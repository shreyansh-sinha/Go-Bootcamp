package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Grade string

const (
	A Grade = "A"
	B Grade = "B"
	C Grade = "C"
	F Grade = "F"
)

type student struct {
	firstName, lastName, university                string
	test1Score, test2Score, test3Score, test4Score int
}

type studentStat struct {
	student
	finalScore float32
	grade      Grade
}

func constructStudentsList(data [][]string) []student {

	var studentList []student

	for i, line := range data {
		if i > 0 { // omit header line
			var studentRecord student
			for j, field := range line {
				switch j {
				case 0:
					studentRecord.firstName = field
				case 1:
					studentRecord.lastName = field
				case 2:
					studentRecord.university = field
				case 3:
					value, err := strconv.Atoi(field)
					if err != nil {
						log.Fatal("cannot convert to int")
					}
					studentRecord.test1Score = value
				case 4:
					value, err := strconv.Atoi(field)
					if err != nil {
						log.Fatal("cannot convert to int")
					}
					studentRecord.test2Score = value
				case 5:
					value, err := strconv.Atoi(field)
					if err != nil {
						log.Fatal("cannot convert to int")
					}
					studentRecord.test3Score = value
				case 6:
					value, err := strconv.Atoi(field)
					if err != nil {
						log.Fatal("cannot convert to int")
					}
					studentRecord.test4Score = value
				}
			}

			studentList = append(studentList, studentRecord)
		}
	}

	return studentList
}

func parseCSV(filePath string) []student {

	f, err := os.Open(filePath)

	if err != nil {
		log.Fatal("Error opening file")
		return nil
	}

	defer f.Close()

	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	student := constructStudentsList(data)

	return student

}

func calculateGrade(students []student) []studentStat {

	size := len(students)
	studentsStat := make([]studentStat, size)

	for i := 0; i < size; i++ {

		currStudent := students[i]
		var currStudentStat studentStat
		var grade Grade

		finalScore := float32(currStudent.test1Score+currStudent.test2Score+currStudent.test3Score+currStudent.test4Score) / float32(4)
		currStudentStat.student = currStudent
		currStudentStat.finalScore = finalScore

		if finalScore >= 70 {
			grade = A
		} else if finalScore >= 50 && finalScore < 70 {
			grade = B
		} else if finalScore >= 35 && finalScore < 50 {
			grade = C
		} else {
			grade = F
		}

		currStudentStat.grade = grade

		studentsStat[i] = currStudentStat
	}

	return studentsStat
}

func findOverallTopper(gradedStudents []studentStat) studentStat {

	size := len(gradedStudents)
	var topper studentStat
	maxScore := 0.0
	for i := 0; i < size; i++ {

		currStudentStat := gradedStudents[i]
		if currStudentStat.finalScore > float32(maxScore) {
			maxScore = float64(currStudentStat.finalScore)
			topper = currStudentStat
		}
	}
	return topper
}

func findTopperPerUniversity(gs []studentStat) map[string]studentStat {

	size := len(gs)
	topperUniversityMap := make(map[string]studentStat)

	for i := 0; i < size; i++ {
		currStudentStat := gs[i]
		university := currStudentStat.student.university

		universityMaxStudentStat, check_variable_name := topperUniversityMap[university]

		if !check_variable_name {
			topperUniversityMap[university] = currStudentStat
		} else {
			if currStudentStat.finalScore > universityMaxStudentStat.finalScore {
				topperUniversityMap[university] = currStudentStat
			}
		}
	}

	return topperUniversityMap
}

func main() {

	students := parseCSV("grades.csv")

	studentsStats := calculateGrade(students)

	topperStudentStat := findOverallTopper(studentsStats)

	topperUniversityMap := findTopperPerUniversity(studentsStats)

	fmt.Println("Overall topper is ", topperStudentStat)

	for key, value := range topperUniversityMap {
		fmt.Println(key, value)
	}

}
