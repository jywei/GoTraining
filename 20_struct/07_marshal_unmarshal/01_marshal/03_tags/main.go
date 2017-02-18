package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string `json:"-"` // don't include Last
	Age   int    `json:"wisdom score"`
}

func main() {
	p1 := person{"James", "Bond", 80}
	bs, _ := json.Marshal(p1)
	fmt.Println(string(bs))
}
