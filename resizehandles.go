package gohotdraw

import (
	"math"
)

func AddCornerHandles(f Figure, handles *Set) {
	handles.Push(newSouthEastHandle(f))
	handles.Push(newSouthWestHandle(f))
	handles.Push(newNorthEastHandle(f))
	handles.Push(newNorthWestHandle(f))
}

func AddAllHandles(f Figure, handles *Set) {
	AddCornerHandles(f, handles)
	handles.Push(newSouthHandle(f))
	handles.Push(newEastHandle(f))
	handles.Push(newNorthHandle(f))
	handles.Push(newWestHandle(f))
}

type resizeHandle struct {
	*LocatorHandle
}

func newResizeHandle(owner Figure, loc Locator) *resizeHandle {
	handle := &resizeHandle{}
	handle.LocatorHandle = NewLocatorHandle(owner, loc)
	return handle
}

func InvokeStart(x, y int, view DrawingView) {
	//TODO
}

func InvokeEnd(x, y int, view DrawingView) {
	//TODO
}


type northEastHandle struct {
	*resizeHandle
}

func newNorthEastHandle(owner Figure) *northEastHandle {
	return &northEastHandle{newResizeHandle(owner, CreateNorthEastLocator())}
}

func (this *northEastHandle) InvokeStep(x, y, anchorX, anchorY int, view DrawingView) {
	r := this.owner.GetDisplayBox()
	this.owner.SetDisplayBox(
		this.owner, 
		&Point{r.X, int(math.Fmin(float64(r.Y+r.Height), float64(y)))}, 
		&Point{int(math.Fmax(float64(r.X), float64(x))), r.Y + r.Height})
}

type eastHandle struct {
	*resizeHandle
}

func newEastHandle(owner Figure) *eastHandle {
	return &eastHandle{newResizeHandle(owner, CreateEastLocator())}
}

func (this *eastHandle) InvokeStep(x, y, anchorX, anchorY int, view DrawingView) {
	r := this.owner.GetDisplayBox()
	this.owner.SetDisplayBox(
		this.owner, 
		&Point{r.X, r.Y}, 
		&Point{int(math.Fmax(float64(r.X), float64(x))), r.Y + r.Height})
}

type northHandle struct {
	*resizeHandle
}

func newNorthHandle(owner Figure) *northHandle {
	return &northHandle{newResizeHandle(owner, CreateNorthLocator())}
}

func (this *northHandle) InvokeStep(x, y, anchorX, anchorY int, view DrawingView) {
	r := this.owner.GetDisplayBox()
	topLeft := &Point{r.X, int(math.Fmin(float64(r.Y+r.Height), float64(y)))}
	bottomRight := &Point{r.X + r.Width, r.Y + r.Height}
	this.owner.SetDisplayBox(this.owner, topLeft, bottomRight)
}

type northWestHandle struct {
	*resizeHandle
}

func newNorthWestHandle(owner Figure) *northWestHandle {
	return &northWestHandle{newResizeHandle(owner, CreateNorthWestLocator())}
}

func (this *northWestHandle) InvokeStep(x, y, anchorX, anchorY int, view DrawingView) {
	r := this.owner.GetDisplayBox()
	topLeft := &Point{int(math.Fmin(float64(r.X+r.Width), float64(x))), int(math.Fmin(float64(r.Y+r.Height), float64(y)))}
	bottomRight := &Point{r.X + r.Width, r.Y + r.Height}
	this.owner.SetDisplayBox(this.owner, topLeft, bottomRight)
}

type southEastHandle struct {
	*resizeHandle
}

func newSouthEastHandle(owner Figure) *southEastHandle {
	return &southEastHandle{newResizeHandle(owner, CreateSouthEastLocator())}
}

func (this *southEastHandle) InvokeStep(x, y, anchorX, anchorY int, view DrawingView) {
	r := this.owner.GetDisplayBox()
	this.owner.SetDisplayBox(this.owner, &Point{r.X, r.Y}, 
		&Point{int(math.Fmax(float64(r.X), float64(x))), int(math.Fmax(float64(r.Y), float64(y)))})
}

type southHandle struct {
	*resizeHandle
}

func newSouthHandle(owner Figure) *southHandle {
	return &southHandle{newResizeHandle(owner, CreateSouthLocator())}
}

func (this *southHandle) InvokeStep(x, y, anchorX, anchorY int, view DrawingView) {
	r := this.owner.GetDisplayBox()
	this.owner.SetDisplayBox(this.owner, &Point{r.X, r.Y}, 
		&Point{r.X + r.Width, int(math.Fmax(float64(r.Y), float64(y)))})
}

type southWestHandle struct {
	*resizeHandle
}

func newSouthWestHandle(owner Figure) *southWestHandle {
	return &southWestHandle{newResizeHandle(owner, CreateSouthWestLocator())}
}

func (this *southWestHandle) InvokeStep(x, y, anchorX, anchorY int, view DrawingView) {
	r := this.owner.GetDisplayBox()
	topLeft := &Point{int(math.Fmin(float64(r.X+r.Width), float64(x))), r.Y}
	bottomRight := &Point{r.X + r.Width, int(math.Fmax(float64(r.Y), float64(y)))}
	this.owner.SetDisplayBox(this.owner, topLeft, bottomRight)
}

type westHandle struct {
	*resizeHandle
}

func newWestHandle(owner Figure) *westHandle {
	return &westHandle{newResizeHandle(owner, CreateWestLocator())}
}

func (this *westHandle) InvokeStep(x, y, anchorX, anchorY int, view DrawingView) {
	r := this.owner.GetDisplayBox()
	this.owner.SetDisplayBox(this.owner, 
		&Point{int(math.Fmin(float64(r.X+r.Width), float64(x))), r.Y}, 
		&Point{r.X + r.Width, r.Y + r.Height})
}
