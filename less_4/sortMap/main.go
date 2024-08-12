// Задание 4: Сортировка мапы
// Создайте мапу с ключами-строками и значениями-целыми числами. Напишите функцию, которая возвращает слайс ключей, отсортированных по значениям в мапе.
// Для одинаковых значений ключи должны быть отсортированы в алфавитном порядке.

package main

import (
	"fmt"
	"sort"
)

type KeyValue struct {
	Key   string
	Value int
}

func main() {
	mapa := map[string]int{
		"apple":  11,
		"juice":  12,
		"orange": 13,
		"coffee": 14,
		"rat":    11,
	}

	fmt.Println(sortMap(mapa))
}

func sortMap(m map[string]int) []string {
	var kvSlice []KeyValue

	for key, value := range m {
		kvSlice = append(kvSlice, KeyValue{Key: key, Value: value})
	}

	sort.Slice(kvSlice, func(i, j int) bool {
		if kvSlice[i].Value == kvSlice[j].Value {
			return kvSlice[i].Key < kvSlice[j].Key
		}
		return kvSlice[i].Value < kvSlice[j].Value
	})

	var sortedKeys []string
	for _, kv := range kvSlice {
		sortedKeys = append(sortedKeys, kv.Key)
	}

	return sortedKeys
}
