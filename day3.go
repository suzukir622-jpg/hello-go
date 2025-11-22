package main
import "fmt"
func main() {
	fmt.Println("Welcome to Day 3 of Go Lang!")
	numbers := []int{34, 12, 5, 67, 23, 89, 2}
	max, min := maxMin(numbers)
	fmt.Println("Maximum number is:", max)
	fmt.Println("Minimum number is:", min)
	reversed := reverse(numbers)
	fmt.Println("Reversed slice is:", reversed)
}


func maxMin(nums []int) (int, int) {
	max:= nums[0]
	min:= nums[0]
	for _, n := range nums {
		if n > max {
			max = n
		}
		if n < min {
			min = n
		}
	}
	return max, min
}

func reverse(nums []int) []int {
	reversed := make([]int, len(nums))
	for i, n := range nums {
		reversed[len(nums)-1-i] = n
	}
	return reversed
}