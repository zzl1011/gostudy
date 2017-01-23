/**
 * author:zhaozhilu
 *
 *
 */
package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	a := "foobar"
	arr := strings.Split(a, "")
	for i := 0; i < len(arr); i++ {
		length := len(arr) - i - 1
		s += arr[length]
	}
	fmt.Println(s)
}
