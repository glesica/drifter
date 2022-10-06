package geometry

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNormal(t *testing.T) {
	var norm Vec

	p0 := Vec{0, 0, 0}
	p1 := Vec{0, 1, 0}
	p2 := Vec{1, 0, 0}

	norm = Normal(p0, p1, p2)
	assert.Equal(t, 0.0, norm.X)
	assert.Equal(t, 0.0, norm.Y)
	assert.Equal(t, -1.0, norm.Z)

	norm = Normal(p0, p2, p1)
	assert.Equal(t, 0.0, norm.X)
	assert.Equal(t, 0.0, norm.Y)
	assert.Equal(t, 1.0, norm.Z)
}

func TestSlopes(t *testing.T) {
	var dx, dy float64

	p0 := Vec{0, 0, 0}
	p1 := Vec{0, 1, 0}
	p2 := Vec{1, 0, 0}

	dx, dy = Slopes(p0, p1, p2)
	assert.Equal(t, 0.0, dx)
	assert.Equal(t, 0.0, dy)

	p0.Z = 1.0
	dx, dy = Slopes(p0, p1, p2)
	assert.Equal(t, -1.0, dx)
	assert.Equal(t, -1.0, dy)
}
