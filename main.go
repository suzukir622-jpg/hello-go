package main

import "fmt"

func main() {
    fmt.Println("Hello, Go!")
    fmt.Println(3+5)
    fmt.Println("Go is fun!")
    fmt.Println("Day2 of Go lang")

    a := 10
    b := 20
    sum := a + b
    fmt.Println("Sum is:", sum)

    greet()
    result()
}

func greet() {
    name:= "Rin"
    greeting := "Hello " + name
    fmt.Println(greeting)

}

func calculate() {
    var x int = 7
    var y int = 2.5
    result := x + y
    fmt.Println("Result is: ", result)

}