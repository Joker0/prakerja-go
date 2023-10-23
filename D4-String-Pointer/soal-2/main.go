package main

import (
	"fmt"
)

type Student struct {
	Name  string
	Score float64
}

type StudentList []Student

func (sl StudentList) AverageScore() float64 {
	total := 0.0
	for _, student := range sl {
		total += student.Score
	}
	return total / float64(len(sl))
}

func (sl StudentList) MinScore() (string, float64) {
	if len(sl) == 0 {
		return "", 0.0
	}

	minStudent := sl[0]
	for _, student := range sl {
		if student.Score < minStudent.Score {
			minStudent = student
		}
	}
	return minStudent.Name, minStudent.Score
}

func (sl StudentList) MaxScore() (string, float64) {
	if len(sl) == 0 {
		return "", 0.0
	}

	maxStudent := sl[0]
	for _, student := range sl {
		if student.Score > maxStudent.Score {
			maxStudent = student
		}
	}
	return maxStudent.Name, maxStudent.Score
}

func main() {
	students := make(StudentList, 0)

	for i := 1; i <= 5; i++ {
		var name string
		var score float64

		fmt.Printf("Input %d Student's Name: ", i)
		fmt.Scanln(&name)

		fmt.Printf("Input %d Student's Score: ", i)
		fmt.Scanln(&score)

		student := Student{Name: name, Score: score}
		students = append(students, student)
	}

	averageScore := students.AverageScore()
	minName, minScore := students.MinScore()
	maxName, maxScore := students.MaxScore()

	fmt.Printf("Average Score: %.2f\n", averageScore)
	fmt.Printf("Lowest Score of students: %s (%.2f)\n", minName, minScore)
	fmt.Printf("Highest Score of student: %s (%.2f)\n", maxName, maxScore)
}
