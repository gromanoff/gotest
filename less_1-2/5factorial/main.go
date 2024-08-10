package main

import (
	"fmt"
)

func main() {
	a := 5
	fmt.Println(factorial(a))

}

func factorial(x int) int {
	if x == 0 || x == 1 {
		return 1
	}
	return x * factorial(x-1)
}
