/**
 * author:zhaozhilu
 * strings包Index、LastIndex、Replace、 Count方法
 *
 */
package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string = "Hello, how is it going, Hugo?"
	var manyG = "ggggggggg"

	fmt.Printf("Index of H's in %s is: %d\n", str, strings.Index(str, "H"))
	fmt.Printf("Last index of g's in %s is: %d\n", str, strings.LastIndex(str, "g"))
	fmt.Printf("Number of g's in %s is: %d\n", manyG, strings.Count(manyG, "g"))
	fmt.Println(strings.Replace(manyG, "g", "s", 2))
	fmt.Println(strings.Replace(manyG, "g", "m", -1))
}
