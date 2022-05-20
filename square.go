package square

type Point struct {
	x, y int
}

type Square struct {
	start Point
	a     uint
}

func (s *Square) End() Point {
	return Point{uint(s.start.x) + s.a, uint(s.start.y) - s.a}
}

func (s *Square) Area() uint {
	return s.a * s.a
}

func (s *Square) Perimeter() uint {
	return s.a * 4
}
