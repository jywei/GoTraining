package main

import (
	"fmt"
	"log"
)

// NorgateMathError is a custom error type
type NorgateMathError struct {
	lat, long string // latitude and longitude
	err       error
}

func (n *NorgateMathError) Error() string { // this func turns the NorgateMathError struct into an error type interface
	return fmt.Sprintf("a norgate math error occured: %v %v %v", n.lat, n.long, n.err)
} // if one uses a pointer as an input(receiver), then the output(value) can only be a pointer

func main() {
	_, err := Sqrt(-10.23)
	if err != nil {
		log.Println(err)
	}
}

// Sqrt is a custom square function
func Sqrt(f float64) (float64, error) {
	if f < 0 {
		nme := fmt.Errorf("norgate math redux: square root of negative number: %v", f)
		return 0, &NorgateMathError{"50.2289 N", "99.4656 W", nme}
		// return 0, nme
	}
	// implementation
	return 42, nil
}

// see use of structs with error type in standard library:
// http://www.goinggo.net/2014/11/error-handling-in-go-part-ii.html
//
// http://golang.org/pkg/net/#OpError
// http://golang.org/src/pkg/net/dial.go
// http://golang.org/src/pkg/net/net.go
//
// http://golang.org/src/pkg/encoding/json/decode.go
