package main

// Сумма чисел
// Напишите функцию, которая принимает два целых числа и возвращает их сумму.

import (
	"fmt"
)

func main() {
	a := 2
	b := 3

	fmt.Println(sum(a, b))
}

func sum(a, b int) int {
	result := a + b

	return result
}
