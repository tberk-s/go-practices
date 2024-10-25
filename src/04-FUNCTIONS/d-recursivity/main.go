package main

import "fmt"

/*
	--------------------------------------------------------------

	The function can returns itself and these functions are called:

	RECURSIVE functions.. situation is just called "RECURSIVITY"

	We can see a factorial calculation example

	--------------------------------------------------------------
*/

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1) // kendini çağırdı.
}

func main() {
	fmt.Println(fact(3)) // 6
}
