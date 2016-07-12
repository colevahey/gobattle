package main

import (
	"fmt"
	c "github.com/skilstak/go-colors"
	"github.com/skilstak/go-input"
)

func main() {
	fmt.Println(c.Clear + c.Red + "GOBATTLE")
	var player = input.Ask("What is your player name? >> " + c.Magenta)
	fmt.Println(c.Red + "Welcome " + player)
}
