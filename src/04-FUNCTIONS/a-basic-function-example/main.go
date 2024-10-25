package main

import "fmt"

// "..." means it is a Variadic... we pass or receive n amount of arguments
func greet(names ...string) { // function signature
	for _, name := range names {
		fmt.Println("hello", name, "!")
	}
}

func main() {
	greet("taha") // hello taha !

	greet("taha", "berk")
	// hello taha !
	// hello berk !

	users := []string{"turbo", "max", "move"}
	// pass n amount
	greet(users...)
	// hello turbo !
	// hello max !
	// hello move !
}
