package main

import (
	"encoding/json"
	"os"
)

type person struct {
	First       string
	Last        string
	Age         int
	notExported int
}

func main() {
	p1 := person{"James", "Bond", 20, 007}
	json.NewEncoder(os.Stdout).Encode(p1) // os.Stdout == writer
	// NewEncoder(os.Stdout) will take a write and return a encoded pointer (*Encoder)
	// Stdout will call Write method, which is attached to the file (Encode(p1))
}
