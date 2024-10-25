package main

import (
	"fmt"

	"github.com/tberk-s/go-practices/src/02-DATA-TYPES-IN-GO/e-structs/struct-with-field-access-getter/person"
)

func main() {
	p1 := person.Person{} // boş bir kopya (instance)

	p1.FirstName = "Taha"
	p1.LastName = "Berk"

	fmt.Printf("%+v\n", p1) // {FirstName:Taha LastName:Berk secret:}

	p1.SetSecret("very secret password")

	fmt.Printf("%+v\n", p1)  // {FirstName:Taha LastName:Berk secret:<secret>}
	fmt.Println(p1.Secret()) // <secret>

	fmt.Println(p1.FullName("p1")) // Taha Berk

	p2 := person.Person{FirstName: "Ahmet", LastName: "Mehmet"}

	p2.SetSecret("Another secret password")
	fmt.Println(p2.FullName("p2"))
}
