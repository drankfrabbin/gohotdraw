package gohotdraw

type Locator interface {
	Locate(owner Figure) *Point
}

type DefaultLocator struct{}

type RelativeLocator struct {
	*DefaultLocator
	relativeX float
	relativeY float
}

func NewDefaultRelativeLocator() *RelativeLocator {
	return NewRelativeLocator(0.0, 0.0)
}

func NewRelativeLocator(relativeX, relativeY float) *RelativeLocator {
	return &RelativeLocator{relativeX: relativeX, relativeY: relativeY}
}

func (this *RelativeLocator) Locate(owner Figure) *Point {
	r := owner.GetDisplayBox()
	xValue := r.X + (int)((float)(r.Width)*this.relativeX)
	yValue := r.Y + (int)((float)(r.Height)*this.relativeY)
	return &Point{xValue, yValue}
}

func CreateNorthLocator() Locator {
	return NewRelativeLocator(0.5, 0.0)
}

func CreateWestLocator() Locator {
	return NewRelativeLocator(0.0, 0.5)
}

func CreateSouthLocator() Locator {
	return NewRelativeLocator(0.5, 1.0)
}

func CreateEastLocator() Locator {
	return NewRelativeLocator(1.0, 0.5)
}

func CreateNorthEastLocator() Locator {
	return NewRelativeLocator(1.0, 0.0)
}

func CreateNorthWestLocator() Locator {
	return NewRelativeLocator(0.0, 0.0)
}

func CreateSouthEastLocator() Locator {
	return NewRelativeLocator(1.0, 1.0)
}

func CreateSouthWestLocator() Locator {
	return NewRelativeLocator(0.0, 1.0)
}

func CreateCenterLocator() Locator {
	return NewRelativeLocator(0.5, 0.5)
}
