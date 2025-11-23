package main
import (
	"fmt"
	"sort"
)

func sum(nums []int) int {
	total := 0
	for _, n:= range nums {
		total += n
	}
	return total
}

func avg(nums []int) float64 {
	total := sum(nums)
	return float64(total) / float64(len(nums))
}

func splitEvenOdd(nums []int) ([]int, []int) {
	even := []int{}
	odd := []int{}
	for _, n:= range nums {
		if n%2 == 0{
			even = append(even, n)
		} else {
			odd = append(odd, n)
		}
	}
	return even, odd
}

func evenSorted(nums []int) ([]int) {
	even := []int{}
	for _, n := range nums {
		if n%2 == 0 {
		even = append(even, n)
		}
	}
	sort. Ints(even)
	return even
}

func sortEven(nums []int) ([]int) {
	even := []int{}
	sort.Ints(nums)
  for _, n := range nums {
		if n%2 == 0 {
			even = append(even, n)
		}
	}
	return even
}

func reverseStrings(strs []string) []string {
	reversed := make([]string, len(strs))
	for i, s := range strs {
		reversed[len(strs)-1-i]= s
	}
	return reversed
}

func sumslice(num []int) int {
	total := 0
	for _, n := range num {
		total += n
	}
	return total
} 





func main() {
	numbers := []int{10, 20, 30, 40, 50}

	slice1 := numbers[1:4]
	fmt.Println(slice1)
	slice2 := numbers[:3]
	fmt.Println(slice2)
	slice3 := numbers[2:]
	fmt.Println(slice3)

	sumsum := sum(numbers)
	fmt.Println(sumsum)

	avgslice := avg(numbers)
	fmt.Println(avgslice)

	even, odd := splitEvenOdd(numbers)
	fmt.Println("Even numbers: ", even)
	fmt.Println("Odd numbers: ", odd)

  input := []int{2, 5, 7, 4, 6, 3}
	output := evenSorted(input)
	fmt.Println(output)

	input2 := []int{42, 17, 8, 99, 23, 5}
	sortedEven := sortEven(input2)
	fmt.Println(sortedEven)

	input3 := []string{"banana", "apple", "orange", "grape", "kiwi"}
  reversed := reverseStrings(input3)
	fmt.Println(reversed)

	numbers2 := []int{1, 2, 3, 4, 5}
  sum := sumslice(numbers2)
  fmt.Println(sum)
}

