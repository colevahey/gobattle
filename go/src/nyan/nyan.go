package main

import (
	"fmt"
	c "github.com/skilstak/go/colors"
	"math/rand"
	"strings"
)

func main() {
	for {
		spaces := strings.Repeat(" ", rand.Intn(20))
		fmt.Print(c.Rc() + "nyan" + spaces)
	}
}
