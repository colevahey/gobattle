package hehe

import (
	"encoding/json"
)

var This = "IT WORKS"

type Message struct {
	Name string
	Age  string
}

var m = Message{"Jose", "12"}

//The _ is the error which comes out as <nil>
var b, _ = json.Marshal(m)
var Data Message
var _ = json.Unmarshal(b, &Data)
