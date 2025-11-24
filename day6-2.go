package main
import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func (p Person) IsAdult() bool {
	return p.Age >= 18
}


func (p *Person) HaveBirthday() {
	p.Age += 1
}

func (p *Person) Greet() {
	fmt.Printf("Hello, my name is %s and I am %d years old.\n", p.Name, p.Age)
}


func (p Person) judgeEvenOdd() bool {
	return p.Age % 2 == 0
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

	people := []*Person {
		{"Alice", 30},
		{"Bob", 25},
		{"Charlie", 35},
		{"Diana", 28},
	}

	fmt.Println("Adults in the list:")
	for _, person := range people {
		if person.IsAdult() {
			fmt.Printf("%s is an adult\n", person.Name)
		}
	}

	for _, person := range people {
		person.HaveBirthday()
		fmt.Printf("Happy birthday %s! You are now %d years old.\n", person.Name, person.Age)
	}

	for _, person := range people {
		if person.judgeEvenOdd() {
			fmt.Printf("%s's age %d is even.\n", person.Name, person.Age)
		} else {
			fmt.Printf("%s's age %d is odd.\n", person.Name, person.Age)
		}
	}

	fmt.Println("Greetings from all people:")


	for _, person := range people {
		person.Greet()
	}
}