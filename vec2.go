package vec

import (
	"fmt"
)

// NewVec2 creates a new Vec2
func NewVec2(x, y float64) *Vec2 {
	return &Vec2{
		X: x,
		Y: y,
	}
}

// Vec2 is a vector containing 2 floats
type Vec2 struct {
	X float64
	Y float64
}

// String converts the Vec2 to a string
func (vec *Vec2) String() string {
	return "Vec2{" + fmt.Sprint(vec.X) + ", " + fmt.Sprint(vec.Y) + "}"
}
