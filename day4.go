package main 
import "fmt"

func main() {

	students := map[string]int{
		"Alice": 85,
		"Bob": 90,
		"Charlie": 78,
	}

fmt.Println("===Problem 1 ===")
for name, age := range students {
	fmt.Printf("%s is %d years old\n", name, age)
}

fmt.Println("===Problem 2 ===")
students["Bob"] = 95
for name, age := range students {
	fmt.Printf("%s is %d years old\n", name, age)
}

fmt.Println("===Problem 3 ===")
delete(students, "Charlie")
for name, age := range students {
	fmt.Printf("%s is %d years old\n", name, age)
}

scores := map[string][]int{
	"Alice": {85, 90, 88},
	"Bob": {78, 82, 80},
}

fmt.Println("===Problem 4 ===")
for name, scoreList := range scores {
	sum := 0
	for _, score := range scoreList {
		sum += score
	}
	average := float64(sum) / float64(len(scoreList))
	fmt.Printf("%s's average score is %.2f\n", name, average)
}

}