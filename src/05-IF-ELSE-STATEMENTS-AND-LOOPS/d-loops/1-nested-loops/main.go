package main

import "fmt"

func main() {
	users := []string{"taha", "berk"}
	scores := []int{10, 20, 30, 40, 50}

	// Nested loop example
	for _, user := range users {
		for _, score := range scores {
			fmt.Printf("%s scored %d points\n", user, score)
		}
	}
}
