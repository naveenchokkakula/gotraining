package main

import "fmt"

func main() {
	print1to100()
	printEvenNumbers()
	multiplicationTable()
	factorial()
	fibonacci()
}

// 1️⃣ Print numbers from 1 to 100
func print1to100() {
	fmt.Println("Numbers from 1 to 100:")
	for i := 1; i <= 100; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println("\n")
}

// 2️⃣ Print even numbers between 1 and 50
func printEvenNumbers() {
	fmt.Println("Even numbers between 1 and 50:")
	for i := 2; i <= 50; i += 2 {
		fmt.Print(i, " ")
	}
	fmt.Println("\n")
}

// 3️⃣ Multiplication table
func multiplicationTable() {
	var num int
	fmt.Print("Enter a number for multiplication table: ")
	fmt.Scan(&num)

	for i := 1; i <= 10; i++ {
		fmt.Printf("%d x %d = %d\n", num, i, num*i)
	}
	fmt.Println()
}

// 4️⃣ Factorial
func factorial() {
	var n int
	fmt.Print("Enter a number to find factorial: ")
	fmt.Scan(&n)

	fact := 1
	for i := 1; i <= n; i++ {
		fact *= i
	}
	fmt.Println("Factorial:", fact)
	fmt.Println()
}

// 5️⃣ Fibonacci series
func fibonacci() {
	var terms int
	fmt.Print("Enter number of Fibonacci terms: ")
	fmt.Scan(&terms)

	a := 0
	b := 1

	fmt.Println("Fibonacci series:")
	for i := 1; i <= terms; i++ {
		fmt.Print(a, " ")
		next := a + b
		a = b
		b = next
	}
	fmt.Println()
}
