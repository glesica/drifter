package terrain

import (
	"github.com/glesica/drifter/internal/geometry"
)

// Map defines a vector field of a particular shape and size that can
// apply acceleration and damping to bodies.
type Map interface {
	// Acceleration provides the acceleration in the X and Y directions
	// applied at the given point on the field.
	//
	// If this function is called with an invalid point its return value
	// is invalid and undefined, implementations may choose to panic.
	Acceleration(x, y float64) (float64, float64)

	// Damping is the resistance to acceleration at the given point on
	// the field as a fraction of the total (a value between 0.0 and 1.0).
	// A simulation can choose whether to use it or not.
	//
	// If this function is called with an invalid point its return value
	// is invalid and undefined, implementations may choose to panic.
	//
	// For example, if damping is 0.1, then acceleration will be reduced
	// by 0.1 or 10%.
	Damping(x, y float64) float64

	// Valid indicates whether the given point is a valid point on the field.
	Valid(x, y float64) bool

	// Wrap will wrap a point off the field to the correct point on the
	// field by applying a periodic boundary condition, or by other means
	// in the case of more complicated fields.
	//
	// If this function is called with an invalid point its return value
	// is invalid and undefined, implementations may choose to panic.
	Wrap(x, y float64) (float64, float64)
}

func NewMapFromDEM(dem *geometry.DEM) *FuncMap {
	w, h := dem.Size()
	return &FuncMap{
		AccelFunc: dem.AccelAt,
		DampFunc:  ScalarZero,
		ValidFunc: MakeBoxValidator(float64(w), float64(h)),
		WrapFunc:  MakeWrapper(float64(w), float64(h)),
	}
}
