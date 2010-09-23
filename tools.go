package gohotdraw

import (
	"math"
)

type Tool interface {
	MouseListener
	MouseMotionListener
	IsActive() bool
	Activate()
	Deactivate()
}

type DefaultTool struct {
	editor           DrawingEditor
	isActive         bool
	anchorX, anchorY int
}

func newDefaultTool(editor DrawingEditor) *DefaultTool {
	tool := &DefaultTool{}
	tool.editor = editor
	return tool
}

func (this *DefaultTool) Activate() {
	this.editor.GetView().ClearSelection()
	this.isActive = true
}

func (this *DefaultTool) Deactivate() {
	this.isActive = false
}

func (this *DefaultTool) IsActive() bool {
	return this.isActive
}

func (this *DefaultTool) MouseDown(e *MouseEvent) {
	this.anchorX = e.X
	this.anchorY = e.Y
}

func (this *DefaultTool) MouseDrag(e *MouseEvent) {}
func (this *DefaultTool) MouseUp(e *MouseEvent)   {}


type CreationTool struct {
	*DefaultTool
	anchor        *Point
	createdFigure Figure
	prototype     Figure
}

func NewCreationTool(editor DrawingEditor, prototype Figure) *CreationTool {
	tool := &CreationTool{}
	tool.DefaultTool = newDefaultTool(editor)
	tool.prototype = prototype
	return tool
}

func (this *CreationTool) MouseDown(e *MouseEvent) {
	this.anchor = e.GetPoint()
	this.createdFigure = this.createFigure()
	this.createdFigure.SetDisplayBox(this.createdFigure, this.anchor, this.anchor)
	this.editor.GetView().Add(this.createdFigure)
}

func (this *CreationTool) MouseDrag(e *MouseEvent) {
	this.createdFigure.SetDisplayBox(this.createdFigure, this.anchor, e.GetPoint())
}

func (this *CreationTool) MouseUp(e *MouseEvent) {
	if this.createdFigure.IsEmpty(this.createdFigure) {
		this.editor.GetView().Remove(this.createdFigure)
	}
	this.createdFigure = nil
	this.editor.ToolDone()
}

func (this *CreationTool) createFigure() Figure {
	if this.prototype == nil {
		panic("No prototype defined")
	}
	return this.prototype.Clone()
}


type NullTool struct {
	*DefaultTool
}

func NewNullTool(editor DrawingEditor) *NullTool {
	return &NullTool{newDefaultTool(editor)}
}


type SelectionTool struct {
	*DefaultTool
	currentTool Tool
}

func NewSelectionTool(editor DrawingEditor) *SelectionTool {
	tool := &SelectionTool{}
	tool.DefaultTool = newDefaultTool(editor)
	return tool
}

func (this *SelectionTool) MouseDown(e *MouseEvent) {
	selectedHandle := this.editor.GetView().FindHandle(e.GetPoint())
	//fmt.Printf("handle selected: %v\n", selectedHandle)
	if selectedHandle != nil {
		this.currentTool = NewHandleTracker(this.editor, selectedHandle)
	} else {
		selectedFigure := this.editor.GetView().GetDrawing().FindFigure(e.GetPoint())
		if selectedFigure != nil {
			this.currentTool = NewDragTracker(this.editor, selectedFigure)
			//rect := selectedFigure.GetDisplayBox()
			//fmt.Printf("found figure - X: %v, Y: %v, W: %v, H: %v\n", rect.X, rect.Y, rect.Width, rect.Height)
		} else {
			if !e.IsShiftDown() {
				this.editor.GetView().ClearSelection()
			}
			this.currentTool = NewAreaTracker(this.editor)
		}
	}
	this.currentTool.MouseDown(e)
	//this.currentTool.Activate()
}

func (this *SelectionTool) MouseDrag(e *MouseEvent) {
	if this.currentTool != nil {
		this.currentTool.MouseDrag(e)
	}
}

func (this *SelectionTool) MouseUp(e *MouseEvent) {
	if this.currentTool != nil {
		this.currentTool.MouseUp(e)
		this.currentTool.Deactivate()
		this.currentTool = nil
	}
}


type DragTracker struct {
	*DefaultTool
	anchorFigure Figure
	hasMoved     bool
	lastX, lastY int
}

func NewDragTracker(editor DrawingEditor, anchorFigure Figure) *DragTracker {
	tracker := &DragTracker{}
	tracker.DefaultTool = newDefaultTool(editor)
	tracker.anchorFigure = anchorFigure
	return tracker
}

