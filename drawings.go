package gohotdraw

import (
	_ "fmt"
)

type Drawing interface {
	Figure
	Add(figure Figure) Figure
	Remove(figure Figure) Figure
	FindFigure(point *Point) Figure
	GetTitle() string
	SetTitle(title string)
}

type StandardDrawing struct {
	*CompositeFigure
	title string
}

func NewStandardDrawing() *StandardDrawing {
	drawing := &StandardDrawing{}
	drawing.CompositeFigure = NewCompositeFigure()
	return drawing
}

func (this *StandardDrawing) GetHandles() *Set {
	handles := NewSet()
	handles.Push(NewNullHandle(this, CreateNorthWestLocator()))
	handles.Push(NewNullHandle(this, CreateNorthEastLocator()))
	handles.Push(NewNullHandle(this, CreateSouthWestLocator()))
	handles.Push(NewNullHandle(this, CreateSouthEastLocator()))
	return handles
}

func (this *StandardDrawing) SetTitle(title string) {
	this.title = title
}

func (this *StandardDrawing) GetTitle() string {
	return this.title
}

func (this *StandardDrawing) FindFigure(point *Point) Figure {
	for figure := range this.figures.Iter() {
		if figure.(Figure).Contains(point) {
			return figure.(Figure)
		}
	}
	return nil
}

//func (this *StandardDrawing) FigureRequestRemove(event *FigureChangeEvent) {
//	figure := event.GetFigure()
//	if Contains(figure, this.figures) {
//		Remove(figure, this.figures)
//		figure.RemoveFigureChangeListener(this)
//		figure.Release()
//	}
//}

//func (this *StandardDrawing) FigureInvalidated(event *FigureChangeEvent) {
//	for i := 0; i < this.listeners.Len(); i++ {
//		currentListener := this.listeners.At(i).(DrawingChangeListener)
//		currentListener.DrawingInvalidated(NewDrawingChangeEvent(this, event.GetInvalidatedRect()))
//	}
//}

//func (this *StandardDrawing) FigureRequestUpdate(event *FigureChangeEvent) {
//	for i := 0; i < this.listeners.Len(); i++ {
//		currentListener := this.listeners.At(i).(DrawingChangeListener)
//		currentListener.DrawingRequestUpdate(NewDrawingChangeEvent(this, nil))
//	}
//}
