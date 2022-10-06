package terrain

import (
	"github.com/glesica/drifter/internal/history"
	"math/rand"
)

func RandLayout(count int, w, h float64) []*history.Trace {
	traces := make([]*history.Trace, count)
	for i := 0; i < count; i++ {
		x := rand.Float64() * w
		y := rand.Float64() * h
		traces[i] = history.NewTrace(x, y)
	}

	return traces
}
