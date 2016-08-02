package main

import (
	"fmt"
	c "github.com/skilstak/go-colors"
	"github.com/skilstak/go-input"
	"highscore"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

var Player = input.Ask(c.Clear + c.B3 + "Welcome to Gobattle\nWhat is your player name? >> ")

func main() {
	fmt.Println(c.Clear + "Welcome " + strings.Title(Player))
	time.Sleep(1 * time.Second)
	fmt.Println("Gobattle is a single player game where you battle the global average to answer questions.")
	time.Sleep(1 * time.Second)
	fmt.Println("There are two teams you can play with. The Blue Team and the Red Team.")
	time.Sleep(1 * time.Second)
	fmt.Println(c.Blue + "The Blue Team focuses on addition and division")
	time.Sleep(1 * time.Second)
	fmt.Println(c.Red + "The Red Team focuses on subtraction and multiplication")
	time.Sleep(1 * time.Second)
	fmt.Println(c.B3 + "Which team would you like to join?")
	fmt.Println(c.Blue + "Blue Team" + c.B3 + " or " + c.Red + "Red Team" + c.B3 + "?")
	var team = input.Ask(c.B3 + "")
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
	time.Sleep(2 * time.Second)
	points := 0
	timenow := time.Now().UnixNano()
	rand.Seed(timenow)
	var addnum1 = rand.Intn(100)
	var addnum2 = rand.Intn(100)
	var addnum3 = rand.Intn(1000)
	var addnum4 = rand.Intn(1000)
	var addnum5 = rand.Intn(1500)
	var addnum6 = rand.Intn(1500)
	var divnum1 = rand.Intn(150)
	var divnum2 = 2
	var divnum3 = rand.Intn(80)
	var divnum4 = 3
	var divnum5 = rand.Intn(60)
	var divnum6 = 4
	fmt.Println(c.Clear + "Level 1: Questions are worth one point each.")
	time.Sleep(2 * time.Second)
	fmt.Println(c.Clear + "Q1: What is " + strconv.Itoa(addnum1) + " + " + strconv.Itoa(addnum2) + "?")
	var answer = input.Ask(c.B3 + ">>> ")
	var comp = addnum1 + addnum2
	if answer == strconv.Itoa(comp) {
		fmt.Println(c.Blue + "Correct")
		points += 1
	} else {
		fmt.Println(c.Blue + "Incorrect")
		fmt.Println("The correct answer was " + strconv.Itoa(comp))
	}
	time.Sleep(1 * time.Second)
	fmt.Println(c.Clear + "Q2: What is " + strconv.Itoa(divnum1) + " / " + strconv.Itoa(divnum2) + "?")
	answer = input.Ask(c.B3 + ">>> ")
	comp = divnum1 / divnum2
	if answer == strconv.Itoa(comp) {
		fmt.Println(c.Blue + "Correct")
		points += 1
	} else {
		fmt.Println(c.Blue + "Incorrect")
		fmt.Println("The correct answer was " + strconv.Itoa(comp))
	}
	if points != 1 {
		fmt.Println("You have " + strconv.Itoa(points) + " points")
	} else {
		fmt.Println("You have 1 point")
	}
	time.Sleep(2 * time.Second)
	fmt.Println(c.Clear + "Level 2: Questions are worth two points each.")
	time.Sleep(2 * time.Second)
	fmt.Println(c.Clear + "Q3: What is " + strconv.Itoa(addnum3) + " + " + strconv.Itoa(addnum4) + "?")
	answer = input.Ask(c.B3 + ">>> ")
	comp = addnum3 + addnum4
	if answer == strconv.Itoa(comp) {
		fmt.Println(c.Blue + "Correct")
		points += 2
	} else {
		fmt.Println(c.Blue + "Incorrect")
		fmt.Println("The correct answer was " + strconv.Itoa(comp))
	}
	time.Sleep(1 * time.Second)
	fmt.Println(c.Clear + "Q4: What is " + strconv.Itoa(divnum3) + " / " + strconv.Itoa(divnum4) + "?")
	answer = input.Ask(c.B3 + ">>> ")
	comp = divnum3 / divnum4
	if answer == strconv.Itoa(comp) {
		fmt.Println(c.Blue + "Correct")
		points += 2
	} else {
		fmt.Println(c.Blue + "Incorrect")
		fmt.Println("The correct answer was " + strconv.Itoa(comp))
	}
	if points != 1 {
		fmt.Println("You have " + strconv.Itoa(points) + " points")
	} else {
		fmt.Println("You have 1 point")
	}
	time.Sleep(2 * time.Second)
	fmt.Println(c.Clear + "Level 3: Questions are worth three points each.")
	time.Sleep(2 * time.Second)
	fmt.Println(c.Clear + "Q5: What is " + strconv.Itoa(addnum5) + " + " + strconv.Itoa(addnum6) + "?")
	answer = input.Ask(c.B3 + ">>> ")
	comp = addnum5 + addnum6
	if answer == strconv.Itoa(comp) {
		fmt.Println(c.Blue + "Correct")
		points += 3
	} else {
		fmt.Println(c.Blue + "Incorrect")
		fmt.Println("The correct answer was " + strconv.Itoa(comp))
	}
	time.Sleep(1 * time.Second)
	fmt.Println(c.Clear + "Q6: What is " + strconv.Itoa(divnum5) + " / " + strconv.Itoa(divnum6) + "?")
	answer = input.Ask(c.B3 + ">>> ")
	comp = divnum5 / divnum6
	if answer == strconv.Itoa(comp) {
		fmt.Println(c.Blue + "Correct")
		points += 3
	} else {
		fmt.Println(c.Blue + "Incorrect")
		fmt.Println("The correct answer was " + strconv.Itoa(comp))
	}
	if points != 1 {
		fmt.Println("You finished with " + strconv.Itoa(points) + " points")
	} else {
		fmt.Println("You finished with 1 point")
	}
	time.Sleep(2 * time.Second)
	highlist()
}

func red() {
	fmt.Println(c.Clear + "Welcome to the Red Team.")
	time.Sleep(2 * time.Second)
	points := 0
	timenow := time.Now().UnixNano()
	rand.Seed(timenow)
	var subnum1 = rand.Intn(100)
	var subnum2 = rand.Intn(100)
	var subnum3 = rand.Intn(1000)
	var subnum4 = rand.Intn(1000)
	var subnum5 = rand.Intn(1500)
	var subnum6 = rand.Intn(1500)
	var mulnum1 = rand.Intn(20)
	var mulnum2 = rand.Intn(20)
	var mulnum3 = rand.Intn(30)
	var mulnum4 = rand.Intn(30)
	var mulnum5 = rand.Intn(45)
	var mulnum6 = rand.Intn(45)
	fmt.Println(c.Clear + "Level 1: Questions are worth one point each.")
	time.Sleep(2 * time.Second)
	fmt.Println(c.Clear + "Q1: What is " + strconv.Itoa(subnum1) + " - " + strconv.Itoa(subnum2) + "?")
	var answer = input.Ask(c.B3 + ">>> ")
	var comp = subnum1 - subnum2
	if answer == strconv.Itoa(comp) {
		fmt.Println(c.Red + "Correct")
		points += 1
	} else {
		fmt.Println(c.Red + "Incorrect")
		fmt.Println("The correct answer was " + strconv.Itoa(comp))
	}
	time.Sleep(1 * time.Second)
	fmt.Println(c.Clear + "Q2: What is " + strconv.Itoa(mulnum1) + " * " + strconv.Itoa(mulnum2) + "?")
	answer = input.Ask(c.B3 + ">>> ")
	comp = mulnum1 * mulnum2
	if answer == strconv.Itoa(comp) {
		fmt.Println(c.Red + "Correct")
		points += 1
	} else {
		fmt.Println(c.Red + "Incorrect")
		fmt.Println("The correct answer was " + strconv.Itoa(comp))
	}
	if points != 1 {
		fmt.Println("You have " + strconv.Itoa(points) + " points")
	} else {
		fmt.Println("You have 1 point")
	}
	time.Sleep(2 * time.Second)
	fmt.Println(c.Clear + "Level 2: Questions are worth two points each.")
	time.Sleep(2 * time.Second)
	fmt.Println(c.Clear + "Q3: What is " + strconv.Itoa(subnum3) + " - " + strconv.Itoa(subnum4) + "?")
	answer = input.Ask(c.B3 + ">>> ")
	comp = subnum3 - subnum4
	if answer == strconv.Itoa(comp) {
		fmt.Println(c.Red + "Correct")
		points += 2
	} else {
		fmt.Println(c.Red + "Incorrect")
		fmt.Println("The correct answer was " + strconv.Itoa(comp))
	}
	time.Sleep(1 * time.Second)
	fmt.Println(c.Clear + "Q4: What is " + strconv.Itoa(mulnum3) + " * " + strconv.Itoa(mulnum4) + "?")
	answer = input.Ask(c.B3 + ">>> ")
	comp = mulnum3 * mulnum4
	if answer == strconv.Itoa(comp) {
		fmt.Println(c.Red + "Correct")
		points += 2
	} else {
		fmt.Println(c.Red + "Incorrect")
		fmt.Println("The correct answer was " + strconv.Itoa(comp))
	}
	if points != 1 {
		fmt.Println("You have " + strconv.Itoa(points) + " points")
	} else {
		fmt.Println("You have 1 point")
	}
	time.Sleep(2 * time.Second)
	fmt.Println(c.Clear + "Level 3: Questions are worth three points each.")
	time.Sleep(2 * time.Second)
	fmt.Println(c.Clear + "Q5: What is " + strconv.Itoa(subnum5) + " - " + strconv.Itoa(subnum6) + "?")
	answer = input.Ask(c.B3 + ">>> ")
	comp = subnum5 - subnum6
	if answer == strconv.Itoa(comp) {
		fmt.Println(c.Red + "Correct")
		points += 3
	} else {
		fmt.Println(c.Red + "Incorrect")
		fmt.Println("The correct answer was " + strconv.Itoa(comp))
	}
	time.Sleep(1 * time.Second)
	fmt.Println(c.Clear + "Q6: What is " + strconv.Itoa(mulnum5) + " * " + strconv.Itoa(mulnum6) + "?")
	answer = input.Ask(c.B3 + ">>> ")
	comp = mulnum5 * mulnum6
	if answer == strconv.Itoa(comp) {
		fmt.Println(c.Red + "Correct")
		points += 3
	} else {
		fmt.Println(c.Red + "Incorrect")
		fmt.Println("The correct answer was " + strconv.Itoa(comp))
	}
	if points != 1 {
		fmt.Println("You finished with " + strconv.Itoa(points) + " points")
	} else {
		fmt.Println("You finished with 1 point")
	}
	time.Sleep(2 * time.Second)
	highlist()
}

func highlist() {
	highscore.Highscore()
}
