package main

import "fmt"

// Student struct
type Student struct {
	Name  string
	Marks int
}

// Method to display details
func (s Student) DisplayDetails() {
	fmt.Println("Student Name:", s.Name)
	fmt.Println("Marks:", s.Marks)
}

// Method to update marks
func (s *Student) UpdateMarks(newMarks int) {
	s.Marks = newMarks
}

// Method to check pass/fail
func (s Student) IsPassed() bool {
	return s.Marks >= 40
}

// Method to get grade
func (s Student) Grade() string {
	switch {
	case s.Marks >= 90:
		return "A+"
	case s.Marks >= 80:
		return "A"
	case s.Marks >= 70:
		return "B+"
	case s.Marks >= 60:
		return "B"
	case s.Marks >= 50:
		return "C"
	case s.Marks >= 40:
		return "D"
	default:
		return "F"
	}
}

func main() {
	student1 := Student{Name: "Alice", Marks: 75}
	student1.DisplayDetails()
	fmt.Println("Passed?", student1.IsPassed())
	fmt.Println("Grade:", student1.Grade())

	student1.UpdateMarks(35)
	fmt.Println("\nAfter updating marks:")
	student1.DisplayDetails()
	fmt.Println("Passed?", student1.IsPassed())
	fmt.Println("Grade:", student1.Grade())
}
