package geometry

// Normal returns the normal vector of the plane defined by the
// three given points.
func Normal(pt0, pt1, pt2 Vec) Vec {
	vec0 := pt1.Minus(pt0)
	vec1 := pt2.Minus(pt0)

	return Vec{
		X: vec0.Y*vec1.Z - vec0.Z*vec1.Y,
		Y: vec0.Z*vec1.X - vec0.X*vec1.Z,
		Z: vec0.X*vec1.Y - vec0.Y*vec1.X,
	}
}

// Slopes produces the slope of the plane defined by the three
// given points in the X and Y dimensions, respectively.
func Slopes(pt0, pt1, pt2 Vec) (dx float64, dy float64) {
	norm := Normal(pt0, pt1, pt2)
	dx = -norm.X / norm.Z
	dy = -norm.Y / norm.Z
	return
}
