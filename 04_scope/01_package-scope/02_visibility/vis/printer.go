package vis

import "fmt"

// in the package vis, i can access to the variable i defined in name.go

// PrintVar is exported because it starts with a capital letter
func PrintVar() {
	fmt.Println(MyName)
	fmt.Println(yourName)
}
