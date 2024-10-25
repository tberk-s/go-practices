package main

import "fmt"

/*
	--------------------------------------------------------------

	Better to not use it, we have to follow the signature of function

	Every time to see what it returns

	--------------------------------------------------------------
*/

func sum(a, b int) (result int) { // <--------------- NAMED RESULT
	result = a + b // buradaki result, (result int)'deki result

	// geri dönen şey ne? aaa pardon (result int)'deki result
	return // <--------------- NAKED RESULT
}

func main() {
	fmt.Println(sum(1, 2)) // 3
}
