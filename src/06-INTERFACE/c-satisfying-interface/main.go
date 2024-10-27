package main

import (
	"fmt"
	"io"
	"os"
)

// Interface
type Describer interface {
	Describe() string
}

type Person struct {
	Name string
	Age  int
}

// This func can satisfy "Describer" interface since it has "Describe" method
func (p Person) Describe() string {
	return fmt.Sprintf("Name: '%s' Age: %v", p.Name, p.Age)
}

type Car struct {
	Model string
	Year  int
}

// Also satisfies "Describer" interface
func (c Car) Describe() string {
	return fmt.Sprintf("Car: %v %s", c.Year, c.Model)
}

func printDescription(d Describer) {
	fmt.Println(d.Describe())
}

// We can pass the interface as an argument to the function, "io.Reader"
func readStream(r io.Reader) (string, error) {
	b := make([]byte, 1024)
	n, err := r.Read(b)
	if err != nil {
		return "", fmt.Errorf("read error: %w", err)
	}
	return string(b[:n]), nil
}

func main() {
	b := Person{Name: "Taha", Age: 23}
	c := Car{Model: "Random Car", Year: 2024}

	printDescription(b)
	printDescription(c)

	// You need to create test.txt by hand if you didnt use "git clone"
	file, err := os.Open("src/06-INTERFACE/c-satisfying-interface/test.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	content, err := readStream(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
	} else {
		fmt.Println("File read ->", content)
	}
}
