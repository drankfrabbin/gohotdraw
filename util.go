package gohotdraw

import (
	_ "fmt"
	"math"
)

type Dimension struct {
	Width  int
	Height int
}

type Point struct {
	X int
	Y int
}

type Rectangle struct {
	X      int
	Y      int
	Width  int
	Height int
}

func NewRectangle() *Rectangle {
	return &Rectangle{0, 0, 0, 0}
}

func NewRectangleFromRect(r *Rectangle) *Rectangle {
	return &Rectangle{r.X, r.Y, r.Width, r.Height}
}

func NewRectangleFromPoint(p *Point) *Rectangle {
	return &Rectangle{p.X, p.Y, 0, 0}
}

func NewRectangleFromPoints(p1, p2 *Point) *Rectangle {
	rect := NewRectangleFromPoint(p1)
	rect.AddPoint(p2)
	return rect
}

func (this *Rectangle) AddPoint(p *Point) {
	this.Add(p.X, p.Y)
}

func (this *Rectangle) Add(newX, newY int) {
	if (this.Width | this.Height) < 0 {
		this.SetBounds(newX, newY, 0, 0)
		return
	}
	x1 := this.X
	y1 := this.Y
	x2 := this.Width
	y2 := this.Height
	x2 += x1
	y2 += y1
	if x1 > newX {
		x1 = newX
	}
	if y1 > newY {
		y1 = newY
	}
	if x2 < newX {
		x2 = newX
	}
	if y2 < newY {
		y2 = newY
	}
	x2 -= x1
	y2 -= y1
	if x2 > math.MaxInt32 {
		x2 = math.MaxInt32
	}
	if y2 > math.MaxInt32 {
		y2 = math.MaxInt32
	}
	this.SetBounds(x1, y1, x2, y2)
}

func (this *Rectangle) SetBounds(x, y, width, height int) {
	this.X = x
	this.Y = y
	this.Width = width
	this.Height = height
}

func (this *Rectangle) Translate(dx, dy int) {
	oldv := this.X
	newv := oldv + dx
	if dx < 0 {
		// moving leftward
		if newv > oldv {
			// negative overflow
			// Only adjust width if it was valid (>= 0).
			if this.Width >= 0 {
				// The right edge is now conceptually at
				// newv+width, but we may move newv to prevent
				// overflow.  But we want the right edge to
				// remain at its new location in spite of the
				// clipping.  Think of the following adjustment
				// conceptually the same as:
				// width += newv newv = MIN_VALUE width -= newv
				this.Width += newv - math.MinInt32
				// width may go negative if the right edge went past
				// MIN_VALUE, but it cannot overflow since it cannot
				// have moved more than MIN_VALUE and any non-negative
				// number + MIN_VALUE does not overflow.
			}
			newv = math.MinInt32
		}
	} else {
		// moving rightward (or staying still)
		if newv < oldv {
			// positive overflow
			if this.Width >= 0 {
				// Conceptually the same as:
				// width += newv newv = MAX_VALUE width -= newv
				this.Width += newv - math.MaxInt32
				// With large widths and large displacements
				// we may overflow so we need to check it.
				if this.Width < 0 {
					this.Width = math.MaxInt32
				}
			}
			newv = math.MaxInt32
		}
	}
	this.X = newv

	oldv = this.Y
	newv = oldv + dy
	if dy < 0 {
		// moving upward
		if newv > oldv {
			// negative overflow
			if this.Height >= 0 {
				this.Height += newv - math.MinInt32
				// See above comment about no overflow in this case
			}
			newv = math.MinInt32
		}
	} else {
		// moving downward (or staying still)
		if newv < oldv {
			// positive overflow
			if this.Height >= 0 {
				this.Height += newv - math.MaxInt32
				if this.Height < 0 {
					this.Height = math.MaxInt32
				}
			}
			newv = math.MaxInt32
		}
	}
	this.Y = newv
}

