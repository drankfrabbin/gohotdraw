package gohotdraw

import (
	"container/vector"
	"fmt"
)

type Figure interface {
	basicMoveBy(dx int, dy int)
	GetDisplayBox() *Rectangle
	Draw(g Graphics)
	GetHandles() *vector.Vector
	GetFigures() *vector.Vector
	setBasicDisplayBox(topLeft, bottomRight *Point)
	Includes(figure Figure) bool

	GetListeners() *vector.Vector
	AddFigureListener(l FigureListener)
	RemoveFigureListener(l FigureListener)
	Release()

	GetZValue() int
	SetZValue(zValue int)
	Clone() Figure
	Contains(point *Point) bool
}

type DefaultFigure struct {
	zValue    int
	listeners *vector.Vector
}

func newDefaultFigure() *DefaultFigure {
	return &DefaultFigure{listeners: new(vector.Vector)}
}

// Does nothing. Subclasses override.
func (this *DefaultFigure) basicMoveBy(dx int, dy int) {
	panic(UNIMPLEMENTED)
}

// Does nothing. Subclasses override.
func (this *DefaultFigure) setBasicDisplayBox(topLeft, bottomRight *Point) {
	panic(UNIMPLEMENTED)
}

// Does nothing. Subclasses override.
func (this *DefaultFigure) GetDisplayBox() *Rectangle {
	panic(UNIMPLEMENTED)
	return nil
}

func (this *DefaultFigure) GetHandles() *vector.Vector {
	panic(UNIMPLEMENTED)
	return nil
}

func (this *DefaultFigure) GetFigures() *vector.Vector {
	figures := new(vector.Vector)
	figures.Push(this)
	return figures
}

func (this *DefaultFigure) Includes(figure Figure) bool {
	panic(UNIMPLEMENTED)
	return false
}

func (this *DefaultFigure) Release() {
	for i := 0; i < this.listeners.Len(); i++ {
		currentListerner := this.listeners.At(i).(FigureListener)
		currentListerner.FigureRemoved(NewFigureEvent(this))
	}
}

func (this *DefaultFigure) GetZValue() int {
	return this.zValue
}

func (this *DefaultFigure) SetZValue(zValue int) {
	this.zValue = zValue
}

// Does nothing. Subclasses override.
func (this *DefaultFigure) Draw(g Graphics) {
	panic(UNIMPLEMENTED)
}

func (this *DefaultFigure) GetListeners() *vector.Vector {
	return this.listeners
}

func (this *DefaultFigure) AddFigureListener(l FigureListener) {
	if !Contains(l, this.listeners) {
		this.listeners.Push(l)
	}
}

func (this *DefaultFigure) RemoveFigureListener(l FigureListener) {
	for i := 0; i < this.listeners.Len(); i++ {
		currentListerner := this.listeners.At(i).(FigureListener)
		if currentListerner == l {
			this.listeners.Delete(i)
			return
		}
	}
}

func (this *DefaultFigure) Clone() Figure {
	return newDefaultFigure()
}

func (this *DefaultFigure) Contains(point *Point) bool {
	panic(UNIMPLEMENTED)
}


func AddToContainer(figure Figure, l FigureListener) {
	figure.AddFigureListener(l)
	changed(figure)
}

func RemoveFromContainer(figure Figure, l FigureListener) {
	changed(figure)
	figure.RemoveFigureListener(l)
}

func MoveBy(figure Figure, dx int, dy int) {
	//	willChange(figure)
	figure.basicMoveBy(dx, dy)
	changed(figure)
}

func SetDisplayBoxRect(figure Figure, rect *Rectangle) {
	SetDisplayBox(figure, &Point{rect.X, rect.Y}, &Point{rect.X + rect.Width, rect.Y + rect.Height})
}

func SetDisplayBox(figure Figure, topLeft, bottomRight *Point) {
	//willChange(figure)
	figure.setBasicDisplayBox(topLeft, bottomRight)
	changed(figure)
}

func GetSize(figure Figure) *Dimension {
	return &Dimension{Width: figure.GetDisplayBox().Width, Height: figure.GetDisplayBox().Height}
}

