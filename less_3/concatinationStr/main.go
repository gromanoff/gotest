// 7. Конкатенация строк
// Задача: Напишите функцию, которая принимает срез строк и возвращает их конкатенацию в одну строку.

package main

import (
	"fmt"
)

func main() {
	s := []string{"hel", "lo", "w", "or", "l", "d"}

	fmt.Println(conctinationStr(s))
}

func conctinationStr(s []string) string {
	result := ""
	for i := 0; i < len(s); i++ {
		result += s[i]

	}
	return result
}
