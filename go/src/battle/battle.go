package main

import (
	"fmt"
	c "github.com/skilstak/go-colors"
	"github.com/skilstak/go-input"
	"strings"
	"time"
)

func main() {
	fmt.Println(c.Clear + c.B3 + "Welcome to Gobattle")
	var player = input.Ask("What is your player name? >> " + c.Green)
	player = strings.Title(player)
	fmt.Println(c.B3 + "Welcome " + player)
	time.Sleep(2 * time.Second)
	fmt.Println("There are two teams you can play with. The Red Team and the Blue Team.")
	time.Sleep(2 * time.Second)
	fmt.Println("Which team would you like to join?")
	fmt.Println(c.Blue + "Blue Team" + c.B3 + " or " + c.Red + "Red Team" + c.B3 + "?")
	var team = input.Ask(c.Green + "")
	team = strings.Title(team)
	if team == "Blue Team" {
		fmt.Println(c.Blue + "You have joined the Blue Team")
	} else if team == "Red Team" {
		fmt.Println(c.Red + "You have joined the Red Team")
	} else {
		fmt.Println(c.B3 + "That is not a team. Please choose again.")
		time.Sleep(3 * time.Second)
		main()
	}
}
