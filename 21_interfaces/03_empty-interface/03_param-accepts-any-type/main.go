package main

import "fmt"

type animal struct {
	sound string
}

type dog struct {
	animal
	friendly bool
}

type cat struct {
	animal
	annoying bool
}

func specs(a interface{}) { // for empty interface you can throw anything in it (all types)
	fmt.Println(a)
}

func main() {
	fido := dog{animal{"woof"}, true}
	fifi := cat{animal{"meow"}, true}
	specs(fido)
	specs(fifi)
}
