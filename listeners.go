package gohotdraw

type EventListener interface{}

type FigureListener interface {
	EventListener
	//FigureInvalidated(event *FigureEvent)
	FigureChanged(event *FigureEvent)
	
	//Sent when a figure was added to a drawing
	FigureAdded(event *FigureEvent)
	FigureRemoved(event *FigureEvent)
}

type FigureSelectionListener interface {
	FigureSelectionChanged(view DrawingView)
}

type InputListener interface {
	ExposeListener
	MouseListener
	MouseMotionListener
	KeyListener
}

type MouseListener interface {
	MouseDown(e *MouseEvent)
	MouseUp(e *MouseEvent)
}

type MouseMotionListener interface {
	MouseDrag(e *MouseEvent)
	//MouseMove(e *MouseEvent)
}

type KeyListener interface {
	KeyDown(e *KeyEvent)
	KeyUp(e *KeyEvent)
}

type ExposeListener interface {
	ExposeHappened(e *ExposeEvent)
}

//type ToolListener interface {
//	//ToolStarted(event *ToolEvent)
//	ToolDone(event *ToolEvent)
//}
