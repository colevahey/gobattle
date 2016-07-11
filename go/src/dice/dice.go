package main

import (
	"fmt"
	c "github.com/skilstak/go/colors"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	var num = rand.Intn(5)
	if num == 0 {
		fmt.Println(c.Clear + c.Rc() + "1")
	} else if num == 1 {
		fmt.Println(c.Clear + c.Rc() + "2")
	} else if num == 2 {
		fmt.Println(c.Clear + c.Rc() + "3")
	} else if num == 3 {
		fmt.Println(c.Clear + c.Rc() + "4")
	} else if num == 4 {
		fmt.Println(c.Clear + c.Rc() + "5")
	} else if num == 5 {
		fmt.Println(c.Clear + c.Rc() + "6")
	} else {
		fmt.Println(c.Clear + c.Rc() + "Fatal error")
	}
}
