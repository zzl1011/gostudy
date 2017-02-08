/**
 * author:zhaozhilu
 * strings包Repeat、ToLower、ToUpper方法
 *
 */
package main

import (
	"fmt"
	"strings"
)

func main() {
	var oldStr string = "Hi there! "
	var newStr string
	var lowerStr string
	var upperStr string

	newStr = strings.Repeat(oldStr, 3)
	fmt.Printf("The new repeated string is: %s\n", newStr)

	lowerStr = strings.ToLower(oldStr)
	fmt.Printf("The new lower string is: %s\n", lowerStr)

	upperStr = strings.ToUpper(oldStr)
	fmt.Printf("The new upper string is: %s\n", upperStr)

}
