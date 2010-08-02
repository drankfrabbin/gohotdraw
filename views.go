package gohotdraw

import (
	_ "fmt"
)

type DrawingView interface {
	GetDrawing() Drawing
	SetDrawing(drawing Drawing)
	GetGraphics() Graphics
	SetGraphics(g Graphics)
	SetEditor(editor DrawingEditor)
	Add(figure Figure) Figure
	Remove(figure Figure) Figure

	AddFigureSelectionListener(l FigureSelectionListener)
	RemoveFigureSelectionListener(l FigureSelectionListener)
	FireSelectionChanged()

	IsFigureSelected(figure Figure) bool
	AddToSelection(figure Figure)
	RemoveFromSelection(figure Figure)
	ClearSelection()
	AddAllToSelection(figures *Set)
	GetSelection() *Set
	ToggleSelection(figure Figure)

	FindHandle(p *Point) Handle

	SetUpdateStrategy(strategy Painter)
	// Paints the drawing view. The actual drawing is delegated to
	// the current update strategy.
	Repaint()

	// Draws the contents of the drawing view.
	// The view has three layers: background, drawing, handles.
	// The layers are drawn in back to front order.
	Draw(g Graphics)
	//drawDrawing(g Graphics)
}

type StandardDrawingView struct {
	drawing            Drawing
	eventHandler       *EventHandler
	editor             DrawingEditor
	updateStrategy     Painter
	graphics           Graphics
	selection          *Set
	selectionListeners *Set
	selectionHandles   *Set
}

func NewStandardDrawingView() *StandardDrawingView {
	view := &StandardDrawingView{}
	view.eventHandler = NewEventHandler(view)
	view.updateStrategy = &SimpleUpdateStrategy{}
	view.selection = NewSet()
	view.selectionListeners = NewSet()
	return view
}

func (this *StandardDrawingView) GetDrawing() Drawing {
	return this.drawing
}

func (this *StandardDrawingView) SetDrawing(drawing Drawing) {
	if this.drawing != nil {
		//TODO this.ClearSelection()
		//TODO this.drawing.RemoveFigureListener(this)
	}
	this.drawing = drawing
	if this.drawing != nil {
		this.drawing.AddFigureListener(this.eventHandler)
	}
	//TODO this.CheckMinimumSize()
	this.Repaint()
}

func (this *StandardDrawingView) SetEditor(editor DrawingEditor) {
	this.editor = editor
}

func (this *StandardDrawingView) SetGraphics(g Graphics) {
	this.graphics = g
}

func (this *StandardDrawingView) GetGraphics() Graphics {
	return this.graphics
}

func (this *StandardDrawingView) SetUpdateStrategy(p Painter) {
	this.updateStrategy = p
}

func (this *StandardDrawingView) Add(figure Figure) Figure {
	return this.drawing.Add(figure)
}

func (this *StandardDrawingView) Remove(figure Figure) Figure {
	return this.drawing.Remove(figure)
}

func (this *StandardDrawingView) GetSelection() *Set {
	return this.selection.Clone()
}

func (this *StandardDrawingView) IsFigureSelected(figure Figure) bool {
	return this.selection.Contains(figure)
}

func (this *StandardDrawingView) AddToSelection(figure Figure) {
	if !this.IsFigureSelected(figure) {
		this.selection.Push(figure)
		this.selectionHandles = nil
		this.FireSelectionChanged()
	}
}

func (this *StandardDrawingView) AddAllToSelection(figures *Set) {
	for currentFigure := range figures.Iter() {
		this.AddToSelection(currentFigure.(Figure))
	}
}

func (this *StandardDrawingView) RemoveFromSelection(figure Figure) {
	if !this.IsFigureSelected(figure) {
		for i := 0; i < this.selection.Len(); i++ {
			currentFigure := this.selection.At(i).(Figure)
			if currentFigure == figure {
				this.selection.Delete(i)
			}
		}
		this.selectionHandles = nil
		this.FireSelectionChanged()
	}
}

func (this *StandardDrawingView) ClearSelection() {
	if this.selectionHandles == nil {
		return
	}
	this.selection = NewSet()
	this.selectionHandles = nil
	this.FireSelectionChanged()
}

func (this *StandardDrawingView) ToggleSelection(figure Figure) {
	if this.IsFigureSelected(figure) {
		this.RemoveFromSelection(figure)
	} else {
		this.AddToSelection(figure)
	}
	this.FireSelectionChanged()
}


func (this *StandardDrawingView) AddFigureSelectionListener(l FigureSelectionListener) {
	this.selectionListeners.Add(l)
}
func (this *StandardDrawingView) RemoveFigureSelectionListener(l FigureSelectionListener) {
	for i := 0; i < this.selectionListeners.Len(); i++ {
		currentListener := this.selectionListeners.At(i).(FigureSelectionListener)
		if currentListener == l {
			this.selectionListeners.Delete(i)
			return
		}
	}
}

func (this *StandardDrawingView) FireSelectionChanged() {
	for i := 0; i < this.selectionListeners.Len(); i++ {
		currentListener := this.selectionListeners.At(i).(FigureSelectionListener)
		currentListener.FigureSelectionChanged(this)
	}
}

func (this *StandardDrawingView) GetSelectionHandles() *Set {
	if this.selectionHandles == nil {
		this.selectionHandles = NewSet()
		selectedFigures := this.GetSelection()
		for f := 0; f < selectedFigures.Len(); f++ {
			currentFigure := selectedFigures.At(f).(Figure)
			currentHandles := currentFigure.GetHandles()
			for h := 0; h < currentHandles.Len(); h++ {
				this.selectionHandles.Push(currentHandles.At(h))
			}
		}
	}
	return this.selectionHandles
}

func (this *StandardDrawingView) FindHandle(p *Point) Handle {
	var currentHandle Handle
	handles := this.GetSelectionHandles()
	for i := 0; i < handles.Len(); i++ {
		currentHandle = handles.At(i).(Handle)
		if HandleContainsPoint(currentHandle, p) {
			return currentHandle
		}
	}
	return nil
}

func (this *StandardDrawingView) drawHandles(g Graphics) {
	selectionHandles := this.GetSelectionHandles()
	for i := 0; i < selectionHandles.Len(); i++ {
		currentHandle := selectionHandles.At(i).(Handle)
		DrawHandle(currentHandle, g)
	}
}

//Draws background, drawing, foreground, and handles
func (this *StandardDrawingView) Draw(g Graphics) {
	this.drawBackground(g)
	this.drawForeground(g)
	this.drawHandles(g)
}


func (this *StandardDrawingView) drawBackground(g Graphics) {
	width := g.GetWindowSize().Width
	height := g.GetWindowSize().Height
	//fmt.Printf("height: %v, width: %v\n", width, height)
	g.SetFGColor(White)
	//fmt.Println("paint background")
	g.DrawRectFromRect(&Rectangle{0, 0, width, height})
}

//Draws the drawing of the view
func (this *StandardDrawingView) drawForeground(g Graphics) {
	this.drawing.Draw(g)
}

func (this *StandardDrawingView) Repaint() {
	if this.graphics != nil {
		if this.updateStrategy != nil {
			this.updateStrategy.Draw(this.graphics, this)
		}
	}
}
