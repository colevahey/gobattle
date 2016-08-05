package marshal

import (
	"encoding/json"
	"fmt"
	"highscore"
)

var T = Top{highscore.Player, 0, "red"}

type Top struct {
	Name   string
	Points int
	Color  string
}

//The _ is the error which comes out as <nil>
var b, _ = json.Marshal(T)
var Data Top
var _ = json.Unmarshal(b, &Data)

func main() {
	fmt.Println(Data.Name)
}
