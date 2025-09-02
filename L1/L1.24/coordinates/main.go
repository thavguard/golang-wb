package coordinates

import "math"

type Point struct {
	x, y float64
}

// d = √((x₂ - x₁)² + (y₂ - y₁)²) - формула
func (p *Point) Distance(other *Point) float64 {
	x1 := p.x
	y1 := p.y

	x2 := other.x
	y2 := other.y

	return math.Sqrt(math.Pow((x2-x1), 2) + math.Pow((y2-y1), 2))
}

func NewPoint(x, y float64) *Point {
	point := new(Point)

	point.x = x
	point.y = y

	return point
}
