/**
 * author: zhaozhilu 
 *
 *
 */
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	b := []byte("asSASA ddd dsjkdsjs dk时间")
	var total int
	for len(b) > 0 {
		_, size := utf8.DecodeRune(b)
		//fmt.Printf("%c %v\n", r, size)
		total += size
		b = b[size:]
	}
	fmt.Println(total)
}
