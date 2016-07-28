package main

import (
	"encoding/json"
	"fmt"
	"github.com/skilstak/go-input"
	"strconv"
)

//func Marshal(v interface{}) ([]byte, error)
//func Unmarshal(data []byte, v interface{}) error

type Message struct {
	Name string
	Body string
	Time int
}

func main() {
	namer := input.Ask("What is your name?")
	m := Message{namer, "Hello", 12}
	b, _ := json.Marshal(m)
	var here Message
	var _ = json.Unmarshal(b, &here)
	fmt.Println("Name: " + here.Name)
	fmt.Println("Body: " + here.Body)
	fmt.Println("Time: " + strconv.Itoa(here.Time))
}
