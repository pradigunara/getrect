package rectangle

type Rectangle struct {
	X      int `json:"x"`
	Y      int `json:"y"`
	Width  int `json:"w" validate:"gt=0"`
	Height int `json:"h" validate:"gt=0"`
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
