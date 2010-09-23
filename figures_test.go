package gohotdraw

import (
	"testing"
	_ "fmt"
)

func TestDefaultRectFigure(t *testing.T) {
	fig := NewRectangleFigure()
	db := fig.GetDisplayBox()
	expX := 0
	expY := 0
	expWidth := 0
	expHeight := 0
	testRect(t, db, expX, expY, expWidth, expHeight)
}

func TestRectFigure(t *testing.T) {
	fig := NewBorderDecorator(NewRectangleFigureFromPoints(&Point{1, 2}, &Point{4, 6}))
	db := fig.GetDisplayBox()
	expX := 1
	expY := 2
	expWidth := 3
	expHeight := 4
	testRect(t, db, expX, expY, expWidth, expHeight)
}

func TestRectFigureGetSize(t *testing.T) {
	fig := NewBorderDecorator(NewRectangleFigureFromPoints(&Point{1, 2}, &Point{4, 6}))
	dim := fig.GetSize(fig)
	expWidth := 3
	expHeight := 4
	if dim.Width != expWidth {
		t.Errorf("%T: Width) expected %d, got %d", dim, expWidth, dim.Width)
	}
	if dim.Height != expHeight {
		t.Errorf("%T: Height) expected %d, got %d", dim, expHeight, dim.Height)
	}
}

func TestMoveRectFigure(t *testing.T) {
	fig := NewBorderDecorator(NewRectangleFigureFromPoints(&Point{3, 3}, &Point{4, 4}))
	fig.MoveBy(fig, 2, -1)
	db := fig.GetDisplayBox()
	expX := 5
	expY := 2
	expWidth := 1
	expHeight := 1
	testRect(t, db, expX, expY, expWidth, expHeight)
}

func TestCompositeAddFigures(t *testing.T) {
	cf := NewCompositeFigure()
	fig1 := NewBorderDecorator(NewRectangleFigureFromPoints(&Point{1, 1}, &Point{1, 3}))
	fig2 := NewBorderDecorator(NewRectangleFigureFromPoints(&Point{2, 2}, &Point{5, 5}))
	cf.Add(fig1)
	cf.Add(fig2)
	if cf.GetFigures().Len() != 2 {
		t.Errorf("%T: figure count) expected %d, got %d", cf, 2, cf.GetFigures().Len())
	}
	if cf.GetFigureCount() != 2 {
		t.Errorf("%T: figure count) expected %d, got %d", cf, 2, cf.GetFigureCount())
	}
}

func TestMoveComposite(t *testing.T) {
	cf := NewCompositeFigure()
	fig1 := NewBorderDecorator(NewRectangleFigureFromPoints(&Point{1, 1}, &Point{3, 3}))
	fig2 := NewBorderDecorator(NewRectangleFigureFromPoints(&Point{2, 2}, &Point{5, 5}))
	cf.Add(fig1)
	cf.Add(fig2)
	cf.MoveBy(cf, 1, -1)
	fig1Box := fig1.GetDisplayBox()
	testRect(t, fig1Box, 2, 0, 2, 2)
	fig2Box := fig2.GetDisplayBox()
	testRect(t, fig2Box, 3, 1, 3, 3)
}

func TestDiplayBoxComposite(t *testing.T) {
	cf := NewCompositeFigure()
	fig1 := NewRectangleFigureFromPoints(&Point{1, 1}, &Point{3, 3})
	fig2 := NewRectangleFigureFromPoints(&Point{2, 2}, &Point{5, 5})
	cf.Add(fig1)
	cf.Add(fig2)
	box := cf.GetDisplayBox()
	testRect(t, box, 1, 1, 4, 4)
}

func TestCompIncludes(t *testing.T) {
	cf := NewCompositeFigure()
	fig1 := NewRectangleFigureFromPoints(&Point{1, 1}, &Point{3, 3})
	fig2 := NewRectangleFigureFromPoints(&Point{1, 1}, &Point{3, 3})
	cf.Add(fig1)
	if !fig1.Includes(fig1) {
		t.Errorf("fig1 should inlcude itself")
	}
	if cf.Includes(fig2) {
		t.Errorf("fig2 should not be in composite")
	}
	if !cf.Includes(fig1) {
		t.Errorf("fig1 should be in composite")
	}
}

