package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string
	Age   int
}

func main() {
	var p1 person
	fmt.Println(p1.First) // print empty string
	fmt.Println(p1.Last)  // print empty string
	fmt.Println(p1.Age)   // print zero

	bs := []byte(`{"First": "James", "Last": "Bond", "Age": 35}`)
	json.Unmarshal(bs, &p1)

	fmt.Println("-----------------")
	fmt.Println(p1.First)
	fmt.Println(p1.Last)
	fmt.Println(p1.Age)
	fmt.Printf("%T \n", p1)
}
