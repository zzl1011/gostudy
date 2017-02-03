/**
 * author: zhaozhilu
 * 斐波纳契数列
 *
 */
package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		nums := fibonacci(i)
		fmt.Println(nums)
	}
}

func fibonacci(i int) int {
	if i < 2 {
		return i
	}

	return fibonacci(i-1) + fibonacci(i-2)
}
