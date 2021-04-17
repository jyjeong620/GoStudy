package shapes

import "fmt"

const pi = 3.141592653589793238462643383279502884197169399375105820974944

type Circle struct {
	shape

	r float64 // 원의 반지름
}

func NewCircle(r float64) *Circle {
	return &Circle{shape{"원"}, r}
}

func (s *Circle) Area() float64 {
	return s.r * s.r * pi
}

func (s *Circle) String() string {
	return fmt.Sprintf("Type=[%T], Name=[%s], Area=[%f]", s, s.name, s.Area())
}
