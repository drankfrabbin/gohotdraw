package gohotdraw

import (
	"testing"
)

func TestGetColor(t *testing.T) {
	var expR, expG, expB uint32 = 100,255,155
	rgb := NewColor(expR, expG, expB)
	r,g,b := rgb.GetChannels()
	if expR != r || expG != g || expB != b {
		t.Errorf("expected: %d, %d, %d; actual: %d, %d, %d",expR, expG, expB,r,g,b)
	}
}
