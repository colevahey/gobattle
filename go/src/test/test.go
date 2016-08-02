package main

import (
	"fmt"
	"hehe"
)

//func Marshal(v interface{}) ([]byte, error)
//func Unmarshal(data []byte, v interface{}) error

type Message struct {
	Name string
	Age  string
}

func main() {
	fmt.Println("Name: " + hehe.Data.Name)
	fmt.Println("Score: " + hehe.Data.Age)
	fmt.Println(hehe.This)
}
