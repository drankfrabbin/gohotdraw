package gohotdraw

import (
	"testing"
	"fmt"
)

func TestDrawingView(t *testing.T) {
	//TODO How to test Listeners?
	view := NewStandardDrawingView(0, 0)
	view.SetGraphics(NewConsoleGraphics())
	view.SetUpdateStrategy(&SimpleUpdateStrategy{})
	drawing := NewStandardDrawing()
	view.SetDrawing(drawing)
	fig1 := NewRectangleFigureFromPoints(&Point{1, 1}, &Point{3, 3})
	fig2 := NewRectangleFigureFromPoints(&Point{2, 1}, &Point{3, 3})
	fig3 := NewRectangleFigureFromPoints(&Point{3, 1}, &Point{3, 3})
	fmt.Println("before")
	drawing.Add(fig1)
	fmt.Println("added fig1")
	drawing.Add(fig2)
	fmt.Println("added fig2")
	drawing.Add(fig3)
	fmt.Println("added fig3")
	drawing.Remove(fig2)
	fmt.Println("removed fig2")
	MoveBy(fig1, 0, 2)
	fmt.Println("moved fig1")
	MoveBy(fig2, 0, 2)
	drawing.Add(fig2)
}
