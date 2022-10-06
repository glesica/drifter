package history

// Trace records the current and historical position and
// movement data for an entity within a simulation.
type Trace struct {
	x, y    float64
	vx, vy  float64
	ax, ay  float64
	history []float64
}

func NewTrace(x, y float64) *Trace {
	return &Trace{
		x: x,
		y: y,
	}
}

// Update applies the given acceleration, accounting for the
// given time delta, to the trace and updates its position. The
// Velocity Verlet method, which is highly numerically stable,
// is used to update the position and velocity of the entity.
func (t *Trace) Update(ax, ay, dt float64) {
	x1 := t.x + t.vx*dt + 0.5*t.ax*dt*dt
	y1 := t.y + t.vy*dt + 0.5*t.ay*dt*dt

	t.vx = t.vx + 0.5*(t.ax+ax)*dt
	t.vy = t.vy + 0.5*(t.ay+ay)*dt

	t.x = x1
	t.y = y1

	t.ax = ax
	t.ay = ay

	t.history = append(t.history, t.x, t.y, t.vx, t.vy)
}

func (t *Trace) X() float64 {
	return t.x
}

func (t *Trace) Y() float64 {
	return t.y
}

func (t *Trace) VX() float64 {
	return t.vx
}

func (t *Trace) VY() float64 {
	return t.vy
}

// History produces an iterator from the current state of
// the trace. Note that the iterator won't necessarily
// yield any snapshots because the trace may be empty.
func (t *Trace) History() *Iterator {
	return &Iterator{
		history: t.history,
	}
}

type Iterator struct {
	index   int
	history []float64
}

// Next produces the next snapshot in the iterator. The first
// returned value indicates whether the iterator was able to
// produce another snapshot (if false, then no snapshot was
// returned).
func (i *Iterator) Next() (bool, Snapshot) {
	if i.index+3 >= len(i.history) {
		return false, nil
	}

	s := snapshot(i.history[i.index : i.index+4])
	i.index += 4

	return true, s
}

type Snapshot interface {
	X() float64
	Y() float64
	VX() float64
	VY() float64
}

type snapshot []float64

func (s snapshot) X() float64 {
	return s[0]
}

func (s snapshot) Y() float64 {
	return s[1]
}

func (s snapshot) VX() float64 {
	return s[2]
}

func (s snapshot) VY() float64 {
	return s[3]
}
