package main

import (
	"fmt"
	"strings"
)

func main() {

	var name string = "Taha\n"
	var str string = "İı"
	var raw_string string = `"Raw" 
							string in Go\n`

	// Non-ASCII literal. Go source is UTF-8
	symbol := 'Σ' //rune type, rune is an alias for int32, holds a unicode code point.

	fmt.Println(name)       // Will process "\n" as new line
	fmt.Println(raw_string) // `\n` can be seen on output
	fmt.Println(symbol)     // 932 is byte representation..

	fmt.Printf("Length of name : %v\n", len(name))
	fmt.Printf("Length of raw_string : %v\n", len(raw_string))
	fmt.Printf("Length of string with Turkish characters : %v\n", len(str)) // It says 4... why??

	// symbol is actually rune...
	// fmt.Println(len(symbol)) //invalid argument: symbol (variable of type rune)

	fmt.Println([]byte(str)) // It has 4 bytes, 2 for each because these are Turkish characters
	fmt.Println(raw_string[6:])

	// String Concatenation

	hello := "Hello"
	taha := "taha"

	result := hello + " " + taha + "!!!"

	fmt.Println(result)

	// Compare strings

	fruit1 := "apple"
	fruit2 := "pear"

	fmt.Println(fruit1 == fruit2) // False
	fmt.Println(fruit1 < fruit2)  // True because a is before p

	// Some string functions

	fmt.Println(strings.Contains(fruit1, "app"))         // True
	fmt.Println(strings.Index(fruit1, "le"))             // 3, Starting index of "le"
	fmt.Println(strings.Replace(fruit1, "le", "app", 1)) // appapp
	fmt.Println(strings.ToUpper(fruit1))                 // Also "ToLower()" can be used
}
