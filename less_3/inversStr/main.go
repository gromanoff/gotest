// 9. Инверсия строки
// Задача: Напишите функцию, которая принимает строку и возвращает её перевернутый вариант.

package main

import (
	"fmt"
)

func main() {
	s := "dlroW olleH"
	fmt.Printf(inverseStr(s))
}

func inverseStr(s string) string {
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
