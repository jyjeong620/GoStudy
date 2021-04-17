package shapes

import "fmt"

type Rectangle struct {
	// <1>
	shape

	// <2>
	w float64
	h float64
}

// <3>
func NewRectangle(w, h float64) *Rectangle {
	return &Rectangle{shape{"사각형"}, w, h}
}

// <4>
func (s *Rectangle) Area() float64 {
	return s.w * s.h
}

func (s *Rectangle) String() string {
	return fmt.Sprintf("Type=[%T], Name=[%s], Area=[%f]", s, s.name, s.Area())
}
