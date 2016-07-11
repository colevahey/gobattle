package main

import (
	"fmt"
	c "github.com/skilstak/go/colors"
	"time"
)

func main() {
	fmt.Print(c.Clear)
	var num = 1
	var n = "1"
	for num <= 12 {
		if num == 1 {
			n = "1"
		} else if num == 2 {
			n = "2"
		} else if num == 3 {
			n = "3"
		} else if num == 4 {
			n = "4"
		} else if num == 5 {
			n = "5"
		} else if num == 6 {
			n = "6"
		} else if num == 7 {
			n = "7"
		} else if num == 8 {
			n = "8"
		} else if num == 9 {
			n = "9"
		} else if num == 10 {
			n = "10"
		} else if num == 11 {
			n = "11"
		} else if num == 12 {
			n = "12"
		} else {
			fmt.Println("FATAL ERROR")
		}
		fmt.Println(c.Multi(n + " badger"))
		time.Sleep(150 * time.Millisecond)
		num++
	}
}
