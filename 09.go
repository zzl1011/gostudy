/**
 * author: zhaozhilu
 *
 *
 */
package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "asSASA ddd dsjkdsjs dk"
	fmt.Println(strings.Replace(str, "SA ", "abc", 1))
}
