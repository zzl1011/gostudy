/**
* author: zhaozhilu
 *
 *
*/
package main

import "fmt"
import "strings"

func main() {
	for i := 1; i <= 100; i++ {
		fmt.Println(strings.Repeat("A", i))
	}
}
