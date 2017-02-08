/**
 * author:zhaozhilu
 * strings包TrimSpace、Trim、TrimLeft、TrimRight、Split、Fields、
 *
 */
package main

import (
	"fmt"
	"strings"
)

func main() {
	var originStr string = " This is a space test "
	var str string = "@sfsfsf@"

	fmt.Printf("TrimSpace:%s\n", strings.TrimSpace(originStr))
	fmt.Printf("TrimLeft: %s\n", strings.TrimLeft(str, "@"))
	fmt.Printf("TrimRight: %s\n", strings.TrimRight(str, "@"))

	for _, v := range strings.Split(originStr, " ") {
		fmt.Println(v)
	}
}
