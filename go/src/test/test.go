package main

import (
	"encoding/json"
	"fmt"
	"github.com/skilstak/go-input"
)

//func Marshal(v interface{}) ([]byte, error)
//func Unmarshal(data []byte, v interface{}) error

type Message struct {
	Name string
	Age  string
}

func main() {
	namer := input.Ask("What is your name? ")
	ager := input.Ask("How old are you? ")
	m := Message{namer, ager}
	b, _ := json.Marshal(m)
	var data Message
	var _ = json.Unmarshal(b, &data)
	fmt.Println("Name: " + data.Name)
	fmt.Println("Age: " + data.Age)
}
