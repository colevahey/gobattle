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
	fmt.Println(c.B3 + "Welcome " + c.Green + strings.Title(player))
}
