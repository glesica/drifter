package geometry

type Vec struct {
	X, Y, Z float64
}

func (v Vec) Minus(other Vec) Vec {
	return Vec{
		X: v.X - other.X,
		Y: v.Y - other.Y,
		Z: v.Z - other.Z,
	}
}
