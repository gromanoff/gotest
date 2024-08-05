package main

import "fmt"

func main() {
	var a, b int
	fmt.Print("Введите первое число: ")
	fmt.Scanf("%d", &a)
	fmt.Print("Введите второе число: ")
	fmt.Scanf("%d", &b)

	fmt.Print("Введите действие (+, -, *, /): ")
	var action string
	fmt.Scanf("%s", &action)

	fmt.Printf("Результат: %d\n", step(a, b, action))
}

func step(a, b int, action string) int {
	result := 0

	switch action {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		if b != 0 {
			result = a / b
		} else {
			fmt.Println("Ошибка: деление на ноль")
			return 0
		}
	default:
		fmt.Println("Ошибка: неизвестное действие")
		return 0
	}

	return result
}
