package shapes

import "math"

func distance(x1, y1, x2, y2 float64) float64 {
	x := x2 - x1
	y := y2 - y1
	return math.Sqrt(x * x + y * y)
}

type Shape interface {
	area() float64
	perimeter() float64
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

func (r Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}

func (r Rectangle) perimeter() float64 {
	return 2 * (r.x2 - r.x1 + r.y2 - r.y1)
}

type Circle struct {
	x, y, r float64
}

func (c Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func (c Circle) perimeter() float64 {
	return 2 * math.Pi * c.r
}
