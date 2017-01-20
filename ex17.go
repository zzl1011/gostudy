/**
 * author:zhaozhilu
 *
 *
 */
package main

import "fmt"

func SumAndProduct(A, B int) (int, int) {
	return A + B, A * B
}

func main() {
	x := 3
	y := 4

	sum, product := SumAndProduct(x, y)
	fmt.Printf("%d + %d = %d\n", x, y, sum)
	fmt.Printf("%d * %d = %d\n", x, y, product)
}
