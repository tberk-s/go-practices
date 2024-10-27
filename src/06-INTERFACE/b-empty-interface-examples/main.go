package main

import "fmt"

// any is actually interface{} which is empty interface!!
func printAny(i any) {
	fmt.Printf("Type: %T, value: %[1]v\n", i)
}

// We check if the given type is match with the type we assert
// i.(T), where T is acceptable and "ok" will be true
func assertion(i interface{}) {
	if v, ok := i.(string); ok {
		// We use v here instead of "i" because this way we can treat the value
		// as a string instead of empty interface
		// Which allows us to use string operations
		fmt.Println("The value is a string : ", v)
		fmt.Println("Length of the string : ", len(v))
		return
	}
	fmt.Printf("NOT A STRING! Type : %T\n", i)
}

func switchHandle(i any) {
	switch v := i.(type) {
	case int:
		fmt.Println("Value is integer : ", v)
	case string:
		fmt.Println("Value is string : ", v)
	case bool:
		fmt.Println("Value is boolean : ", v)
	default:
		fmt.Printf("Unknown type : %T and value : %[1]v\n", v)
	}
}

func main() {
	fmt.Println("PRINT ANY")
	printAny(42)
	printAny("Hello!")
	printAny(1.2)
	printAny(nil)
	fmt.Println("")

	// -------------------------- TYPE ASSERTION EXAMPLE WITH EMPTY INTERFACE --------------------------
	fmt.Println("TYPE ASSERTION EXAMPLES")
	assertion("GO")
	assertion(12)
	fmt.Println("")

	// -------------------------- TYPE SWITCH EXAMPLE WITH EMPTY INTERFACE -----------------------------
	fmt.Println("TYPE SWITCH EXAMPLES")
	switchHandle("Hello")
	switchHandle(3.14)
	switchHandle(12)
	switchHandle(true)
	switchHandle(nil)
	fmt.Println("")

	// -------------------------- STORE DIFFERENT TYPE OF ITEMS IN A SLICE WITH EMPTY INTERFACE ---------
	fmt.Println("SLICE EXAMPLE")
	data := []any{}
	data = append(data, 12)
	data = append(data, "hello")
	data = append(data, 3.14)
	data = append(data, true)

	for _, v := range data {
		fmt.Printf("Value : %v, Type : %[1]T\n", v)
	}
	fmt.Println("")

	// -------------------------- DIFFERENT TYPE OF VALUES IN A MAP WITH EMPTY INTERFACE ------------
	fmt.Println("MAP EXAMPLE")
	m := make(map[string]any)
	m["foo"] = 14
	m["hello"] = "world"
	m["This is true"] = true

	for key, value := range m {
		fmt.Printf("KEY : %v | Value : %v, Type : %[2]T\n", key, value)
	}
	fmt.Println("")

}
