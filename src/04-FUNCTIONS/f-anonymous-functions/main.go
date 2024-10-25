package main

import "fmt"

/*
	-------------------------------------------------------------

	// anonim fonksiyon
	func() {
   		fmt.Println("anonymous")
	}()

	--------------------------------------------------------------
*/

func main() {
	counter := func() func(int) {
		// Return an anonymous function that counts up to the given number
		return func(n int) {
			for i := 1; i <= n; i++ {
				fmt.Println(i) // Print the current count
			}
		}
	}()

	// Call the counter multiple times
	counter(4)

	// Create another counter
	anotherCounter := func() func() int {
		count := 10 // Start from a different initial value

		return func() int {
			count++
			return count
		}
	}()

	// Call the new counter
	fmt.Println(anotherCounter()) // Output: 11
	fmt.Println(anotherCounter()) // Output: 12
}
