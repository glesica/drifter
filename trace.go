package drifter

import "math/rand"

type Trace struct {
	ID                    int
	X                     float64
	Y                     float64
	VX                    float64
	VY                    float64
	AX                    float64
	AY                    float64
	PathX                 []float64
	PathY                 []float64
	PathVX                []float64
	PathVY                []float64
	minCaptureDelta       float64
	deltaSinceLastCapture float64
}

func (t *Trace) Advance(ax, ay, delta float64) {
	xNext := t.X + t.VX*delta + 0.5*t.AX*delta*delta
	yNext := t.Y + t.VY*delta + 0.5*t.AY*delta*delta

	vxNext := t.VX + 0.5*(t.AX+ax)*delta
	vyNext := t.VY + 0.5*(t.AY+ay)*delta

	t.VX = vxNext
	t.VY = vyNext

	t.X = xNext
	t.Y = yNext

	t.AX = ax
	t.AY = ay
}

func (t *Trace) Capture(delta float64) {
	t.deltaSinceLastCapture += delta
	if t.deltaSinceLastCapture >= t.minCaptureDelta {
		t.deltaSinceLastCapture = 0.0
		t.PathX = append(t.PathX, t.X)
		t.PathY = append(t.PathY, t.Y)
		t.PathVX = append(t.PathVX, t.VX)
		t.PathVY = append(t.PathVY, t.VY)
	}
}

func (t *Trace) Damp(amount float64, delta float64) {
	t.VX *= 1.0 - amount*delta
	t.VY *= 1.0 - amount*delta
}

func (t *Trace) MoveTo(x, y float64) {
	t.X = x
	t.Y = y
}

func (t *Trace) Nudge(max float64) {
	t.AX += (rand.Float64() - 0.5) * 2.0 * max
	t.AY += (rand.Float64() - 0.5) * 2.0 * max
}

func (t *Trace) Wrap(width, height float64) {
	if t.X >= width {
		t.X -= width
	} else if t.X < 0.0 {
		t.X += width
	}

	if t.Y >= height {
		t.Y -= height
	} else if t.Y < 0.0 {
		t.Y += height
	}
}
