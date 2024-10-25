package main

import "fmt"

func main() {
	var a [3]int

	fmt.Printf("Values of array A :  %v %[1]T\n", a) // [0 0 0] [3]int

	a[0] = 1
	a[1] = 2
	a[2] = 3

	fmt.Printf("Values of array A : %v\n", a) // [1 2 3]

	fmt.Printf("%v\n", len(a))

	x := [3]int{10, 20, 30} // short declaration...

	fmt.Printf("Values of array X :  %v %[1]T\n\n", x)

	// a := [...]int{10, 20, 30} We let compiler to decide the len of array

	var fruits = [...]string{"apple", "melon"}

	otherFruits := fruits

	otherFruits[0] = "banana"

	fmt.Println("otherFruits", otherFruits, len(otherFruits), cap(otherFruits))
	// otherFruits [banana melon] 2 2

	fmt.Println("fruits", fruits, len(fruits), cap(fruits))
	// fruits [apple melon] 2 2

	// & ile değişkenin hafızadaki işaret (point) ettiği yeri alırız.
	fmt.Printf("fruits: %p\n", &fruits)           // fruits: 0x10276d3e0
	fmt.Printf("otherFruits: %p\n", &otherFruits) // otherFruits: 0x14000060020

}
