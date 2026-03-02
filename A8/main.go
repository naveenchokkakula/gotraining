package main

import "fmt"

type Student struct {
	Name  string
	Marks int
}

func modifyValue(x *int) {
	*x += 10
}

func swap(a, b *int) {
	*a, *b = *b, *a
}

func modifyStudent(s *Student) {
	s.Marks += 20
	s.Name += " Jr."
}

func main() {
	// Pointer with a variable
	num := 50
	ptr := &num
	*ptr = 100
	fmt.Println("num after pointer change:", num)

	// Modify variable using function
	modifyValue(&num)
	fmt.Println("num after modifyValue:", num)

	// Swap two numbers
	a, b := 10, 20
	swap(&a, &b)
	fmt.Println("After swap: a =", a, "b =", b)

	// Modify struct using pointer
	student := Student{Name: "Alice", Marks: 75}
	modifyStudent(&student)
	fmt.Println("Modified student:", student)
}
