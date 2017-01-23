/**
 * author: zhaozhilu
 *
 *
 */
package main

import "fmt"

var a = 6

func main() {
	p()
	q()
	p()
}

func p() {
	fmt.Println(a)
}

func q() {
//	a := 5
	a = 5
	fmt.Println(a)
}
