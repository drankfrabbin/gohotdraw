package gohotdraw

import "testing"

func testRect(t *testing.T, rect *Rectangle, expX, expY, expWidth, expHeight int) {
	if rect.X != expX {
		t.Errorf("%T: X) expected %d, got %d", rect, expX, rect.X)
	}
	if rect.Y != expY {
		t.Errorf("%T: Y) expected %d, got %d", rect, expY, rect.Y)
	}
	if rect.Width != expWidth {
		t.Errorf("%T: Width) expected %d, got %d", rect, expWidth, rect.Width)
	}
	if rect.Height != expHeight {
		t.Errorf("%T: Height) expected %d, got %d", rect, expHeight, rect.Height)
	}
}

func assertEqInt(t *testing.T, message string, a, b int) {
	if a != b {
		t.Errorf("%s - expected: %v, was: %v", message, a, b)
	}
}

func assertTrue(t *testing.T, message string, value bool) {
	if !value {
		t.Errorf("%s - expected: %t, was: %t", message, true, value)
	}
}
