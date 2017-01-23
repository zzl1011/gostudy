/**
 * author:zhaozhilu
 * struct
 */
package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	var P person

	P.name = "zzl"
	P.age = 30
	fmt.Printf("The person's name is %s\n", P.name)
	fmt.Printf("The person's age is %d", P.age)
}
