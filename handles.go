package gohotdraw

import (
	_ "fmt"
)

const (
	HANDLESIZE = 8
)

func GetHandleDisplayBox(handle Handle) *Rectangle {
	p := handle.Locate()
	return &Rectangle{p.X - HANDLESIZE/2, p.Y - HANDLESIZE/2, HANDLESIZE, HANDLESIZE}
}

func HandleContainsPoint(handle Handle, p *Point) bool {
	return GetHandleDisplayBox(handle).ContainsPoint(p)
}

func DrawHandle(handle Handle, g Graphics) {
	r := GetHandleDisplayBox(handle)
	g.SetFGColor(LightGray)
	g.DrawBorderedRectFromRect(r)
}


type Handle interface {
	Locate() *Point
	InvokeStart(x, y int, view DrawingView)
	InvokeStep(x, y, anchorX, anchorY int, view DrawingView)
	InvokeEnd(x, y int, view DrawingView)
	GetOwner() Figure
}

type DefaultHandle struct {
	owner Figure
}

func NewDefaultHandle(owner Figure) *DefaultHandle {
	return &DefaultHandle{owner: owner}
}

func (this *DefaultHandle) GetOwner() Figure {
	return this.owner
}

func (this *DefaultHandle) InvokeStart(x, y int, view DrawingView) {
	//do nothing, subclasses can implement
}

func (this *DefaultHandle) InvokeStep(x, y, anchorX, anchorY int, view DrawingView) {
	//do nothing, subclasses can implement
}

func (this *DefaultHandle) InvokeEnd(x, y int, view DrawingView) {
	//do nothing, subclasses can implement
}

type LocatorHandle struct {
	*DefaultHandle
	locator Locator
}

func NewLocatorHandle(owner Figure, l Locator) *LocatorHandle {
	handle := &LocatorHandle{}
	handle.DefaultHandle = NewDefaultHandle(owner)
	handle.locator = l
	return handle
}

func (this *LocatorHandle) GetLocator() Locator {
	return this.locator
}

//Locates the handle on the figure.
//The handle is drawn centered around the returned point.
func (this *LocatorHandle) Locate() *Point {
	return this.locator.Locate(this.GetOwner())
}


//type NullHandle struct {
//	*LocatorHandle
//}

//func NewNullHandle(owner Figure, locator Locator) *NullHandle {
//	handle := &NullHandle{}
//	handle.LocatorHandle = NewLocatorHandle(owner, locator)
//	return handle
//}

//func (this *NullHandle) Draw(g Graphics) {
//	r := GetHandleDisplayBox(this)
//	g.SetFGColor(Red)
//	g.DrawBorderedRect(r.X, r.Y, r.Width, r.Height)
//}
