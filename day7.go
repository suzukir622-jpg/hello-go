package main
import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age int
}

func (p Person) IsAdult() bool {
	return p.Age >= 18
}

func (p *Person) HappyBirthday() {
	p.Age += 1
}

func (p *Person) Greet() {
	fmt.Printf("Hello my name is %s and I'm %d years old\n", p.Name, p.Age)
}


func reverse(nums []int) []int {
	reversed := make([]int, len(nums))
	for i, n := range nums {
		reversed[len(nums)-1-i] = n
	}
	return reversed
}


func main() {
	numbers := []int{1, 5, 6, 7, 9}
	reversed := reverse(numbers)
	fmt.Println(reversed)

	ages := map[string]int {
		"Aaron": 30,
		"Beck": 45,
		"Carol": 59,
	}

	ages["David"] = 24

	fmt.Println("Beck's age is", ages["Beck"])

	ages["Aaron"] = 23

	delete(ages, "Carol")

	for name, age := range ages {
		fmt.Printf("%s's age is %d\n", name, age)
	}

	people := []*Person {
		{"Aaron", 34},
		{"Beck", 45},
		{"Charles", 23},
	}

	for _, person := range people{
		if person.IsAdult() {
			fmt.Printf("%s is adult\n", person.Name)
		} 
	}

	for _, person := range people{
		person.HappyBirthday()
		person.Greet()
	}

	sort.Slice(people, func(i,j int) bool {
		return people[i].Name < people[j].Name
	})

	for _, p := range people {
			fmt. Printf("%s is %d years old\n", p.Name, p.Age)
	}
}