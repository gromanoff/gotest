// 6. Проверка наличия элемента в мапе
// Задача: Напишите функцию, которая проверяет, содержится ли заданный ключ в мапе.

package main

import (
	"fmt"
)

func main() {
	mapa := map[string]int{
		"apple":  5,
		"orange": 6,
		"juce":   11,
	}

	word := "juce"

	fmt.Println(ifExists(mapa, word))

}

func ifExists(m map[string]int, s string) string {
	for key := range m {
		if key == s {
			return "Есть такое"
		}
	}
	return "Такого нет"
}
