// 8. Удаление дубликатов из среза
// Задача: Напишите функцию, которая удаляет дубликаты из среза строк.

package main

import (
	"fmt"
)

func main() {
	s := []string{"Atem", "apple", "tea", "banana", "Atem", "orange"}
	fmt.Println(DeleteDouble(s))
}

func DeleteDouble(s []string) []string {
	mapa := make(map[string]int)
	deleteDouble := []string{}
	element := ""

	for i := 0; i < len(s); i++ {
		element = s[i]
		mapa[element] = i
	}

	for word, _ := range mapa {
		deleteDouble = append(deleteDouble, word)
	}

	return deleteDouble
}
