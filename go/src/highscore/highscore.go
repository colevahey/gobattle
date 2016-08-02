package highscore

import (
	"encoding/json"
	"fmt"
	c "github.com/skilstak/go-colors"
	"strconv"
)

var Name = ""
var Points = 0
var Color = ""

type Top struct {
	Name   string
	Points int
	Color  string
}

var t = Top{Name, Points, Color}

//The _ is the error which comes out as <nil>
var b, _ = json.Marshal(t)
var Data Top
var _ = json.Unmarshal(b, &Data)

func Highscore() {
	if Data.Color == "blue" {
		fmt.Println(c.Clear + c.B3 + "High Score: " + c.Blue + Data.Name + ", " + strconv.Itoa(Data.Points))
	} else if Data.Color == "red" {
		fmt.Println(c.Clear + c.B3 + "High Score: " + c.Red + Data.Name + ", " + strconv.Itoa(Data.Points))
	} else {
		fmt.Println("ERROR")
	}
}
