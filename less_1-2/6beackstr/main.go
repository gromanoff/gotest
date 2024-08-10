package main

import (
	"fmt"
)

func main() {
	s := "Hello"
	fmt.Println(beackstr(s))
}

func beackstr(s string) string {

	runes := []rune(s)
	left := 0
	right := len(runes) - 1

	for left < right {
		runes[left], runes[right] = runes[right], runes[left]
		left++
		right--
	}

	return string(runes)
}
