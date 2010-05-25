package gohotdraw

import (
	"testing"
)

func TestCreateConnection(t *testing.T) {
	g := NewDefaultXGBGraphics()
	c := g.createConnection()
	if c == nil {
		t.Errorf("connection not established")
	}
}

func TestCreateWindow(t *testing.T) {
	g := NewDefaultXGBGraphics()
	winId := g.createWindow(DEFAULT_X, DEFAULT_Y, DEFAULT_WIDTH, DEFAULT_HEIGHT)
	if winId == 0 {
		t.Errorf("winId not created")
	}
}

func TestCreateContext(t *testing.T) {
	g := NewDefaultXGBGraphics()
	contextId := g.createContext()
	if contextId == 0 {
		t.Errorf("contextId not created")
	}
}
