package drifter

type VectorFunc func(x, y float64) (float64, float64)

type Field struct {
	vector VectorFunc
}

func NewField(vector VectorFunc) *Field {
	return &Field{
		vector: vector,
	}
}

func (f *Field) Acceleration(x, y float64) (float64, float64) {
	return f.vector(x, y)
}

func (f *Field) Damping(x, y float64) float64 {
	return 0.0
}

func (f *Field) Wrap(x, y float64) (float64, float64) {
	return x, y
}
