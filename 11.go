/**
 * author:zhaozhilu
 *
 *
 */
package main

import "fmt"

func main() {
	a := "foobar"
	s := []rune(a)
	for i,j := 0, len(s)-1; i < j; i,j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	fmt.Printf("%s\n", string(s))
}
