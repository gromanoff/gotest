package main

// Подсчет слов в строке
// Задача: Напишите функцию, которая принимает строку и возвращает количество слов в ней.

import (
	"fmt"
	"strings"
)

func main() {
	s := "You are gey"
	fmt.Println(lenStr(s))
}

func lenStr(s string) int {
	arr := strings.Split(s, " ")
	result := len(arr)

	return result
}
