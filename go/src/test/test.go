package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	time := time.Now().UnixNano()
	rand.Seed(time)
	fmt.Println("Test")
	here := rand.Intn(100)
	herea := strconv.Itoa(here)
	fmt.Println(herea)
}
