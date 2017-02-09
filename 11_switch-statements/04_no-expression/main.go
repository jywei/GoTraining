package main

import "fmt"

func main() {

	myFriendsName := "Mar"

	switch {
	case len(myFriendsName) == 2:
		fmt.Println("What a short name, like your penis!")
	case myFriendsName == "Jenny":
		fmt.Println("Jenny, you agian...")
	case myFriendsName == "Marcus", myFriendsName == "Medicci":
		fmt.Println("Marcus and Medicci, you bastards...")
	case myFriendsName == "Julien":
		fmt.Println("Julien, you agian...")
	case myFriendsName == "Alan":
		fmt.Println("Alan, you agian...")
	case myFriendsName == "Suzuki":
		fmt.Println("Suzuki, you agian...")
	default:
		fmt.Println("There is no one...")
	}
}

/*
  expression not needed
  -- if no expression provided, go checks for the first case that evals to true
  -- makes the switch operate like if/if else/else
  cases can be expressions
*/
