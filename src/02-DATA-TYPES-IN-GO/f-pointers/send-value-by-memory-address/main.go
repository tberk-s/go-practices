package main

import (
	"fmt"
	"log"
)

/*
	--------------------------------------------------------------

	In this example we change the value by sending memory address!

	--------------------------------------------------------------
*/

func increment(n *int) error {
	if n == nil {
		return fmt.Errorf("n nil geldi")
	}
	fmt.Println("increment", "gelen n", *n, "memory adresi:", n)

	*n++

	fmt.Println("increment", "arttıktan sonra n", *n)
	return nil
}

func main() {
	num := 10
	fmt.Println("num:", num, "memory adresi:", &num)

	fmt.Printf("num with increment func: %v\n", increment(&num))

	fmt.Println("new num!!:", num, "new memory address!!: ", &num) //memory address same!

	if err := increment(nil); err != nil {
		log.Fatal(err)
	}
}
