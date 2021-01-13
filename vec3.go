package vec

import "fmt"

// Vec3 is a vector containing 3 floats
type Vec3 struct {
	X float64
	Y float64
	Z float64
}

// NewVec3 creates a new Vec3
func NewVec3(x, y, z float64) *Vec3 {
	return &Vec3{
		X: x,
		Y: y,
		Z: z,
	}
}

// String converts the Vec3 to a string
func (vec *Vec3) String() string {
	return "Vec3{" + fmt.Sprint(vec.X) + ", " + fmt.Sprint(vec.Y) + ", " + fmt.Sprint(vec.Z) + "}"
}
