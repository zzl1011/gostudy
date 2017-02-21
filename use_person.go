package main

import (
	"fmt"
	"./person"
)

func main() {
	p := new(person.Person)
	//fmt.Println(p.firstName)

	p.SetFirstName("Eric")
	fmt.Println(p.FirstName())
}
