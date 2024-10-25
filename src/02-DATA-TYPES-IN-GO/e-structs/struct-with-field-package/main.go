package main

import (
	"fmt"

	"github.com/tberk-s/go-practices/src/02-DATA-TYPES-IN-GO/e-structs/struct-with-field-package/person"
)

func main() {
	p := person.Person{} // boş bir kopya (instance)

	p.FirstName = "Taha"
	p.LastName = "Berk"

	fmt.Printf("p: %#v\n", p)           // p: person.Person{FirstName:"Taha", LastName:"Berk", secret:""}
	fmt.Printf("p: %+v\n", p.FirstName) // p: Taha

	//(fmt.Println(p.secret) // p.secret undefined (type person.Person has no field or method secret)
}
