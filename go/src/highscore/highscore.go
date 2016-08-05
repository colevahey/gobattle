package highscore

import (
	"encoding/json"
	"fmt"
	c "github.com/skilstak/go-colors"
	//"marshal"
	"strconv"
)

var Player = "Turtle"
var T = Top{"Player", 0, "red"}

type Top struct {
	Name   string
	Points int
	Color  string
}

//The _ is the error which comes out as <nil>
var b, _ = json.Marshal(T)
var Data Top
var _ = json.Unmarshal(b, &Data)

func Highscore() {
	fmt.Println(c.Clear + c.B3 + "High Score: " + Data.Name + ", " + strconv.Itoa(Data.Points))
}
