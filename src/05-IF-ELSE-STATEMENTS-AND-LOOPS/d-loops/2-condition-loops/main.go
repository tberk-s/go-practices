package main

import (
	"fmt"
)

func main() {
	users := []string{"taha", "berk", "ahmet", "mehmet", "ali", "veli"}

	// Loop with conditions
	for i, user := range users {
		if user == "taha" {
			fmt.Println(user, "is an ADMIN user!!!")
		} else if i == 2 {
			fmt.Println("The index is at : ", i, "and user is : ", user)
		} else {
			fmt.Println(user, "is just a regular user.")
		}
	}
	var result int
	for i := range users {
		if i == 4 {
			break
		}
		result = i + result
		fmt.Println(result)
	}
}
