package drifter

import (
	"encoding/json"
	"io"
	"math"
)

// A Logger tracks some number of Traces using arbitrary storage and can
// write its contents to a file for later replace or extension.
type Logger struct {
	history *History
}

func NewLogger() *Logger {
	return &Logger{
		history: &History{
			Traces: make(map[int]*TraceHistory),
		},
	}
}

func (l *Logger) Add(id int, ts, x, y, vx, vy float64) {
	frame := &TraceFrame{ts, id, x, y, vx, vy}
	trace, found := l.history.Traces[id]
	if !found {
		trace = &TraceHistory{id, []*TraceFrame{}}
		l.history.Traces[id] = trace
	}
	trace.Frames = append(trace.Frames, frame)

	if ts > l.history.LastTimestamp {
		l.history.LastTimestamp = ts
	}
}

func (l *Logger) History() *History {
	return l.history
}

func (l *Logger) Load(r io.Reader) error {
	decoder := json.NewDecoder(r)
	l.history = &History{}
	return decoder.Decode(l.history)
}

func (l *Logger) Save(w io.Writer) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(l.history)
}

type TraceFrame struct {
	Timestamp float64 `json:"timestamp"`
	ID        int     `json:"id"`
	X         float64 `json:"x"`
	Y         float64 `json:"y"`
	VX        float64 `json:"vx"`
	VY        float64 `json:"vy"`
}

func (t *TraceFrame) Speed() float64 {
	return math.Sqrt(t.VX*t.VX + t.VY*t.VY)
}

type TraceHistory struct {
	ID     int `json:"id"`
	Frames []*TraceFrame
}

func (t *TraceHistory) FirstX() float64 {
	return t.Frames[0].X
}

func (t *TraceHistory) FirstY() float64 {
	return t.Frames[0].Y
}

func (t *TraceHistory) LastX() float64 {
	return t.Frames[len(t.Frames)-1].X
}

func (t *TraceHistory) LastY() float64 {
	return t.Frames[len(t.Frames)-1].Y
}

func (t *TraceHistory) Length() int {
	return len(t.Frames)
}

type History struct {
	Traces        map[int]*TraceHistory `json:"traces"`
	LastTimestamp float64               `json:"last-timestamp"`
}

func (h *History) SpeedBounds() (float64, float64) {
	maxSpeed := 0.0
	minSpeed := math.MaxFloat64
	for _, h := range h.Traces {
		for _, t := range h.Frames {
			s := t.Speed()
			if s > maxSpeed {
				maxSpeed = s
			}
			if s < minSpeed {
				minSpeed = s
			}
		}
	}
	return minSpeed, maxSpeed
}
