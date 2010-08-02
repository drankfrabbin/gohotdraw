package gohotdraw

import (
	"testing"
)

func TestDrawingView(t *testing.T) {
	//TODO How to test Listeners?
	view := NewStandardDrawingView()
	view.SetGraphics(NewDefaultXGBGraphics())
	view.SetUpdateStrategy(&SimpleUpdateStrategy{})
	drawing := NewStandardDrawing()
	view.SetDrawing(drawing)
	fig1 := NewRectangleFigureFromPoints(&Point{1, 1}, &Point{3, 3})
	fig2 := NewRectangleFigureFromPoints(&Point{2, 1}, &Point{3, 3})
	fig3 := NewRectangleFigureFromPoints(&Point{3, 1}, &Point{3, 3})
	drawing.Add(fig1)
	drawing.Add(fig2)
	drawing.Add(fig3)
	drawing.Remove(fig2)
	fig1.MoveBy(fig1, 0, 2)
	fig2.MoveBy(fig2, 0, 2)
	drawing.Add(fig2)
}
