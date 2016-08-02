package main

import (
	"encoding/json"
	"fmt"
	c "github.com/skilstak/go-colors"
)

type Top struct {
	Name  string
	Score string
	Color string
}

var t = Top{"Jose", "12", "red"}

//The _ is the error which comes out as <nil>
var b, _ = json.Marshal(t)
var Data Top
var _ = json.Unmarshal(b, &Data)

func main() {
	if Data.Color == "blue" {
		fmt.Println(c.Clear + c.B3 + "High Score: " + c.Blue + Data.Name + ", " + Data.Score)
	} else if Data.Color == "red" {
		fmt.Println(c.Clear + c.B3 + "High Score: " + c.Red + Data.Name + ", " + Data.Score)
	} else {
		fmt.Println("ERROR")
	}
}
