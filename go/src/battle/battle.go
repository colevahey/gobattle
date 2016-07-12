package main

import (
	"fmt"
	c "github.com/skilstak/go-colors"
	"github.com/skilstak/go-input"
	"strings"
)

func main() {
	fmt.Println(c.Clear + c.B3 + "GOBATTLE")
	var player = input.Ask("What is your player name? >> " + c.Green)
	player = strings.Title(player)
	fmt.Println(c.B3 + "Welcome " + c.Green + player)
	teampick()
}

func teampick() {
	fmt.Println(c.Clear + c.B3 + "Which team would you like to join?")
	fmt.Println(c.Blue + "Blue Team" + c.B3 + " or " + c.Red + "Red Team" + c.B3 + "?")
	var team = input.Ask("")
	team = strings.Title(team)
	if team == "Blue Team" {
		fmt.Println(c.Blue + "Welcome to the Blue Team " + player)
	} else if team == "Red Team" {
		fmt.Println(c.Red + "Welcome to the Red Team " + player)
	} else {
		fmt.Println(c.B3 + "That is not a team. Please choose again.")
		teampick()
	}
}
