package game

import (
	"flag"
	"fmt"
	"log"
)

func GetUserInput(input *int) {
	_, err := fmt.Scanln(input)
	if err != nil {
		log.Fatal(err)
	}
}

// Create the csv flag to read files from it
func Flags() (string, int) {
	csvFilename := flag.String(
		"csv",
		"problems.csv",
		"a csv file in the format of 'question,answer'",
	)
	timeLimit := flag.Int("limit", 30, "The time limit for the quiz in seconds")
	flag.Parse()
	return *csvFilename, *timeLimit
}
