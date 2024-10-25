package main

import "fmt"

func main() {

	var result bool

	fmt.Printf("%t\n", result) // false, initial value

	if 2 > 1 {
		result = true // evet, 2 büyüktür 1
	}

	fmt.Printf("%v\n", result) // value’su: true
	fmt.Printf("%t\n", result) // boolean olarak value’su: true

	fmt.Printf("true && false -> %t\n", true && false)
	fmt.Printf("false && true -> %t\n", false && true)

	fmt.Printf("true || false -> %t\n", true || false)
	fmt.Printf("false || true -> %t\n", false || true)

	fmt.Printf("!true -> %t\n", !true)
	fmt.Printf("!false -> %t\n", !false)
}
