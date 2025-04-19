package main

import (
	"fmt"

	"github.com/tiaszz/quiz-game/game"
)

func main() {
	var input int
	file := game.Flags()
	record := game.ReadCsvFile(file)
	answers := game.GetAnswerInt(record)

	for i, p := range record {
		fmt.Printf("Problem #%d: %s = ", i+1, p[0])
		game.GetUserInput(&input)
		if input == answers[i] {
			fmt.Println("Correct answer")
		} else {
			fmt.Println("Incorrect answer")
		}
	}
}
