package main

/*
	type S struct {
	  Field fieldtype `key1:"value1" key2:"value2"`
	}

*/

import (
	"fmt"
	"log"

	"github.com/isacikgoz/defaults" //go get...
)

type User struct {
	Name     string `validate:"notempty"`
	Email    string `validate:"email"`
	Homepage string `validate:"url"      default:"https://example.com"`
}

func main() {
	var u User

	// set defaults
	if err := defaults.Set(&u); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", u)

	u.Name = "Taha Berk"
	u.Email = "taha@example.com"

	fmt.Printf("%+v\n", u)

	if err := defaults.Validate(&u); err != nil {
		log.Fatal(err)
	}
}
