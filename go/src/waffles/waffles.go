package main

import (
	"fmt"
	c "github.com/skilstak/go/colors"
	"github.com/skilstak/go/input"
)

func main() {
	var first = input.Ask(c.Clear + "Do you like waffles? ")
	if first == "yes" {
		fmt.Println(c.Multi("YES YOU LIKE WAFFLES"))
	} else if first == "no" {
		fmt.Println(c.Multi("YOU DON'T LIKE WAFFLES?"))
	} else {
		var second = input.Ask("Well do you like waffles? ")
		if second == "yes" {
			fmt.Println(c.Multi("YES YOU LIKE WAFFLES"))
		} else if second == "no" {
			fmt.Println(c.Multi("YOU DON'T LIKE WAFFLES?"))
		} else {
			fmt.Println("Maybe you just don't like waffles")
		}
	}
}
