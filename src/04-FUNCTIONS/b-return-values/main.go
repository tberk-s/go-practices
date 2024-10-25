package main

import (
	"fmt"
	"log"
)

type users map[string]struct{}

func greetFromMap(u users, name string) (string, error) { // Function signature, returns a string and an error!!!
	if _, ok := u[name]; !ok {
		return "", fmt.Errorf("%s not found in map", name)
	}
	return "hello " + name, nil
}

func main() {
	u := users{
		"taha":   {},
		"kek":    {},
		"kuyruk": {},
	}

	peopleToSayHello := []string{"taha", "berk", "kuyruk", "ahmet"}

	g, err := greetFromMap(u, "taha")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(g) // hello taha

	for _, name := range peopleToSayHello {
		result, err := greetFromMap(u, name)
		if err != nil {
			log.Printf("inside of the foor loop, error : %v", err)
			continue
		}
		fmt.Println(result)

	}
	g, err = greetFromMap(u, "berk")
	if err != nil {
		log.Fatal(err)
		// 2023/08/06 13:15:10 berk not found in map
		// exit status 1
	}

	fmt.Println(g)
}