func IsEmpty(figure Figure) bool {
	dimension := GetSize(figure)
	return dimension.Width < 3 || dimension.Height < 3
}

func changed(figure Figure) {
	//	invalidate(figure)
	for i := 0; i < figure.GetListeners().Len(); i++ {
		currentListener := figure.GetListeners().At(i).(FigureListener)
		currentListener.FigureChanged(NewFigureEvent(figure))
	}
}


//func invalidate(figure Figure) {
//	rect := figure.GetDisplayBox()
//	rect.Grow(HANDLESIZE, HANDLESIZE)
//	for i := 0; i < figure.GetListeners().Len(); i++ {
//		currentListener := figure.GetListeners().At(i).(FigureListener)
//		currentListener.FigureInvalidated(NewFigureEventRect(figure, rect))
//	}
//}

//func willChange(figure Figure) {
//	invalidate(figure)
//}



type CompositeFigure struct {
	*DefaultFigure
	//TODO lowestZ
	//TODO highestZ
	figures *vector.Vector
}

func NewCompositeFigure() *CompositeFigure {
	return &CompositeFigure{DefaultFigure: newDefaultFigure(), figures: new(vector.Vector)}
}

func (this *CompositeFigure) Add(figure Figure) Figure {
	if !this.Includes(figure) {
		this.figures.Push(figure)
		AddToContainer(figure, this)
	}
	fmt.Printf("figure count: %v\n", this.figures.Len())
	return figure
}

func (this *CompositeFigure) AddAll(figures *vector.Vector) {
	for i := 0; i < figures.Len(); i++ {
		currentFigure := figures.At(i).(Figure)
		this.Add(currentFigure)
	}
}

func (this *CompositeFigure) Remove(figure Figure) Figure {
	for i := 0; i < this.figures.Len(); i++ {
		currentFigure := this.figures.At(i).(Figure)
		if currentFigure == figure {
			this.figures.Delete(i)
			RemoveFromContainer(figure, this)
			return figure
		}
	}
	return figure
}

func (this *CompositeFigure) RemoveAll(figures *vector.Vector) {
	for i := 0; i < figures.Len(); i++ {
		currentFigure := figures.At(i).(Figure)
		this.Remove(currentFigure)
	}
}

func (this *CompositeFigure) Replace(toBeReplaced, replacement Figure) {
	for i := 0; i < this.figures.Len(); i++ {
		currentFigure := this.figures.At(i).(Figure)
		if currentFigure == toBeReplaced {
			this.figures.Set(i, replacement)
		}
	}
}

func (this *CompositeFigure) Draw(g Graphics) {
	for i := 0; i < this.figures.Len(); i++ {
		currentFigure := this.figures.At(i).(Figure)
		currentFigure.Draw(g)
	}
}

func (this *CompositeFigure) GetFigureAt(i int) Figure {
	return this.figures.At(i).(Figure)
}

func (this *CompositeFigure) GetFigures() *vector.Vector {
	return this.figures
}

func (this *CompositeFigure) GetFigureCount() int {
	return this.figures.Len()
}

func (this *CompositeFigure) Includes(figure Figure) bool {
	if Figure(this) == figure {
		return true
	}
	return Contains(figure, this.figures)
}

func (this *CompositeFigure) basicMoveBy(x, y int) {
	for i := 0; i < this.figures.Len(); i++ {
		currentFigure := this.figures.At(i).(Figure)
		MoveBy(currentFigure, x, y)
	}
}

func (this *CompositeFigure) Release() {
	this.DefaultFigure.Release()
	for i := 0; i < this.figures.Len(); i++ {
		currentFigure := this.figures.At(i).(Figure)
		currentFigure.Release()
	}
}

func (this *CompositeFigure) GetDisplayBox() *Rectangle {
	if this.figures.Len() > 0 {
		displayBox := this.figures.At(0).(Figure).GetDisplayBox()
		for i := 1; i < this.figures.Len(); i++ {
			currentDisplayBox := this.figures.At(i).(Figure).GetDisplayBox()
			displayBox = displayBox.Union(currentDisplayBox)
		}
		return displayBox
	}
	return &Rectangle{}
}

