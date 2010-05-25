package gohotdraw

import "testing"

func TestTranslate(t *testing.T) {
	rect := &Rectangle{2, 3, 1, 1}
	rect.Translate(-2, 2)
	expX := 0
	expY := 5
	expWidth := 1
	expHeight := 1
	testRect(t, rect, expX, expY, expWidth, expHeight)
}

func TestGrow(t *testing.T) {
	rect := &Rectangle{5, 5, 1, 1}
	expX := 3
	expY := 3
	expWidth := 5
	expHeight := 5
	rect.Grow(2, 2)
	testRect(t, rect, expX, expY, expWidth, expHeight)
}

func TestUnion(t *testing.T) {
	rect1 := &Rectangle{3, 3, 1, 1}
	rect2 := &Rectangle{1, 5, 2, 1}
	expX := 1
	expY := 3
	expWidth := 3
	expHeight := 3
	u1 := rect1.Union(rect2)
	testRect(t, u1, expX, expY, expWidth, expHeight)
	u2 := rect2.Union(rect1)
	testRect(t, u2, expX, expY, expWidth, expHeight)
}

func TestAddInternal(t *testing.T) {
	rect := &Rectangle{2, 3, 5, 5}
	rect.Add(3, 4)
	expX := 2
	expY := 3
	expWidth := 5
	expHeight := 5
	testRect(t, rect, expX, expY, expWidth, expHeight)
}

func TestAddExternal(t *testing.T) {
	rect := &Rectangle{2, 3, 5, 5}
	rect.Add(8, 8)
	expX := 2
	expY := 3
	expWidth := 6
	expHeight := 5
	testRect(t, rect, expX, expY, expWidth, expHeight)
}

func TestContains(t *testing.T) {
	rect := &Rectangle{2, 3, 5, 5}
	if !rect.Contains(3, 4) {
		t.Errorf("point in rect not contained, but should")
	}
	if rect.Contains(1, 1) {
		t.Errorf("point outside contained, but shouldn't")
	}
	if !rect.Contains(2, 3) {
		t.Errorf("top left conner not contained, but should")
	}
	if rect.Contains(7, 8) {
		t.Errorf("bottom right conner contained, but shouldn't")
	}
	if rect.Contains(4, 8) {
		t.Errorf("bottom boundary contained, but shouldn't")
	}
	if rect.Contains(7, 5) {
		t.Errorf("right boundary contained, but shouldn't")
	}
}
