package sim

import (
	"github.com/glesica/drifter/internal/art"
	"github.com/glesica/drifter/internal/history"
	"github.com/glesica/drifter/internal/terrain"
)

type Driver struct {
	field  terrain.Map
	traces []*history.Trace
}

func NewDriver(field terrain.Map, traces []*history.Trace) *Driver {
	return &Driver{
		field:  field,
		traces: traces,
	}
}

func (d *Driver) Render(drawer art.Drawer) {
	for _, trace := range d.traces {
		drawer(trace.History())
	}
}

func (d *Driver) Update(dt float64) {
	for _, trace := range d.traces {
		ax, ay := d.field.Acceleration(trace.X(), trace.Y())
		trace.Update(ax, ay, dt)
	}
}