package main

import (
	"fmt"
        "os"
        "strconv"
)

func add(a int, b int) int {
	return a+b
}

func plus(a int, b int) int {
	return a-b
}

func multi(a int, b int) int {
	return a*b
}

func division(a int, b int) int {
	return a/b
}

func main() {
	var result int
        args := os.Args
        first, _ := strconv.Atoi(args[1])
	operation := args[2]
        second, _ := strconv.Atoi(args[3])

	switch operation {
            case "+" :
              result = add(first,second)

            case "-":
              result = plus(first,second)

            case "*":
             result  = multi(first,second)

            case "/":
              result = division(first,second)
            default:

        }
        fmt.Printf("%v%v%v的计算结果为%v",first,operation,second,result)
}
