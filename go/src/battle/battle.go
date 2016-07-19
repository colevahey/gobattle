package main

import (
	"fmt"
	c "github.com/skilstak/go-colors"
	"github.com/skilstak/go-input"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func main() {
	fmt.Println(c.Clear + c.B3 + "Welcome to Gobattle")
	fmt.Println("Gobattle is a single player game where you battle the cpu to answer questions.")
	time.Sleep(1 * time.Second)
	fmt.Println("There are two teams you can play with. The Blue Team and the Red Team.")
	time.Sleep(1 * time.Second)
	fmt.Println(c.Blue + "The Blue Team focuses on addition and division")
	time.Sleep(1 * time.Second)
	fmt.Println(c.Red + "The Red Team focuses on subtraction and multiplication")
	time.Sleep(1 * time.Second)
	fmt.Println(c.B3 + "Which team would you like to join?")
	fmt.Println(c.Blue + "Blue Team" + c.B3 + " or " + c.Red + "Red Team" + c.B3 + "?")
	var team = input.Ask(c.Green + "")
	team = strings.Title(team)
	if team == "Blue Team" {
		fmt.Println(c.Blue + "You have joined the Blue Team")
		time.Sleep(1 * time.Second)
		blue()
	} else if team == "Red Team" {
		fmt.Println(c.Red + "You have joined the Red Team")
		time.Sleep(1 * time.Second)
		red()
	} else {
		fmt.Println(c.B3 + "That is not a team. Please choose again.")
		time.Sleep(3 * time.Second)
		main()
	}
}

func blue() {
	fmt.Println(c.Clear + "Welcome to the Blue Team.")
	qeasy()
}

func red() {
	fmt.Println(c.Clear + "Welcome to the Red Team.")
	fmt.Println("Here is your first question")
	qeasy()
}

func qeasy() {
	var thisis = rand.Intn(63)
	rand.Seed(thisis)
	var subnum1 = rand.Intn(100)
	var subnum2 = rand.Intn(100)
	//var addnum1 = rand.Intn(100)
	//var addnum2 = rand.Intn(100)
	//var mulnum1 = rand.Intn(15)
	//var mulnum2 = rand.Intn(15)
	//var divnum1 = rand.Intn(100)
	//var divnum2 = 2
	fmt.Println("What is " + strconv.Itoa(subnum1) + " - " + strconv.Itoa(subnum2) + "?")
	var answer = input.Ask("")
	var thing = subnum1 - subnum2
	if answer == strconv.Itoa(thing) {
		fmt.Println("Correct")
	} else {
		fmt.Println("Incorrect")
	}
}
