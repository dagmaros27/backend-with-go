package main

import (
	"fmt"
)

func evaluateAverageGrade(grades []int) int {
	if len(grades) == 0 {
		return 0
	}
	total := 0
	for _, val := range grades {
		total += val
	}
	average := float64(total) / float64(len(grades))
	return int(average)
}

func isValidGrade(val int) bool {
	return val >= 0 && val <= 100
}

func main() {
	var name string
	fmt.Print("Enter name: ")
	fmt.Scanln(&name)

	subjects := make(map[string]int)
	grades := make([]int, 0)

	fmt.Println("Enter subject and respective grade (e.g., 'English 100'). Type 'q' to finish:")

	for {
		var subject string
		var grade int
		fmt.Scanln(&subject, &grade)

		if subject == "q" {
			break
		}

		if !isValidGrade(grade) {
			fmt.Println("Grade must be between 0 and 100. Please try again.")
			continue
		}

		subjects[subject] = grade
		grades = append(grades, grade)
		fmt.Println("Enter next subject and grade or type 'q' to finish:")
	}

	fmt.Printf("\nName: %s\n", name)
	fmt.Printf("%-10s %s\n", "Subject", "Grade")
	for subject, val := range subjects {
		fmt.Printf("%-10s %d\n", subject, val)
	}
	fmt.Printf("%-10s %d\n", "Average:", evaluateAverageGrade(grades))
}
