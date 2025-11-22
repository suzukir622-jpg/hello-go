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

    numbers := []int{1, 2, 3, 4, 5, 6}
    even, odd := sumEvenOdd(numbers)
    fmt.Println("Sum of even numbers: ", even)
    fmt.Println("Sum of odd numbers: ", odd)


    greet()
    calculate()
    
}

func greet() {
    name:= "Rin"
    greeting := "Hello " + name
    fmt.Println(greeting)

}

func calculate() {
    var x float64 = 7
    var y float64 = 2.5
    result := x + y
    fmt.Println("Result is: ", result)

}

func  sumEvenOdd(nums []int) (int, int) {
    evenSum := 0
    oddSum := 0
     for _, num := range nums {
        if num%2 == 0 {
            evenSum += num
        } else {
                oddSum += num
            }
        }
     return evenSum, oddSum
}