//func (this *CompositeFigure) FigureInvalidated(event *FigureEvent) {
//	for i := 0; i < this.listeners.Len(); i++ {
//		currentListener := this.listeners.At(i).(FigureListener)
//		currentListener.FigureInvalidated(event)
//	}
//}

func (this *CompositeFigure) FigureChanged(event *FigureEvent) {
	for i := 0; i < this.listeners.Len(); i++ {
		currentListener := this.listeners.At(i).(FigureListener)
		currentListener.FigureChanged(NewFigureEvent(this))
	}
}

func (this *CompositeFigure) FigureAdded(event *FigureEvent) {
	for i := 0; i < this.listeners.Len(); i++ {
		currentListener := this.listeners.At(i).(FigureListener)
		currentListener.FigureAdded(NewFigureEvent(this))
	}
}

func (this *CompositeFigure) FigureRemoved(event *FigureEvent) {
	for i := 0; i < this.listeners.Len(); i++ {
		currentListener := this.listeners.At(i).(FigureListener)
		currentListener.FigureRemoved(NewFigureEvent(this))
	}
}

func (this *CompositeFigure) FigureRequestRemove(event *FigureEvent) {
	for i := 0; i < this.listeners.Len(); i++ {
		currentListener := this.listeners.At(i).(FigureListener)
		currentListener.FigureRequestRemove(NewFigureEvent(this))
	}
}
//func (this *CompositeFigure) FigureRequestUpdate(event *FigureEvent) {
//	for i := 0; i < this.listeners.Len(); i++ {
//		currentListener := this.listeners.At(i).(FigureListener)
//		currentListener.FigureRequestUpdate(event)
//	}
//}


func (this *CompositeFigure) Clone() Figure {
	figure := NewCompositeFigure()
	figure.figures = this.figures
	return figure
}

func (this *CompositeFigure) Contains(point *Point) bool {
	for currentFigure := range this.figures.Iter() {
		if currentFigure.(Figure).Contains(point) {
			return true
		}
	}
	return false
}


type RectangleFigure struct {
	displayBox *Rectangle
	*DefaultFigure
}

func NewRectangleFigure() *RectangleFigure {
	return NewRectangleFigureFromPoints(&Point{0, 0}, &Point{0, 0})
}

func NewRectangleFigureFromPoints(topLeft, bottomRight *Point) *RectangleFigure {
	figure := &RectangleFigure{}
	figure.setBasicDisplayBox(topLeft, bottomRight)
	figure.DefaultFigure = newDefaultFigure()
	return figure
}

func NewRectangleFigureFromRect(rectangle *Rectangle) *RectangleFigure {
	return NewRectangleFigureFromPoints(&Point{rectangle.X, rectangle.Y}, &Point{rectangle.X + rectangle.Width, rectangle.Y + rectangle.Height})
}

func (this *RectangleFigure) Draw(g Graphics) {
	g.SetFGColor(230, 230, 230)
	g.DrawBorderedRectFromRect(this.displayBox)
}

func (this *RectangleFigure) setBasicDisplayBox(topLeft, bottomRight *Point) {
	this.displayBox = NewRectangleFromPoint(topLeft)
	this.displayBox.AddPoint(bottomRight)
}

func (this *RectangleFigure) basicMoveBy(xd, yd int) {
	this.displayBox.Translate(xd, yd)
}

func (this *RectangleFigure) GetDisplayBox() *Rectangle {
	return NewRectangleFromRect(this.displayBox)
}

func (this *RectangleFigure) Includes(figure Figure) bool {
	return Figure(this) == figure
}

func (this *RectangleFigure) Clone() Figure {
	figure := NewRectangleFigure()
	figure.displayBox = this.displayBox
	return figure
}

func (this *RectangleFigure) Contains(point *Point) bool {
	return this.displayBox.ContainsPoint(point)
}

func (this *RectangleFigure) GetHandles() *vector.Vector {
	handles := new(vector.Vector)
	AddAllHandles(this, handles)
	return handles
}
