/**
 * author: zhaozhilu
 *
 *
 */
package main

import "fmt"

func main() {
	defer b()
	a := func() {
		fmt.Printf("hello\t")
	}
	a()
}

func b() {
	println("zhaozhilu")
}
