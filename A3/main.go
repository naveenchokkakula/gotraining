package main

import "fmt"

func main() {

	addNumbers()
	checkPrime()
	maxOfTwo()
	reverseNumber()
}

// Function to add two numbers
func addNumbers() {
	var a, b int
	fmt.Print("Enter two numbers to add: ")
	fmt.Scan(&a, &b)

	sum := a + b
	fmt.Println("Sum:", sum)
	fmt.Println()
}

// Function to check whether a number is prime
func checkPrime() {
	var n int
	fmt.Print("Enter a number to check prime: ")
	fmt.Scan(&n)

	if n <= 1 {
		fmt.Println(n, "is not prime")
		return
	}

	isPrime := true
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			isPrime = false
			break
		}
	}

	if isPrime {
		fmt.Println(n, "is prime")
	} else {
		fmt.Println(n, "is not prime")
	}
	fmt.Println()
}

// Function to find maximum of two numbers
func maxOfTwo() {
	var a, b int
	fmt.Print("Enter two numbers to find maximum: ")
	fmt.Scan(&a, &b)

	if a > b {
		fmt.Println("Maximum is:", a)
	} else {
		fmt.Println("Maximum is:", b)
	}
	fmt.Println()
}

// Function to reverse a number
func reverseNumber() {
	var n, rev int
	fmt.Print("Enter a number to reverse: ")
	fmt.Scan(&n)

	temp := n
	rev = 0
	for temp != 0 {
		digit := temp % 10
		rev = rev*10 + digit
		temp /= 10
	}

	fmt.Println("Reversed number:", rev)
	fmt.Println()
}
