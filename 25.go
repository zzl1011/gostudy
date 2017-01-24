/**
 * author ：zhaozhilu
 * package: Errorf
 *
 */
package main

import (
	"fmt"
)

func main() {
	const name, id = "zhaozhilu", 17
	err := fmt.Errorf("user %q(id %d) not found", name, id)
	if err != nil {
		fmt.Println(err) //这里不能使用printf
	}
}
