package rectangle

import (
	"fmt"
)

type Rectangle struct {
	X      int
	Y      int
	Width  int
	Height int
}

func NewRectangle(x, y, w, h int) (Rectangle, error) {
	if w < 0 || h < 0 {
		return Rectangle{}, fmt.Errorf("rectangle width & height cannot be negative: w %d h %d", w, h)
	}

	return Rectangle{
		X:      x,
		Y:      y,
		Width:  w,
		Height: h,
	}, nil
}

func (r *Rectangle) IsColliding(otherRect Rectangle) bool {
	if r.X < otherRect.X+otherRect.Width &&
		r.X+r.Width > otherRect.X &&
		r.Y < otherRect.Y+otherRect.Height &&
		r.Height+r.Y > otherRect.Y {
		return true
	}

	return false
}

func (r *Rectangle) GetIntersection(otherRect Rectangle) Rectangle {
	leftX := Max(r.X, otherRect.X)
	rightX := Min(r.X+r.Width, otherRect.X+otherRect.Width)
	topY := Max(r.Y, otherRect.Y)
	bottomY := Min(r.Y+r.Height, otherRect.Y+otherRect.Height)

	return Rectangle{
		X:      leftX,
		Y:      topY,
		Width:  rightX - leftX,
		Height: bottomY - topY,
	}
}

func Max(x, y int) int {
	if x < y {
		return y
	}

	return x
}

func Min(x, y int) int {
	if x > y {
		return y
	}

	return x
}
