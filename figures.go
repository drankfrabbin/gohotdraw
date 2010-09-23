package gohotdraw

import (
	_"container/vector"
	"fmt"
)

type Figure interface {
	MoveBy(figure Figure, dx int, dy int)
	basicMoveBy(dx int, dy int)
	changed(figure Figure)
	GetDisplayBox() *Rectangle
	GetSize(figure Figure) *Dimension
	IsEmpty(figure Figure) bool
	Includes(figure Figure) bool
	Draw(g Graphics)
	GetHandles() *Set
	GetFigures() *Set
	SetDisplayBoxRect(figure Figure, rect *Rectangle)
	SetDisplayBox(figure Figure, topLeft, bottomRight *Point)
	setBasicDisplayBox(topLeft, bottomRight *Point)
	
	GetListeners() *Set
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
	listeners *Set
}

func newDefaultFigure() *DefaultFigure {
	return &DefaultFigure{listeners: NewSet()}
}

func (this *DefaultFigure) GetFigures() *Set {
	figures := NewSet()
	figures.Add(this)
	return figures
}

func (this *DefaultFigure) Release(figure Figure) {
	for i := 0; i < this.listeners.Len(); i++ {
		currentListerner := this.listeners.At(i).(FigureListener)
		currentListerner.FigureRemoved(NewFigureEvent(figure))
	}
}

func (this *DefaultFigure) GetZValue() int {
	return this.zValue
}

func (this *DefaultFigure) SetZValue(zValue int) {
	this.zValue = zValue
}

func (this *DefaultFigure) GetListeners() *Set {
	return this.listeners
}

func (this *DefaultFigure) AddFigureListener(l FigureListener) {
	this.listeners.Add(l)
}

func (this *DefaultFigure) RemoveFigureListener(l FigureListener) {
	this.listeners.Remove(l)
}

func (this *DefaultFigure) AddToContainer(figure Figure, l FigureListener) {
	figure.AddFigureListener(l)
	figure.changed(figure)
}

func (this *DefaultFigure) RemoveFromContainer(figure Figure, l FigureListener) {
	figure.changed(figure)
	figure.RemoveFigureListener(l)
}

func (this *DefaultFigure) changed(figure Figure) {
	for i := 0; i < figure.GetListeners().Len(); i++ {
		currentListener := figure.GetListeners().At(i).(FigureListener)
		currentListener.FigureChanged(NewFigureEvent(figure))
	}
}

func (this *DefaultFigure) MoveBy(figure Figure, dx int, dy int) {
	figure.basicMoveBy(dx, dy)
	figure.changed(figure)
}

func (this *DefaultFigure) IsEmpty(figure Figure) bool {
	dimension := figure.GetSize(figure)
	return dimension.Width < 3 || dimension.Height < 3
}

func (this *DefaultFigure) SetDisplayBoxRect(figure Figure, rect *Rectangle) {
	figure.SetDisplayBox(
		figure, 
		&Point{rect.X, rect.Y}, 
		&Point{rect.X + rect.Width, rect.Y + rect.Height})
}

func (this *DefaultFigure) SetDisplayBox(figure Figure, topLeft, bottomRight *Point) {
	figure.setBasicDisplayBox(topLeft, bottomRight)
	figure.changed(figure)
}

func (this *DefaultFigure) GetSize(figure Figure) *Dimension {
	return &Dimension{
		Width: figure.GetDisplayBox().Width, 
		Height: figure.GetDisplayBox().Height}
}







type CompositeFigure struct {
	*DefaultFigure
	//TODO lowestZ
	//TODO highestZ
	figures *Set
}

func NewCompositeFigure() *CompositeFigure {
	return &CompositeFigure{DefaultFigure: newDefaultFigure(), figures: NewSet()}
}

func (this *CompositeFigure) Add(figure Figure) Figure {
	this.figures.Add(figure)
	this.AddToContainer(figure, this)
	fmt.Printf("figure count: %v\n", this.figures.Len())
	return figure
}

func (this *CompositeFigure) AddAll(figures *Set) {
	for i := 0; i < figures.Len(); i++ {
		currentFigure := figures.At(i).(Figure)
		this.Add(currentFigure)
	}
}

func (this *CompositeFigure) Remove(figure Figure) Figure {
	this.figures.Remove(figure)
	this.RemoveFromContainer(figure, this)
	return figure
}

func (this *CompositeFigure) RemoveAll(figures *Set) {
	for i := 0; i < figures.Len(); i++ {
		currentFigure := figures.At(i).(Figure)
		this.Remove(currentFigure)
	}
}

func (this *CompositeFigure) Replace(toBeReplaced, replacement Figure) {
	this.figures.Replace(toBeReplaced, replacement)
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

func (this *CompositeFigure) GetFigures() *Set {
	return this.figures
}

func (this *CompositeFigure) GetFigureCount() int {
	return this.figures.Len()
}

func (this *CompositeFigure) Includes(figure Figure) bool {
	if Figure(this) == figure {
		return true
	}
	return this.figures.Contains(figure)
}

func (this *CompositeFigure) basicMoveBy(x, y int) {
	for i := 0; i < this.figures.Len(); i++ {
		currentFigure := this.figures.At(i).(Figure)
		currentFigure.MoveBy(currentFigure, x, y)
	}
}

func (this *CompositeFigure) Release() {
	this.DefaultFigure.Release(this)
	for i := 0; i < this.figures.Len(); i++ {
		currentFigure := this.figures.At(i).(Figure)
		currentFigure.Release()
	}
}

func (this *CompositeFigure) setBasicDisplayBox(topLeft *Point, bottomRight *Point) {
	//do nothing (How would that work anyway?)
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

func (this *CompositeFigure) GetHandles() *Set {
	handles := NewSet()
	for i := 0; i < this.figures.Len(); i++ {
		currentFigure := this.figures.At(i).(Figure)
		AddAllHandles(currentFigure, handles)
	}
	return handles
}

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
	return NewRectangleFigureFromPoints(
		&Point{rectangle.X, rectangle.Y}, 
		&Point{rectangle.X + rectangle.Width, rectangle.Y + rectangle.Height})
}

func (this *RectangleFigure) Draw(g Graphics) {
	g.SetFGColor(Gray)
	g.DrawRectFromRect(this.GetDisplayBox())
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

func (this *RectangleFigure) GetHandles() *Set {
	handles := NewSet()
	AddAllHandles(this, handles)
	return handles
}

func (this *RectangleFigure) Release() {
	this.DefaultFigure.Release(this)
}

