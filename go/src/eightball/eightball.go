package main

import (
	"fmt"
	c "github.com/skilstak/go/colors"
)

var answers = []string{
	"Yes.",
	"No.",
	"Maybe.",
}

var responses = []easterEgg{
	{"die|death", "Don't be so down"},
	{"love", "I do not speak of matters of love"},
}

type easterEgg struct {
	key      string
	response string
}

func main() {
	fmt.Println(c.Clear + c.Multi("Welcome to the magical eightball"))
}
