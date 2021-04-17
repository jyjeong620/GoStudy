package shapes

import "fmt"

// <1>
type shape struct {
	name string
}

// <2>
func (*shape) Area() float64 {
	return 0
}

// <3>
func (s *shape) String() string {
	return fmt.Sprintf("Type=[%T], Name=[%s], Area=[%f]", s, s.name, s.Area())
}
