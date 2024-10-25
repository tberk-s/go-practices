package main

import "fmt"

func main() {
	userScores := map[string]int{
		"taha":   10,
		"berk":   20,
		"kuyruk": 30,
	}

	// Looping through a map
	for user, score := range userScores {
		fmt.Printf("%s has a score of %d\n", user, (score + 30))
	}
}
