
package square

type Point struct {
	X, Y uint
}

type Square struct {
	Start Point
	A     uint
}

func (s *Square) End() Point {
	return Point{s.Start.X + s.A, s.Start.Y - s.A}
}

func (s *Square) Area() uint {
	return s.A * s.A
}

func (s *Square) Perimeter() uint {
	return s.A * 4
}
