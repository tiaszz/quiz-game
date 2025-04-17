package main

import (
	"fmt"

	"github.com/tiaszz/quiz-game/game"
)

func main() {
	// record := game.ReadCsvFile("problems.csv")
	// answers := game.GetAnswerInt(record)
	var userInput int
	game.GetUserInput(&userInput)
	fmt.Println(userInput)
}
