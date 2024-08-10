package main

import (
	"fmt"
)

func main() {
	arr := [5]int{16, 22, 33, 5, 6}
	x := 5
	fmt.Print(search(arr, x))
}

func search(nums [5]int, x int) bool {
	for i := 0; i < 5; i++ {
		if nums[i] == x {
			return true
		}
	}
	return false
}
