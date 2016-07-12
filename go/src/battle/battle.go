package main

import (
	"fmt"
	c "github.com/skilstak/go-colors"
	"github.com/skilstak/go-input"
)

func main() {
	fmt.Println(c.Clear + c.B3 + "GOBATTLE")
	var player = input.Ask("What is your player name? >> " + c.Blue)
	fmt.Println(c.B3 + "Welcome " + player)
}
