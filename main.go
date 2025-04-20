package main

import (
	"fmt"
	"time"

	"github.com/tiaszz/quiz-game/game"
)

func main() {
	var input int
	file, timeLimit := game.Flags()
	record := game.ReadCsvFile(file)
	answers := game.GetAnswerInt(record)
	correctAnswers := 0
	timer := time.NewTimer(time.Duration(timeLimit) * time.Second)

	for i, p := range record {
		fmt.Printf("Problem #%d: %s = ", i+1, p[0])
		answerCh := make(chan int)

		go func() {
			game.GetUserInput(&input)
			answerCh <- input
		}()

		select {
		case <-timer.C:
			fmt.Printf("\nYou scored %d out of %d\n", correctAnswers, len(record))
			return
		case input := <-answerCh:
			if input == answers[i] {
				correctAnswers++
			}
		}
	}

	fmt.Printf("You scored %d out of %d\n", correctAnswers, len(record))
}
