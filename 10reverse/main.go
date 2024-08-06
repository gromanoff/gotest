package main

import "fmt"

// Обратный порядок элементов в срезе
// Задача: Напишите функцию, которая переворачивает срез строк.

func main() {
	arr := [5]string{"home", "apple", "avacado", "bal"}
	fmt.Println(reverseArr(arr[:]))
}

func reverseArr(arr []string) []string {
	left := 0
	right := len(arr) - 1

	for left < right {
		arr[left], arr[right] = arr[right], arr[left]
		left++
		right--
	}

	return arr
}
