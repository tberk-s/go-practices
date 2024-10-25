package main

import "fmt"

func main() {
switchStatement:
	switch 1 {
	case 1:
		fmt.Println("1") // This will print
		for i := 0; i < 5; i++ {
			if i == 2 { // DOES NOT MATTER JUST WROTE THIS TO FIX GOD DAMN LINTER
				break switchStatement
			} // Break out of the switch statement immediately
		}
		fmt.Println("2") // This will not execute
	case 2:
	default:
		fmt.Println("default case...") // This will not execute
	}
	fmt.Println("3") // This will print

	// Output:
	// 1
	// 3

	i := 0
Start:
	fmt.Println(i) // Print the current value of i
	if i > 2 {
		goto End // If i is greater than 2, jump to End
	}
	i += 1     // Increment i
	goto Start // Jump back to Start

End:

	//0
	//1
	//2
	//3

}
