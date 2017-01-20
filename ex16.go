/**
 * author:zhaozhilu
 *
 *
 */
package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}

	return b

}

func main() {
	x := 4
	y := 5
	z := 6

	max_xy := max(x, y)
	max_yz := max(y, z)

	fmt.Printf("max(%d, %d) = %d\n", x, y, max_xy)
	fmt.Printf("max(%d, %d) = %d\n", y, z, max_yz)
	fmt.Printf("max(%d, %d) = %d\n", x, z, max(x,z))
}
