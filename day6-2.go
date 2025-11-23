package main
import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	ages := map[string] int{
		"Alice": 30,
		"Bob": 25,
		"Charlie": 35,
	}

	ages["Diana"] = 28

	fmt.Println("Bob's age is", ages["Bob"])

	fmt.Println("Ages of people:")
	for name, age := range ages {
		fmt.Printf("%s is %d years old\n", name, age)
	}

	people := []Person {
		{"Alice", 30},
		{"Bob", 25},
		{"Charlie", 35},
		{"Diana", 28},
	}




}