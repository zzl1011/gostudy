/**
 * author: zhaozhilu
 *
 */
package main

import "fmt"

func main() {
	i := 0
	HERE:
	if i < 10 {
		fmt.Println(i)
		i++
		goto HERE
	}
}
