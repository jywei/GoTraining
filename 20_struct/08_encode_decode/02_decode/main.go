package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type person struct {
	First       string
	Last        string
	Age         int
	notExported int
}

func main() {
	var p1 person
	rdr := strings.NewReader(`{"First": "James", "Last": "Bond", "Age": 35}`) // NewReader will take a string and return a pointer
	json.NewDecoder(rdr).Decode(&p1)                                          // NewDecoder wants to take a reader, and return a pointer to Decoder

	fmt.Println(p1.First)
	fmt.Println(p1.Last)
	fmt.Println(p1.Age)
}
