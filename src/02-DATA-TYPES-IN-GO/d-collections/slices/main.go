package main

import "fmt"

func main() {
	var users []string // boş slice tanımı, len=0

	fmt.Println("users", users, len(users))                      // users [] 0
	fmt.Printf("users: %v || %#[1]v || %d\n", users, len(users)) // users: [] || []string(nil) || 0

	admins := []string{"taha", "berk"} // 2 elemanı olan slice

	fmt.Println("admins", admins, len(admins)) // admins [taha berk] 2
	fmt.Printf(
		"admins: %v || %#[1]v || %d\n",
		admins,
		len(admins),
	)

	users = append(users, "ahmet", "mehmet")

	fmt.Printf("users: %v || %#[1]v || %d\n", users, len(users))

	/*
		--------------------------------------------------------------

		Slices holds POINTERS, length, and capacity...
		We can define the slice with `make` function
		It is called "preallocate"

		--------------------------------------------------------------
	*/

	slice1 := make([]int, 5) // len: 5, default cap => 5
	slice2 := make([]int, 2) // len: 2, default cap => 2

	fmt.Printf("slice1, len: %d, cap: %d\n", len(slice1), cap(slice1)) // slice1, len: 5, cap: 5
	fmt.Printf("slice2, len: %d, cap: %d\n", len(slice2), cap(slice2)) // slice2, len: 2, cap: 2

	fmt.Println() // boş satır

	slice1 = append(slice1, 1)                                         // 1 eleman ekle, kapasiteyi aş, ek alan ekle
	fmt.Printf("slice1, len: %d, cap: %d\n", len(slice1), cap(slice1)) // slice1, len: 6, cap: 10

	slice2 = append(slice2, 1)                                         // 1 eleman ekle, kapasiteyi aş, ek alan ekle
	fmt.Printf("slice2, len: %d, cap: %d\n", len(slice2), cap(slice2)) // slice2, len: 3, cap: 4

	fmt.Println() // boş satır

	// kapasite varsayılan uzunluğa göre büyür

	slice3 := make([]string, 0, 4)                                     // len: 0, default cap => 4
	fmt.Printf("slice3, len: %d, cap: %d\n", len(slice3), cap(slice3)) // slice3, len: 0, cap: 4

	slice3 = append(slice3, "1 daha")
	fmt.Printf("slice3, len: %d, cap: %d\n", len(slice3), cap(slice3)) // slice3, len: 1, cap: 4

	slice3 = append(slice3, "2 daha")
	fmt.Printf("slice3, len: %d, cap: %d\n", len(slice3), cap(slice3)) // slice3, len: 2, cap: 4

	slice3 = append(slice3, "3 daha")
	fmt.Printf("slice3, len: %d, cap: %d\n", len(slice3), cap(slice3)) // slice3, len: 3, cap: 4

	slice3 = append(slice3, "4 daha")
	fmt.Printf("slice3, len: %d, cap: %d\n", len(slice3), cap(slice3)) // slice3, len: 4, cap: 4

	// güm! kapasiteyi arttır, taştık çünkü!
	slice3 = append(slice3, "more 5, overflow!")
	fmt.Printf("slice3, len: %d, cap: %d\n", len(slice3), cap(slice3)) // slice3, len: 5, cap: 8

	// slice3[100] = "foo" // panic: runtime error: index out of range [100] with length 5

	/*
		--------------------------------------------------------------
		ADD ONE doesn't change the memory address of slice, so it stays same

		BUT

		APPED ONE changes the memory address of the slice, it becomes brand new.
		--------------------------------------------------------------
	*/

	// **** ADD ****
	numbers := []int{1, 2, 3}
	fmt.Println("numbers - başlangıç değeri", numbers) // numbers - başlangıç değeri [1 2 3]

	addOne(numbers)
	fmt.Println("numbers - addOne sonrası", numbers)     // numbers - addOne sonrası [2 3 4]
	fmt.Printf("numbers - hafıza adresi: %p\n", numbers) // numbers - hafıza adresi: 0xc0000ae000

	// **** APPED ****
	numbers2 := []int{1, 2, 3}
	fmt.Println("numbers - başlangıç değeri", numbers) // numbers - başlangıç değeri [1 2 3]

	appendOne(numbers2)
	fmt.Println("numbers - appendOne sonrası", numbers2)  // numbers - appendOne sonrası [1 2 3]
	fmt.Printf("numbers - hafıza adresi: %p\n", numbers2) // numbers - hafıza adresi: 0x140000ac000

}

func addOne(s []int) {
	fmt.Printf("s hafıza adresi: %p\n", s) // s hafıza adresi: 0xc0000ae000
	for i := range s {
		s[i]++ // index i'deki elemanın değerini 1 arttır
	}
}

func appendOne(s []int) {
	fmt.Printf("s hafıza adresi: %p\n", s)          // s hafıza adresi: 0x140000ac000
	s = append(s, 1)                                // bu s artık başka bir slice
	fmt.Printf("s hafıza adresi (append): %p\n", s) // s hafıza adresi (append): 0x140000b2000
}
