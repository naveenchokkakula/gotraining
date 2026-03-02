package main

import (
	"fmt"
	"sort"
)

func main() {
	// 1. Create a slice and print
	numbers := []int{10, 20, 30, 40, 50}
	fmt.Println("Original slice:", numbers)

	// 2. Append elements
	numbers = append(numbers, 60, 70)
	fmt.Println("After append:", numbers)

	// 3. Remove element at index 2
	numbers = append(numbers[:2], numbers[3:]...)
	fmt.Println("After removing element at index 2:", numbers)

	// 4. Sum and average
	sum := 0
	for _, val := range numbers {
		sum += val
	}
	fmt.Println("Sum:", sum)
	fmt.Println("Average:", float64(sum)/float64(len(numbers)))

	// 5. Sort slice
	sort.Ints(numbers)
	fmt.Println("Sorted slice:", numbers)
}
