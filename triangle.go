package golang_united_school_homework

import "math"

// Triangle must satisfy to Shape interface
type Triangle struct {
	Side float64
}

func (t Triangle) CalcPerimeter() float64 {
	return t.Side * 3
}

func (t Triangle) CalcArea() float64 {
	var s float64 = t.Side * 3 * 0.5
	return math.Sqrt(s * (s - t.Side) * (s - t.Side) * (s - t.Side))

}
