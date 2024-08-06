package main

import "fmt"

// Поиск максимального элемента в срезе
// Задача: Напишите функцию, которая возвращает максимальное значение из среза целых чисел.

func main() {
	slice := []int{200, 23, 55, 88, 18, 100, 400}

	fmt.Println(maxElement(slice))
}

func maxElement(arr []int) int {
	max := arr[0]
	for i := 0; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return max
}
