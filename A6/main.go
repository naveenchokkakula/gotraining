package main

import "fmt"

// 1. Define Student struct
type Student struct {
	Name  string
	Age   int
	Marks int
}

func main() {
	// 2. Create and print a single Student object
	student1 := Student{Name: "Alice", Age: 20, Marks: 75}
	fmt.Println("Single Student:")
	fmt.Println(student1)

	// 3. Create a slice of Student structs
	students := []Student{
		{Name: "Alice", Age: 20, Marks: 75},
		{Name: "Bob", Age: 21, Marks: 85},
		{Name: "Charlie", Age: 19, Marks: 65},
	}

	// Print all students using range loop
	fmt.Println("All Students:")
	for _, s := range students {
		fmt.Println(s)
	}

	// 4. Find the student with highest marks
	highest := students[0]
	for _, s := range students {
		if s.Marks > highest.Marks {
			highest = s
		}
	}
	fmt.Println("Student with highest marks:")
	fmt.Println(highest)

	// 5. Update marks of a student (e.g., Bob)
	for i, s := range students {
		if s.Name == "Bob" {
			students[i].Marks = 90
		}
	}

	fmt.Println("\nAfter updating Bob's marks:")
	for _, s := range students {
		fmt.Println(s)
	}
}
