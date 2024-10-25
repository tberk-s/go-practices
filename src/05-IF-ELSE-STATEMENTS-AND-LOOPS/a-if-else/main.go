package main

import (
	"errors"
	"fmt"
	"math/rand"
)

// Function that returns an error
func mayFail() error {
	return errors.New("something went wrong")
}
func main() {
	number := 0

	// If-Else If-Else statement
	if number > 5 {
		fmt.Println("Number is greater than 5")
	} else if number < 5 {
		fmt.Println("Number is less than 5")
	} else {
		fmt.Println("Number is exactly 5")
	}

	/*
		-------------------------------------------------------------

		Short If Statement

		--------------------------------------------------------------
	*/

	if randomNum := rand.Intn(10); randomNum > 5 {
		fmt.Println("Random number is greater than 5:", randomNum)
	} else {
		fmt.Println("Random number is 5 or less:", randomNum)
	}

	/*
		-------------------------------------------------------------

		Can be used for error handling

		--------------------------------------------------------------
	*/

	if err := mayFail(); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Function succeeded!")
	}

}
