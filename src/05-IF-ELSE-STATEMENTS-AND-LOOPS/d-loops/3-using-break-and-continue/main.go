package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i == 5 {
			fmt.Println("Breaking the loop at 5")
			break // Exit the loop
		}
		fmt.Println(i) // Did not print after 5, "break" kills the for loop!
	}

	fmt.Println("Continuing with the next loop")

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue // Skip even numbers NEXT!!
		}
		fmt.Println(i) // Will print only odd numbers
	}
}
