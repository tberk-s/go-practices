package main

import "fmt"

func main() {
	i := 0
	for {
		if i >= 5 {
			fmt.Println("reached 5, exiting...") // Exit the infinite loop when i reaches 5
			break                                // IF we dont... it goes forever and ever and ever...
		}
		fmt.Println(i)
		i++
	}
}
