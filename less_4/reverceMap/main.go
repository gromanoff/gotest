// Задание 2: Реверс мапы
// Напишите функцию, которая принимает мапу с ключами-строками и значениями-целыми числами и возвращает новую мапу, где ключами являются значения оригинальной мапы,
// а значениями — слайсы ключей. Убедитесь, что ваша функция корректно обрабатывает ситуации с одинаковыми значениями в исходной мапе.

package main

import (
	"fmt"
)

func main() {
	mapa := map[string]int{
		"apple":  11,
		"juice":  12,
		"orange": 13,
		"coffee": 14,
		"rat":    11,
	}

	fmt.Println("Original map:", mapa)

	reversed := reverceMap(mapa)
	fmt.Println("Reversed map:", reversed)
}

func reverceMap(mapa map[string]int) map[int][]string {

	result := make(map[int][]string)

	for key, value := range mapa {
		result[value] = append(result[value], key)
	}

	return result
}
