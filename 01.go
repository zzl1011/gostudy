/**
 * @author: zhaozhilu
 * @source: https://leetcode.com/problems/two-sum/
 *
 */
package main

import "fmt"

func main() {
	var result [2]int
	target := 26
	nums := []int{2, 7, 11, 15}
	for i := 0; i <= len(nums)-1; i++ {
		for j := len(nums) - 1; i < j; j-- {
			if nums[i]+nums[j] == target {
				fmt.Printf("The one number is %d\n, the two number is %d\n", i, j)
				result[0] = i
				result[1] = j
			}
		}
	}
	fmt.Println(result)
}
