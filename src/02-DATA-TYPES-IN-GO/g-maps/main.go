package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

var m map[string]int // nil map, key’i string, value’su int...

func main() {
	fmt.Println(m, len(m)) // map[] 0
	// CANNOT ADD KEY INTO THE NIL MAP!!
	// m["foo"] = 5 // panic: assignment to entry in nil map

	m = make(map[string]int)
	m["foo"] = 5
	fmt.Println(m, len(m)) // map[foo:5] 1

	m1 := map[string]int{
		"ocak":  1,
		"şubat": 2,
	}

	m2 := make(map[string]int)
	m2["ocak"] = 1
	m2["şubat"] = 2

	fmt.Println(m1) // map[ocak:1 şubat:2]
	fmt.Println(m2) // map[ocak:1 şubat:2]

	// fmt.Println(m1 == m2)
	// error:
	// invalid operation: m1 == m2 (map can only be compared to nil)

	m1["mart"] = 3
	m2["mart"] = 3

	fmt.Println(m1) // map[mart:3 ocak:1 şubat:2]
	fmt.Println(m2) // map[mart:3 ocak:1 şubat:2]

	delete(m1, "mart") // mart key'ini sil
	fmt.Println(m1)    // map[ocak:1 şubat:2]

	for k, v := range m2 {
		fmt.Println("key", k, "->", v)
	}

	// Maps also have capacity

	fmt.Println(len(m1))           // 2
	fmt.Println(unsafe.Sizeof(m1)) // 8 byte

	fmt.Println(reflect.DeepEqual(m1, m2)) // false

	names := map[string]struct{}{
		"taha":   {},
		"berk":   {},
		"kuyruk": {},
	}

	fmt.Printf("\n%+v\n", names)

	if val, ok := names["taha"]; ok {
		fmt.Println("taha exists") // taha exists inside of the names map
		fmt.Printf("%s", val)      // {} key is empty struct
	}
}
