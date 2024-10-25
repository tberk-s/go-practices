package main

import "fmt"

var user struct {
	firstName string
	lastName  string
	email     string
	password  string
	age       int
}

type users struct {
	firstName string
	//lastName  string
	//email     string
	//password  string
	//age       int
}

// NESTED STRUCTS:
type person struct {
	name    string
	age     int
	address address
}

type address struct {
	city, country string
}

func main() {
	user1 := user
	user1.firstName = "Uğur"
	user1.lastName = "Özyılmazel"
	user1.email = "vigo@xxx.com"
	user1.password = "1234"
	user1.age = 51

	user2 := user
	user2.firstName = "Erhan"
	user2.lastName = "Akpınar"
	user2.email = "erhan@xxx.com"
	user2.password = "1234"
	user2.age = 38
	// anonymous struct
	user3 := struct {
		firstName string
		lastName  string
		email     string
		password  string
		age       int
	}{
		"Ali", // kod okunaklığı açısından iyi değil
		"Desidero",
		"alide@me.com",
		"1234",
		77,
	}

	fmt.Printf("user1.firstName: %s\n", user1.firstName) // Uğur
	fmt.Printf("user2.firstName: %s\n", user2.firstName) // Erhan
	fmt.Printf("user3.firstName: %s\n", user3.firstName) // Ali

	fmt.Printf("user1 :  %v\n", user1)
	fmt.Printf("user2: %v\n", user2)
	fmt.Printf("user3: %v\n", user3)

	// We can skip assigning some fields :
	user4 := user
	user4.age = 23
	user4.firstName = "Taha"
	fmt.Printf("user4: %v\n", user4)
	fmt.Printf("user4: %+v\n", user4)

	/*
		--------------------------------------------------------------

		We can use `new` function to initialize and preallocate the struct in the memory

		new function returns the memory address.

		--------------------------------------------------------------
	*/

	user5 := new(users) // Or &users{} SAME THING
	user6 := users{}

	fmt.Printf("user5: %T\n", user5) // user1: *main.user (pointer geldi)
	fmt.Printf("user6: %T\n", user6)

	fmt.Printf("user5: %v\n", user5) // &{    0}
	fmt.Printf("uesr6: %v\n", user6) // {    0}

	user5.firstName = "Taha"

	fmt.Printf("user5: %v\n", *user5) // DEREFERENCING with *

	/*
		--------------------------------------------------------------

		We can make NESTED STRUCTS too

		--------------------------------------------------------------
	*/

	p1 := person{}
	p1.name = "Taha Berk"
	p1.age = 23
	p1.address = address{
		city:    "İstanbul",
		country: "Türkiye",
	}

	fmt.Printf("%+v\n", p1) // {name:Taha Berk age:23 address:{city:İstanbul country:Türkiye}}

	fmt.Printf("city: %s\n", p1.address.city)       // city: İstanbul
	fmt.Printf("country: %s\n", p1.address.country) // country: Türkiye

}
