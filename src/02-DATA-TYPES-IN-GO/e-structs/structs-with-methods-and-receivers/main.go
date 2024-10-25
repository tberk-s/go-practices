package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x, y, r float64
}

// We use value receiver to make calculations.
func (c Circle) area() float64 {
	return math.Pi * c.r * c.r
}

// We use Pointer receiver to change the radius.
func (c *Circle) rDoubler() {
	c.r *= 2
}

func main() {
	circle := Circle{0, 0, 5}
	result := circle.area()

	fmt.Printf("%v\n", result)  // 78.53981633974483
	fmt.Printf("%#v\n", circle) // r is 5

	circle.rDoubler()
	fmt.Printf("%#v\n", circle) // r is 10
	fmt.Println(circle.area())  // 314.1592653589793

}
