package main

import "fmt"

// TYPE DEFINITIONS!
type Printable string
type Status string
type User struct {
	Name string
}

// CONSTANTS WITH DIFFERENT TYPES!!
const StatusOK Status = "OK"

const StatusERROR Printable = "ERROR"

// FUNCTIONS THAT SATISFIES "Stringer" INTERFACE!!

// FOR METHOD STATUS TYPE
func (s Status) String() string {
	return "Status is: " + string(s)
}

// Implement the String method for User
func (u User) String() string {
	return "User: " + u.Name
}

// METHOD FOR PRINTABLE TYPE
func (p Printable) String() string {
	return "Print is : " + string(p)
}

func main() {
	user := User{Name: "Taha"}

	fmt.Println(user)        // User: Taha
	fmt.Println(StatusOK)    // Status is: OK
	fmt.Println(StatusERROR) // Print is : ERROR

}
