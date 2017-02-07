/**
 * author:zhaozhilu
 * strings包
 * Contains方法
 */

package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string = "Hello world!"
	fmt.Printf("T/F? \"%s\" is have %s \n", str, "o")
	fmt.Printf("%t\n", strings.Contains(str, "o"))
	fmt.Printf("T/F? \"%s\" is have %s \n", str, "g")
	fmt.Printf("%t\n", strings.Contains(str, "g"))

}
