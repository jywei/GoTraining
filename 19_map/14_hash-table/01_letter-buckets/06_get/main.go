package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("http://www.gutenberg.org/files/2701/old/moby10b.txt")
	if err != nil {
		log.Fatal(err)
	}
	bs, err := ioutil.ReadAll(res.Body) // bs == byte slice, res.Body == response body
	res.Body.Close()                    // after using a Get, we should close it
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", bs) // print it as a string
}
