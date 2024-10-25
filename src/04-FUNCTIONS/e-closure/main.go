package main

import "fmt"

// fib function returns a function that generates Fibonacci numbers
func fib() func() int {
	a, b := 0, 1 // Initialize the first two Fibonacci numbers
	return func() int {
		next := a     // Capture the current Fibonacci number
		a, b = b, a+b // Update to the next Fibonacci numbers
		return next   // Return the current Fibonacci number
	}
}

func main() {
	f := fib() // Get the Fibonacci generator function

	for i := 0; i < 10; i++ {
		fmt.Println(f()) // Generate and print the next Fibonacci number
	}
}
