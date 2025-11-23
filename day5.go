package main
import "fmt"

type Person struct {
	FName string
	LName string
	Age int
}

func (p Person) Greet() {
	fmt.Println("こんにちは、私は", p.FName, "です" )
}

func (p *Person) HaveBirthday() {
	p.Age += 1
	fmt.Println("お誕生日おめでとう！私は今", p.Age, "歳です")
}

func (p Person) IsAdult() {
	if p.Age >= 20 {
		fmt.Println(p.FName, "は成人です")
	} else {
		fmt.Println(p.FName, "は未成年です")
	}

	}

func (p Person) FullName() string {
	Full := p.FName + p.LName
	fmt.Println("私のフルネームは", Full, "です")
	return Full
}

func main() {
	p := Person {FName: "Rin", LName: "Suzuki", Age: 25}
	p.Greet()
  p.HaveBirthday()
	p.IsAdult()
	p.FullName()
}

