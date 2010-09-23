package gohotdraw

type DefaultFigureDecorator struct {
	*DefaultFigure
	figure Figure //the decorated figure
}

func NewDefaultFigureDecorator(figure Figure) *DefaultFigureDecorator {
	this := new(DefaultFigureDecorator)
	this.DefaultFigure = newDefaultFigure()
	this.figure = figure
	return this
}

//forward call to contained figure
func (this *DefaultFigureDecorator) Contains(point *Point) bool {
	return this.figure.Contains(point)
}

//forward call to contained figure
func (this *DefaultFigureDecorator) Draw(g Graphics) {
	this.figure.Draw(g)
}

//forward call to contained figure
func (this *DefaultFigureDecorator) GetDisplayBox() *Rectangle {
	return this.figure.GetDisplayBox()
}

//forward call to contained figure
func (this *DefaultFigureDecorator) GetHandles() *Set {
	return this.figure.GetHandles()
}

//forward call to contained figure
func (this *DefaultFigureDecorator) Includes(figure Figure) bool {
	return this.figure.Includes(figure)
}

//forward call to contained figure
func (this *DefaultFigureDecorator) Release() {
	this.DefaultFigure.Release(this.figure)
}

//forward call to contained figure
func (this *DefaultFigureDecorator) basicMoveBy(dx, dy int) {
	this.figure.basicMoveBy(dx,dy)
}

//forward call to contained figure
func (this *DefaultFigureDecorator) setBasicDisplayBox(topLeft *Point, bottomRight *Point) {
	this.figure.setBasicDisplayBox(topLeft , bottomRight)
}




type BorderDecorator struct {
	*DefaultFigureDecorator
}

func NewBorderDecorator(figure Figure) *BorderDecorator {
	this := new(BorderDecorator)
	this.DefaultFigureDecorator = NewDefaultFigureDecorator(figure)
	return this
}

func (this *BorderDecorator) Draw(g Graphics) {
	this.DefaultFigureDecorator.Draw(g)
	g.SetFGColor(Black)
	g.DrawBorderFromRect(this.GetDisplayBox())
}

func (this *BorderDecorator) Clone() Figure {
	return NewBorderDecorator(this.figure)
}
