package main

import "fmt"

/*
	--------------------------------------------------------------

	We can use `new` function to initialize and preallocate the struct in the memory

	new function returns the memory address.

	--------------------------------------------------------------
*/

func increment(n int) {
	fmt.Println("increment", "gelen n", n, "memory adresi:", &n)
	// increment gelen n 10 memory adresi: 0x140000a4030

	n++

	fmt.Println("increment", "arttıktan sonra n", n)
	// increment arttıktan sonra n 11
}

func main() {
	num := 10
	fmt.Println("num:", num, "memory adresi:", &num)
	// num: 10 memory adresi: 0x1400011a020

	increment(num) // num'ın kopyası gönderilir
	fmt.Println("num:", num, "memory adresi:", &num)
	// num: 10 memory adresi: 0x1400011a020
}