func (this *Rectangle) Grow(h, v int) {
	x0 := this.X
	y0 := this.Y
	x1 := this.Width
	y1 := this.Height
	x1 += x0
	y1 += y0

	x0 -= h
	y0 -= v
	x1 += h
	y1 += v

	if x1 < x0 {
		// Non-existant in X direction
		// Final width must remain negative so subtract x0 before
		// it is clipped so that we avoid the risk that the clipping
		// of x0 will reverse the ordering of x0 and x1.
		x1 -= x0
		if x1 < math.MinInt32 {
			x1 = math.MinInt32
		}
		if x0 < math.MinInt32 {
			x0 = math.MinInt32
		} else if x0 > math.MaxInt32 {
			x0 = math.MaxInt32
		}
	} else { // (x1 >= x0)
		// Clip x0 before we subtract it from x1 in case the clipping
		// affects the representable area of the rectangle.
		if x0 < math.MinInt32 {
			x0 = math.MinInt32
		} else if x0 > math.MaxInt32 {
			x0 = math.MaxInt32
		}
		x1 -= x0
		// The only way x1 can be negative now is if we clipped
		// x0 against MIN and x1 is less than MIN - in which case
		// we want to leave the width negative since the result
		// did not intersect the representable area.
		if x1 < math.MinInt32 {
			x1 = math.MinInt32
		} else if x1 > math.MaxInt32 {
			x1 = math.MaxInt32
		}
	}

	if y1 < y0 {
		// Non-existant in Y direction
		y1 -= y0
		if y1 < math.MinInt32 {
			y1 = math.MinInt32
		}
		if y0 < math.MinInt32 {
			y0 = math.MinInt32
		} else if y0 > math.MaxInt32 {
			y0 = math.MaxInt32
		}
	} else { // (y1 >= y0)
		if y0 < math.MinInt32 {
			y0 = math.MinInt32
		} else if y0 > math.MaxInt32 {
			y0 = math.MaxInt32
		}
		y1 -= y0
		if y1 < math.MinInt32 {
			y1 = math.MinInt32
		} else if y1 > math.MaxInt32 {
			y1 = math.MaxInt32
		}
	}
	this.SetBounds(x0, y0, x1, y1)
}

func (this *Rectangle) Union(r *Rectangle) *Rectangle {
	tx2 := this.Width
	ty2 := this.Height
	if (tx2 | ty2) < 0 {
		// This rectangle has negative dimensions...
		// If r has non-negative dimensions then it is the answer.
		// If r is non-existant (has a negative dimension), then both
		// are non-existant and we can return any non-existant rectangle
		// as an answer.  Thus, returning r meets that criterion.
		// Either way, r is our answer.
		return NewRectangleFromRect(r)
	}
	rx2 := r.Width
	ry2 := r.Height
	if (rx2 | ry2) < 0 {
		return NewRectangleFromRect(this)
	}
	tx1 := this.X
	ty1 := this.Y
	tx2 += tx1
	ty2 += ty1
	rx1 := r.X
	ry1 := r.Y
	rx2 += rx1
	ry2 += ry1
	if tx1 > rx1 {
		tx1 = rx1
	}
	if ty1 > ry1 {
		ty1 = ry1
	}
	if tx2 < rx2 {
		tx2 = rx2
	}
	if ty2 < ry2 {
		ty2 = ry2
	}
	tx2 -= tx1
	ty2 -= ty1
	// tx2,ty2 will never underflow since both original rectangles
	// were already proven to be non-empty
	// they might overflow, though...
	if tx2 > math.MaxInt32 {
		tx2 = math.MaxInt32
	}
	if ty2 > math.MaxInt32 {
		ty2 = math.MaxInt32
	}
	return &Rectangle{tx1, ty1, tx2, ty2}
}

func (this *Rectangle) Contains(X, Y int) bool {
	w := this.Width
	h := this.Height
	if (w | h) < 0 {
		// At least one of the dimensions is negative...
		return false
	}
	// Note: if either dimension is zero, tests below must return false...
	x := this.X
	y := this.Y
	if X < x || Y < y {
		return false
	}
	w += x
	h += y
	//    overflow || intersect
	return ((w < x || w > X) &&
		(h < y || h > Y))
}

func (this *Rectangle) ContainsRect(rect *Rectangle) bool {
	return (
		rect.X >= this.X && rect.Y >= this.Y && 
		(rect.X+int(math.Fmax(0, float64(rect.Width)))) <= this.X+int(math.Fmax(0, float64(this.Width))) && 
		(rect.Y+int(math.Fmax(0, float64(rect.Height)))) <= this.Y+int(math.Fmax(0, float64(this.Height))))
}

func (this *Rectangle) ContainsPoint(point *Point) bool {
	return this.Contains(point.X, point.Y)
}

func (this *Rectangle) IsEmpty() bool {
	return (this.Width <= 0) || (this.Height <= 0)
}
