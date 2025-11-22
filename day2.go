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
	even, odd := evenoddreverse(numbers)
	fmt.Println("Reversed even numbers:", even)
	fmt.Println("Reversed odd numbers:", odd)

	sumeven1 := sumeven(numbers)
	fmt.Println("Sum of even numbers is:", sumeven1)
	reverseevensum := sumeven(even)
	fmt.Println("Sum of reversed even numbers is:", reverseevensum)

	avg := average(numbers)
	fmt.Println("Average of numbers is:", avg)

	counteven, countodd := countevenodd(numbers)
	fmt.Println("Count of even numbers:", counteven)
	fmt.Println("Count of odd numbers:", countodd)

	maxVal, minVal, maxIdx, minIdx := maxMinIndex(numbers)
	fmt.Println("Maximum value is:", maxVal, "at index", maxIdx)
	fmt.Println("Minimum value is:", minVal, "at index", minIdx)
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

func evenoddreverse(nums []int) ([]int, []int) {
	even:= []int{}
	odd:= []int{}
	
	for i:= len(nums) - 1; i >=0; i-- {
		if nums[i]%2 == 0 {
			even = append(even, nums[i])
		} else {
			odd = append(odd, nums[i])
		}
	}
	return even, odd
}

func sumeven(nums []int) int {
	sumeven := 0
	for _, n := range nums {
		if n%2 == 0 {
			sumeven += n
		}
	}
	return sumeven
}

func average (nums []int) float64 {
	total := 0
	for _, n := range nums {
		total += n
	}
	return float64(total) / float64(len(nums))
}

func countevenodd (nums []int) (int, int) {
	even:= 0
	odd:= 0
	for _, n := range nums {
		if n%2 == 0 {
			even ++
		} else {
			odd ++
		}
	}
	return even, odd
}

func maxMinIndex(nums []int) (int, int, int, int) {
	maxIdx := 0
	minIdx := 0
	maxVal := nums[0]
	minVal := nums[0]
	
	for i,n := range nums {
		if n > maxVal {
			maxVal = n
			maxIdx = i
		}
		if n < minVal {
			minVal = n
			minIdx = i
		}
	}
return maxVal, minVal, maxIdx, minIdx
}