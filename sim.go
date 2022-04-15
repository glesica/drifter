package drifter

type Sim struct {
	field  *Field
	traces []*Trace
}

func NewSim(field *Field) *Sim {
	return &Sim{
		field: field,
	}
}

func (s *Sim) AddTrace(t *Trace) {
	s.traces = append(s.traces, t)
}

func (s *Sim) Advance(delta float64) {
	for _, trace := range s.traces {
		damping := s.field.Damping(trace.X, trace.Y)
		trace.Damp(damping, delta)

		ax, ay := s.field.Acceleration(trace.X, trace.Y)
		trace.Advance(ax, ay, delta)

		x, y := s.field.Wrap(trace.X, trace.Y)
		trace.MoveTo(x, y)

		trace.Capture(delta)
	}
}
