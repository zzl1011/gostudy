/**
 * author: zhaozhilu 
 * 
 */
package main

import (
	"fmt"
)
func main() {
	var sum [3]int
	arr1 := [3]int{2, 4, 3}
	arr2 := [3]int{5, 6, 4}
	carry := 0
	for i := 0; i < len(arr1); i++ {
		tmp := arr1[i] + arr2[i]
		if tmp % 5 == 0 && tmp % 2 == 0 {
			sum[i] = 0
			carry = 1
		} else {
			if carry == 1 {
				sum[i] = tmp + 1
				carry = 0
			} else {
				sum[i] = tmp
			}
		}

	}
	fmt.Println(sum)
}