func (this *DragTracker) MouseDown(e *MouseEvent) {
	this.DefaultTool.MouseDown(e)
	this.lastX = e.X
	this.lastY = e.Y
	if e.IsShiftDown() {
		this.editor.GetView().ToggleSelection(this.anchorFigure)
		this.anchorFigure = nil
	} else if !this.editor.GetView().IsFigureSelected(this.anchorFigure) {
		this.editor.GetView().ClearSelection()
		this.editor.GetView().AddToSelection(this.anchorFigure)
	}
}

func (this *DragTracker) MouseDrag(e *MouseEvent) {
	this.DefaultTool.MouseDrag(e)
	this.hasMoved = math.Fabs(float64(e.X-this.anchorX)) > 3 ||
		math.Fabs(float64(e.Y-this.anchorY)) > 3
	if this.hasMoved {
		selectedFigures := this.editor.GetView().GetSelection()
		for i := 0; i < selectedFigures.Len(); i++ {
			selectedFigure := selectedFigures.At(i).(Figure)
			selectedFigure.MoveBy(selectedFigure, e.X-this.lastX, e.Y-this.lastY)
		}
	}
	this.lastX = e.X
	this.lastY = e.Y
}


type AreaTracker struct {
	*DefaultTool
	rubberband *Rectangle
}

func NewAreaTracker(editor DrawingEditor) *AreaTracker {
	tracker := &AreaTracker{}
	tracker.DefaultTool = newDefaultTool(editor)
	return tracker
}

func (this *AreaTracker) MouseDown(e *MouseEvent) {
	this.DefaultTool.MouseDown(e)
	this.resizeRubberband(this.anchorX, this.anchorY, this.anchorX, this.anchorY)
}

func (this *AreaTracker) MouseDrag(e *MouseEvent) {
	this.DefaultTool.MouseDrag(e)
	this.drawRubberband()
	this.resizeRubberband(this.anchorX, this.anchorY, e.X, e.Y)
}

func (this *AreaTracker) MouseUp(e *MouseEvent) {
	this.DefaultTool.MouseUp(e)
	this.selectGroup(e.IsShiftDown())
	this.resizeRubberband(-1, -1, 0, 0)
}

func (this *AreaTracker) resizeRubberband(x1, y1, x2, y2 int) {
	this.rubberband = NewRectangleFromPoints(&Point{x1, y1}, &Point{x2, y2})
	this.drawRubberband()
}

func (this *AreaTracker) drawRubberband() {
	g := this.editor.GetView().GetGraphics()
	if g != nil {
		this.editor.GetView().Repaint()
		g.SetFGColor(GreyBlue)
		g.DrawBorderFromRectDirectly(this.rubberband)
	}
}

func (this *AreaTracker) selectGroup(toggle bool) {
	figures := this.editor.GetView().GetDrawing().GetFigures()
	for i := 0; i < figures.Len(); i++ {
		currentFigure := figures.At(i).(Figure)
		rect := currentFigure.GetDisplayBox()
		if this.rubberband.Contains(rect.X, rect.Y) && this.rubberband.Contains(rect.X+rect.Width, rect.Y+rect.Height) {
			if toggle {
				this.editor.GetView().ToggleSelection(currentFigure)
			} else {
				this.editor.GetView().AddToSelection(currentFigure)
			}
		}
	}
}

type HandleTracker struct {
	*DefaultTool
	anchorHandle Handle
}

func NewHandleTracker(editor DrawingEditor, anchorHandle Handle) *HandleTracker {
	tracker := &HandleTracker{}
	tracker.DefaultTool = newDefaultTool(editor)
	tracker.anchorHandle = anchorHandle
	return tracker
}

func (this *HandleTracker) MouseDown(e *MouseEvent) {
	this.DefaultTool.MouseDown(e)
	this.anchorHandle.InvokeStart(e.X, e.Y, this.editor.GetView())
}

func (this *HandleTracker) MouseDrag(e *MouseEvent) {
	this.DefaultTool.MouseDrag(e)
	this.anchorHandle.InvokeStep(e.X, e.Y, this.anchorX, this.anchorY, this.editor.GetView())
}

func (this *HandleTracker) MouseUp(e *MouseEvent) {
	this.DefaultTool.MouseUp(e)
	this.anchorHandle.InvokeEnd(e.X, e.Y, this.editor.GetView())
}
