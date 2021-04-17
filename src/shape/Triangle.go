package shapes

import "fmt"

type Triangle struct {
	shape

	w float64
	h float64
}

func NewTriangle(w, h float64) *Triangle {
	return &Triangle{shape{"삼각형"}, w, h}
}

func (s *Triangle) Area() float64 {
	return s.w * s.h / 2
}

func (s *Triangle) String() string {
	return fmt.Sprintf("Type=[%T], Name=[%s], Area=[%f]", s, s.name, s.Area())
}
