package main

import "fmt"

/*
	--------------------------------------------------------------

	In this example we don't actually change the value of the num

	The updated version lives only inside of the function.

	We need to send the memory address of num... (In second example)

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
