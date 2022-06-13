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
	var sP float64 = t.Side * 3 * 0.5
	return math.Sqrt(sP * (sP * t.Side) * (sP * t.Side) * (sP * t.Side))
}
