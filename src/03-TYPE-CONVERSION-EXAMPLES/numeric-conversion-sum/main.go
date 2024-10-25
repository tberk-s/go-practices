package main

import "fmt"

func main() {
	sum(1, 2, 3, 4.2, "helolololo", 1.1)
}

func sum(kwargs ...any) float64 {
	var count float64

	for _, value := range kwargs {
		fmt.Printf("%v, %[1]T\n", value)

		switch v := value.(type) {
		case int:
			count += float64(v)
		case float64:
			count += v
		}
	}
	fmt.Printf("%.2f\n", count)
	return count
}
