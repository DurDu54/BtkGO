package interfaces

import (
	"fmt"
	"math"
)

func Demo1() {
	r := rectangle{width: 12, height: 13}
	school(r)

	c := circle{radius: 12}
	school(c)
}

type shape interface {
	area() float64
}
type rectangle struct {
	width, height float64
}
type circle struct {
	radius float64
}

func (r rectangle) area() float64 {
	return r.width * r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func school(s shape) {
	fmt.Println("seklin alani : ", s.area())
}
