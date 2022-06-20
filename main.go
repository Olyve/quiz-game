package main

import (
	"fmt"
	"io/ioutil"

	"./quiz"
)

func main() {
	// Read the file containing the quiz
	contents, err := ioutil.ReadFile("problems.csv")
	if err != nil {
		fmt.Println(err)
		return
	}

	game := new(quiz.Quiz)
	game.create(contents)

	return
}
