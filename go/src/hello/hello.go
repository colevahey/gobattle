package main

import (
	"fmt"
	c "github.com/skilstak/go/colors"
)

func printPlain() {
	fmt.Println("Hello world")
}

func printMulti() {
	fmt.Println(c.Multi("Hello world"))
}

func printBlue() {
	fmt.Println(c.Blue + "Hello world")
}

func printRed() {
	fmt.Println(c.Red + "Hello world")
}

func printRandom() {
	fmt.Println(c.Rc() + "Hello world")
}

func main() {
	for {
		printPlain()
		printMulti()
		printBlue()
		printRed()
		printRandom()
	}
}