func TestBorderDecoratorIncludes(t *testing.T) {
	fig1 := NewBorderDecorator(NewRectangleFigureFromPoints(&Point{1, 1}, &Point{3, 3}))
	if !fig1.Includes(fig1) {
		t.Errorf("fig1 should inlcude itself")
	}
}

func TestCompNoDuplicateFigures(t *testing.T) {
	cf := NewCompositeFigure()
	fig := NewRectangleFigureFromPoints(&Point{1, 1}, &Point{3, 3})
	cf.Add(fig)
	cf.Add(fig)
	assertEqInt(t, "figure count", 1, cf.GetFigureCount())
}

func TestCompReplaceFigure(t *testing.T) {
	cf := NewCompositeFigure()
	fig1 := NewBorderDecorator(NewRectangleFigureFromPoints(&Point{1, 1}, &Point{3, 3}))
	figToReplace := NewBorderDecorator(NewRectangleFigureFromPoints(&Point{1, 1}, &Point{3, 3}))
	fig2 := NewBorderDecorator(NewRectangleFigureFromPoints(&Point{1, 1}, &Point{3, 3}))
	figRepl := NewBorderDecorator(NewRectangleFigureFromPoints(&Point{1, 1}, &Point{3, 3}))
	cf.Add(fig1)
	cf.Add(figToReplace)
	cf.Add(fig2)
	cf.Replace(figToReplace, figRepl)

	assertTrue(t, "fig1 included", cf.Includes(fig1))
	assertTrue(t, "fig2 included", cf.Includes(fig2))
	assertTrue(t, "figToReplace not included", !cf.Includes(figToReplace))
	assertTrue(t, "figRepl included", cf.Includes(figRepl))
}

func TestStandardDrawing(t *testing.T) {
	drawing := NewStandardDrawing()
	fig1 := NewRectangleFigureFromPoints(&Point{1, 1}, &Point{3, 3})
	fig2 := NewRectangleFigureFromPoints(&Point{1, 1}, &Point{3, 3})
	drawing.Add(fig1)
	drawing.Add(fig2)
	assertEqInt(t, "figure count", 2, drawing.GetFigureCount())
}

func TestStandardDrawingRemove(t *testing.T) {
	drawing := NewStandardDrawing()
	fig1 := NewRectangleFigureFromPoints(&Point{1, 1}, &Point{3, 3})
	fig2 := NewRectangleFigureFromPoints(&Point{1, 1}, &Point{3, 3})
	drawing.Add(fig1)
	drawing.Add(fig2)
	drawing.Remove(fig1)
	assertEqInt(t, "figure count", 1, drawing.GetFigureCount())
}

func TestStandardDrawingHandles(t *testing.T) {
	drawing := NewStandardDrawing()
	assertEqInt(t, "handle count", 4, drawing.GetHandles().Len())
}

func TestListeners(t *testing.T) {
	drawing := NewStandardDrawing()
	fig1 := NewRectangleFigureFromPoints(&Point{1, 1}, &Point{3, 3})
	assertEqInt(t, "no of listeners: fig1", 0, fig1.GetListeners().Len())
	fig2 := NewRectangleFigureFromPoints(&Point{1, 1}, &Point{3, 3})
	assertEqInt(t, "no of listeners: fig2", 0, fig2.GetListeners().Len())
	drawing.Add(fig1)
	assertEqInt(t, "no of listeners: fig1", 1, fig1.GetListeners().Len())
	drawing.Add(fig2)
	assertEqInt(t, "no of listeners: fig2", 1, fig2.GetListeners().Len())
	drawing.Remove(fig2)
	assertEqInt(t, "no of listeners: fig2", 0, fig2.GetListeners().Len())
}
