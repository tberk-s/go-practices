package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	sum(1, 2, 3, 4, "5", "2.1", "7.5", 1.1)
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
		case string:
			n, err := strconv.ParseFloat(v, 64) //Ascii to float... //// use strconv.Atoi(v) for ascii to integer...
			if err != nil {
				log.Fatal(err) // Error handling and exit...
			}

			count += n

		}

	}
	fmt.Printf("%.2f\n", count)
	return count
}
