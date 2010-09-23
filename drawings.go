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

func (this *StandardDrawing) SetTitle(title string) {
	this.title = title
}

func (this *StandardDrawing) GetTitle() string {
	return this.title
}

func (this *StandardDrawing) FindFigure(point *Point) Figure {
	//this.figures is a Set. A Set can contain any object (interface{})
	for figure := range this.figures.Iter() { // iterates over the elements of figures, type of figure is interface{}
		if figure.(Figure).Contains(point) { // x.(Figure) is type assertion
			return figure.(Figure) // type assertion (http://golang.org/doc/go_spec.html#Type_assertions)
		}
	}
	return nil
}
