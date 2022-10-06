package geometry

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDEM_AtInter(t *testing.T) {
	dem := NewDEM(3, 3)
	dem.Set(1, 1, 1.0)

	for _, c := range []struct {
		x, y, h float64
	}{
		{1.0, 1.0, 0.5},
		{1.25, 1.25, 0.75},
	} {
		h := dem.AtInter(c.x, c.y)
		assert.Equal(t, c.h, h, fmt.Sprintf("h x=%v y=%v", c.x, c.y))
	}
}

func TestDEM_SlopeAt(t *testing.T) {
	dem := NewDEM(3, 3)
	dem.Set(1, 1, 1.0)

	for _, c := range []struct {
		x, y, dx, dy float64
	}{
		// Northeast
		{1.75, 1.75, 1.0, 1.0},
		// Southeast
		{1.75, 1.25, 1.0, -1.0},
		// Southwest
		{1.25, 1.25, -1.0, -1.0},
		// Northwest
		{1.25, 1.75, -1.0, 1.0},
		// North left edge
		{0.25, 1.75, -1.0, 0.0},
		// South right edge
		{2.75, 1.25, 1.0, 0.0},
		// West top edge
		{1.25, 2.75, 0.0, 1.0},
		// East bottom edge
		{1.75, 0.25, 0.0, -1.0},
	} {
		dx, dy := dem.AccelAt(c.x, c.y)
		assert.Equal(t, c.dx, dx, fmt.Sprintf("dx x=%v y=%v", c.x, c.y))
		assert.Equal(t, c.dy, dy, fmt.Sprintf("dy x=%v y=%v", c.x, c.y))
	}
}

func TestDEM_discretize(t *testing.T) {
	dem := NewDEM(3, 3)

	var x, y float64

	x, y = dem.discretize(0.0, 0.5)
	assert.Equal(t, 0.0, x)
	assert.Equal(t, 0.0, y)

	x, y = dem.discretize(3.0, 2.5)
	assert.Equal(t, 2.0, x)
	assert.Equal(t, 2.0, y)
}
