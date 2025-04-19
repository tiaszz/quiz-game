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
func Flags() string {
	csvFilename := flag.String(
		"csv",
		"problems.csv",
		"a csv file in the format of 'question,answer'",
	)
	flag.Parse()
	return *csvFilename
}
