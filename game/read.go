package game

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
)

func ReadCsvFile(file string) [][]string {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}

	r := csv.NewReader(f)
	record, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	return record
}

func GetAnswerInt(record [][]string) []int {
	answers := []int{}
	for _, num := range record {
		answersString, err := strconv.Atoi(num[1])
		if err != nil {
			log.Fatal(err)
		}
		answers = append(answers, answersString)
	}
	return answers
}
