package main

import "fmt"

type Base struct{ 
	id int
}

func (b *Base) Id() int {
	return b.id
}

func (b *Base) SetId(newId int) {
	b.id = newId
}

type Person struct {
	Base
	FirstName string
	LastName  string
}

type Employee struct {
	Person
	salary string
}

func main() {
	e := new(Employee)
	e.SetId(5)
	fmt.Printf("get id value: %d\n", e.Id())
}
