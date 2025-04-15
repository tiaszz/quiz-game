package main

import (
	"fmt"

	readfile "github.com/tiaszz/quiz-game/readFile"
)

func main() {
	record := readfile.ReadCsvFile("problems.csv")
	answers := readfile.GetAnswerInt(record)
	fmt.Println(record)
	fmt.Println(answers)
}
