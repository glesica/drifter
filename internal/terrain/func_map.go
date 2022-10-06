package terrain

import (
	"math"
)

type VectorFunc func(x, y float64) (float64, float64)

type ScalarFunc func(x, y float64) float64

type BoolFunc func(x, y float64) bool

func TrueFunc(x, y float64) bool {
	return true
}

// VectorId is the vector identity function.
func VectorId(x, y float64) (float64, float64) {
	return x, y
}

// VectorZero simply returns a zero vector for all inputs.
func VectorZero(x, y float64) (float64, float64) {
	return 0.0, 0.0
}

// ScalarZero returns a scalar zero for all inputs.
func ScalarZero(x, y float64) float64 {
	return 0.0
}

func MakeBoxValidator(w, h float64) BoolFunc {
	return func(x, y float64) bool {
		return x >= 0.0 && x <= w && y >= 0.0 && y <= h
	}
}

func MakeWrapper(w, h float64) VectorFunc {
	return func(x, y float64) (wx float64, wy float64) {
		wx = x
		wy = y

		if x < 0.0 {
			wx = w + (math.Mod(x, w))
		}

		if y < 0.0 {
			wy = h + (math.Mod(y, h))
		}

		if x > w {
			wx = math.Mod(x, w)
		}

		if y > h {
			wy = math.Mod(y, h)
		}

		return
	}
}

// A FuncMap is a field defined by a set of functions.
type FuncMap struct {
	AccelFunc VectorFunc
	DampFunc  ScalarFunc
	ValidFunc BoolFunc
	WrapFunc  VectorFunc
}

func NewFuncMap() *FuncMap {
	return &FuncMap{
		AccelFunc: VectorZero,
		DampFunc:  ScalarZero,
		ValidFunc: TrueFunc,
		WrapFunc:  VectorId,
	}
}

func (f *FuncMap) Acceleration(x, y float64) (float64, float64) {
	return f.AccelFunc(x, y)
}

func (f *FuncMap) Damping(x, y float64) float64 {
	return f.DampFunc(x, y)
}

func (f *FuncMap) Valid(x, y float64) bool {
	return f.ValidFunc(x, y)
}

func (f *FuncMap) Wrap(x, y float64) (float64, float64) {
	return f.WrapFunc(x, y)
}
