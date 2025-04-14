package main

import (
	readfile "github.com/tiaszz/quiz-game/readFile"
)

func main() {
	readfile.ReadCsvFile("./problems.csv")
}
