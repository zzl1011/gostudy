/**
 * author:zhaozhilu
 * int can be replaced with any type
 *
 */
package main

import "fmt"

func main() {
	for i := range (*[10]int)(nil) {
		fmt.Println(i)
	}
	for i := range [10][0]int{} {
		fmt.Println(i)
	}
	for i := range [10]struct{}{} {
		fmt.Println(i)
	}
}
