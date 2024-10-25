package main

import "fmt"

func main() {
	users := []string{"vigo", "lego", "turbo"} // slice

	// ilk değer her zaman slice'ın index'i
	for i := range users {
		fmt.Println("index", i)
	}
	// index 0
	// index 1
	// index 2

	// index, element durumu;
	// ilk değer her zaman slice'ın index'i
	// diğer değer koleksiyondaki element'in o index'deki değeri
	for i, user := range users {
		fmt.Println(i, user)
	}
	// 0 vigo
	// 1 lego
	// 2 turbo

	// blank identifer kullanımı, ilk değeri yutuyoruz
	// diğer değer koleksiyondaki element'in o index'deki değeri
	for _, user := range users {
		fmt.Println(user)
	}
	// vigo
	// lego
	// turbo
}
