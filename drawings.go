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
	for figure := range this.figures.Iter() { 
		if figure.(Figure).Contains(point) { 
			return figure.(Figure) 
		}
	}
	return nil
}
