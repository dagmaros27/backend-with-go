package main

import (
	"fmt"
	"strconv"
)

func main() {
	name, grades := getStudentInfo()
	displayResults(name, grades)
}

func getStudentInfo() (string, map[string]int) {
	var name string
	fmt.Print("Enter student name: ")
	fmt.Scanln(&name)

	var numSubjects int
	fmt.Print("Enter number of subjects: ")
	fmt.Scanln(&numSubjects)

	grades := make(map[string]int, numSubjects)
	for i := 0; i < numSubjects; i++ {
		var subject string
		var grade int
		var err error

		fmt.Printf("Enter subject %d name: ", i+1)
		fmt.Scanln(&subject)

		for {
			fmt.Printf("Enter grade for %s (0-100): ", subject)
			var gradeInput string
			fmt.Scanln(&gradeInput)

			grade, err = strconv.Atoi(gradeInput)

			if err != nil {
				fmt.Println("Invalid grade. Please enter a value between 0 and 100.")
				continue
			}
			if grade >= 0 && grade <= 100 {
				break
			}
			fmt.Println("Invalid grade. Please enter a value between 0 and 100.")
		}

		grades[subject] = grade
	}

	return name, grades
}

func calculateAverage(grades map[string]int) int {
	var sum int
	for _, grade := range grades {
		sum += grade
	}
	return sum / (len(grades) - 1)

}

func displayResults(name string, grades map[string]int) {
	fmt.Printf("\nStudent Name: %s\n", name)
	for subject, grade := range grades {

		fmt.Printf("%s: %d\n", subject, grade)
	}

	avg := calculateAverage(grades)
	fmt.Printf("%s: %d\n", "Average", avg)

}
