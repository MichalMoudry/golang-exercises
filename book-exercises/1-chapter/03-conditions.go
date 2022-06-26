package main

import "fmt"

func main() {
	testScoreGrade5 := 80
	testScoreGrade4 := 60
	testScoreGrade3 := 50
	testScore := 0

	if testScore >= testScoreGrade5 {
		fmt.Println("Top mark")
	} else if testScore >= testScoreGrade4 {
		fmt.Println("Pass with distinction")
	} else if testScore >= testScoreGrade3 {
		fmt.Println("Pass with distinction")
	} else if testScore < 0 {
		println("Test score shouldn't be nagative")
	} else {
		fmt.Println("Failed")
	}
}
