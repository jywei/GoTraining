package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First       string
	Last        string
	Age         int
	notExported int // not capitalized will not be exported
}

func main() {
	p1 := person{"James", "Bond", 35, 007}
	bs, _ := json.Marshal(p1) // byte slice, and marshal up
	fmt.Println(bs)
	fmt.Printf("%T \n", bs)
	fmt.Println(string(bs)) // covert the byte slice to a string
}
