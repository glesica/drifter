package drifter

type Sim struct {
	timestamp float64
	field     Field
	traces    []*Trace
	logger    *Logger
}

func NewSim(field Field) *Sim {
	return &Sim{
		field:  field,
		logger: NewLogger(),
	}
}

func (s *Sim) AddTrace(t *Trace) {
	s.traces = append(s.traces, t)
	s.logger.Add(t.ID, s.timestamp, t.X, t.Y, t.VX, t.VY)
}

func (s *Sim) Advance(delta float64) {
	s.timestamp += delta

	for _, t := range s.traces {
		if !s.field.Valid(t.X, t.Y) {
			t.Active = false
		}

		if !t.Active {
			continue
		}

		damping := s.field.Damping(t.X, t.Y)
		t.Damp(damping, delta)

		ax, ay := s.field.Acceleration(t.X, t.Y)
		t.Advance(ax, ay, delta)

		x, y := s.field.Wrap(t.X, t.Y)
		t.MoveTo(x, y)

		s.logger.Add(t.ID, s.timestamp, t.X, t.Y, t.VX, t.VY)
	}
}

func (s *Sim) AdvanceTo(delta, ts float64) {
	for s.timestamp < ts {
		s.Advance(delta)
	}
}

func (s *Sim) History() *History {
	return s.logger.History()
}
