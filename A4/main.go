package main

import "fmt"

func main() {
	// Store 5 integers in an array
	var arr [5]int
	fmt.Println("Enter 5 integers:")

	// Input values
	for i := 0; i < len(arr); i++ {
		fmt.Scan(&arr[i])
	}

	// Print array elements
	fmt.Println("Array elements are:")
	for i := 0; i < len(arr); i++ {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	// Calculate sum and average
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	average := float64(sum) / float64(len(arr))
	fmt.Println("Sum of array elements:", sum)
	fmt.Println("Average of array elements:", average)

	// Find largest and smallest numbers
	largest := arr[0]
	smallest := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > largest {
			largest = arr[i]
		}
		if arr[i] < smallest {
			smallest = arr[i]
		}
	}
	fmt.Println("Largest number:", largest)
	fmt.Println("Smallest number:", smallest)

	// Print array in reverse order
	fmt.Println("Array in reverse order:")
	for i := len(arr) - 1; i >= 0; i-- {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()
}
