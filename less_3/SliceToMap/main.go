// 5. Подсчет частоты элементов в срезе
// Задача: Напишите функцию, которая подсчитывает количество вхождений каждого элемента в срезе строк и возвращает результат в виде мапы.

package main

import "fmt"

func main() {
	srez := []string{"one", "two", "three", "two", "five", "one", "one"}
	fmt.Println(countWords(srez))
}

func countWords(s []string) map[string]int {
	words := make(map[string]int)

	for _, word := range s {
		words[word] += 1
	}

	return words
}
