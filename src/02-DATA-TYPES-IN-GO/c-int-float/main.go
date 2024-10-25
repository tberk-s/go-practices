package main

import "fmt"

func main() {
	var x int = 125
	y := -125

	var a float64 = 11.12
	b := -11.12

	// COMPLEX NUMBER DEFINITION:
	// c := 3 + 4i  // complex128

	fmt.Printf("X : Type: %T, value: %[1]v\n", x)
	fmt.Printf("Y : Type: %T, value: %[1]v\n", y)

	fmt.Printf("A : Type: %T, value: %[1]v\n", a)
	fmt.Printf("B : Type: %T, value: %[1]v\n", b)

	fmt.Printf("Sum of X and A : %v\n", (float64(x) + a))
	fmt.Printf("Sum of Y and B : %v\n", (float64(y) + b))

}
