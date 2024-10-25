// We mostly use switch/case if we need to control more than one thing... Better to use switch case instead of if-else!!!!!!!!!

package main

import (
	"fmt"
	"runtime"
)

func main() {
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("Mac OS Hipster, your os is", os)
		// Mac OS Hipster, your os is darwin
	case "linux":
		fmt.Println("Linux Geek, your os is", os)
	default:
		fmt.Println("Other:", os)
	}
	//fmt.Println("os was", os)
	// Lives inside of the switch
	// undefined: os

	// We can check like this too!
	number := 42
	switch {
	case number < 42:
		fmt.Println("küçük")
	case number == 42:
		fmt.Println("eşit")
	case number > 42:
		fmt.Println("büyük")
	}

}
