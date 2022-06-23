package drifter

type Field interface {
	Acceleration(x, y float64) (float64, float64)
	Damping(x, y float64) float64
	Valid(x, y float64) bool
	Wrap(x, y float64) (float64, float64)
}

type VectorFunc func(x, y float64) (float64, float64)
type ScalarFunc func(x, y float64) float64

func VectorId(x, y float64) (float64, float64) {
	return x, y
}

func ScalarZero(x, y float64) float64 {
	return 0.0
}

// A FuncField is a field defined by a set of functions.
type FuncField struct {
	AccelFunc VectorFunc
	DampFunc  ScalarFunc
	WrapFunc  VectorFunc
}

func NewFuncField(accelFunc VectorFunc) *FuncField {
	return &FuncField{
		AccelFunc: accelFunc,
		DampFunc:  ScalarZero,
		WrapFunc:  VectorId,
	}
}

func (f *FuncField) Acceleration(x, y float64) (float64, float64) {
	return f.AccelFunc(x, y)
}

func (f *FuncField) Damping(x, y float64) float64 {
	return f.DampFunc(x, y)
}

func (f *FuncField) Valid(x, y float64) bool {
	return true
}

func (f *FuncField) Wrap(x, y float64) (float64, float64) {
	return f.WrapFunc(x, y)
}
