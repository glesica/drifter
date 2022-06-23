package drifter

import (
	"math"
	"testing"
)

const EPSILON = 0.0001

func TestHeightField_Acceleration(t *testing.T) {
	f := NewHeightField([]float64{
		1, 1, 1, 1, 1,
		1, 2, 2, 2, 1,
		1, 2, 3, 2, 1,
		1, 2, 2, 2, 1,
		1, 1, 1, 1, 1,
	}, 5)

	var ax, ay float64

	for _, c := range []struct {
		x   float64
		y   float64
		eax float64
		eay float64
	}{
		{x: -0.5, y: 0.0, eax: -1.0, eay: 0.5},
		{x: 0.5, y: 0.0, eax: 1.0, eay: 0.5},
		{x: 0.0, y: -0.5, eax: 0.5, eay: -1.0},
		{x: 0.0, y: 0.5, eax: 0.5, eay: 1.0},
	} {
		ax, ay = f.Acceleration(c.x, c.y)

		if !roughly(ax, c.eax) {
			t.Errorf("(%f, %f) expected ax=%f, got %f", c.x, c.y, c.eax, ax)
		}
		if !roughly(ay, c.eay) {
			t.Errorf("(%f, %f) expected ay=%f, got %f", c.x, c.y, c.eay, ay)
		}
	}
}

func TestHeightField_Height(t *testing.T) {
	f := NewHeightField([]float64{
		1, 1, 1, 1, 1,
		1, 2, 2, 2, 1,
		1, 2, 3, 2, 1,
		1, 2, 2, 2, 1,
		1, 1, 1, 1, 1,
	}, 5)

	var h float64

	h = f.Height(2, 2)
	if h != 3.0 {
		t.Fatalf("expected height 3.0, got %f", h)
	}

	h = f.Height(0, 0)
	if h != 1.0 {
		t.Fatalf("expected height 1.0, got %f", h)
	}

	h = f.Height(4, 4)
	if h != 1.0 {
		t.Fatalf("expected height 1.0, got %f", h)
	}
}

func roughly(act, exp float64) bool {
	return math.Abs(exp-act) < EPSILON
}
