package main

// Сумма элементов в срезе
// Задача: Напишите функцию, которая принимает срез целых чисел и возвращает их сумму.

import "fmt"

func main() {
	arr := []int{12, 33, 22, 45, 87}

	fmt.Println(sumsrez(arr))

}

func sumsrez(x []int) int {
	sum := 0
	for _, num := range x {
		sum += num // Увеличиваем значение sum на значение текущего элемента num
	}
	return sum
}
