package main

import "fmt"

func main() {
	evenOdd()
	largestOfThree()
	leapYear()
}

// Q1
func evenOdd() {
	var num int
	fmt.Print("Enter a number: ")
	fmt.Scan(&num)

	if num%2 == 0 {
		fmt.Println("Even number")
	} else {
		fmt.Println("Odd number")
	}
}

// Q2
func largestOfThree() {
	var a, b, c int
	fmt.Print("\nEnter three numbers: ")
	fmt.Scan(&a, &b, &c)

	if a >= b && a >= c {
		fmt.Println("Largest:", a)
	} else if b >= a && b >= c {
		fmt.Println("Largest:", b)
	} else {
		fmt.Println("Largest:", c)
	}
}

// Q3
func leapYear() {
	var year int
	fmt.Print("\nEnter a year: ")
	fmt.Scan(&year)

	if (year%4 == 0 && year%100 != 0) || (year%400 == 0) {
		fmt.Println("Leap year")
	} else {
		fmt.Println("Not a leap year")
	}
}
