/**
 * author:zhaozhilu
 * init函数
 *
 */
package main

import (
	"fmt"
	"math"
)

var Pi float64

func init() {
	Pi = 4 * math.Atan(1) //init() function computes pi
}

func main() {
	fmt.Println("Pi is", Pi)
}
