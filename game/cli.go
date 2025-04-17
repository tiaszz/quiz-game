package game

import (
	"fmt"
	"log"
)

func GetUserInput(input *int) {
	_, err := fmt.Scanln(input)
	if err != nil {
		log.Fatal(err)
	}
}

func Flags() {
}
