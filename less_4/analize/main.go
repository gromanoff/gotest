// Задание 1: Частотный анализ текста
// Напишите функцию, которая принимает строку и возвращает мапу, где ключами являются слова, а значениями — количество их вхождений в текст.
// Убедитесь, что функция не чувствительна к регистру и игнорирует знаки препинания.

package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "After a long day at work, John decided to relax by reading a book, sipping some tea, and listening to his favorite music."
	analize(s)
}

func analize(s string) {
	cleanStr := ""
	// чистим строку от знаков препенания.
	for _, char := range s {
		if char != '.' && char != ',' && char != '!' && char != '?' {
			cleanStr += string(char)
		}
	}
	// разделяем слова по пробелам и добавляем их в масив
	words := strings.Split(cleanStr, " ")
	mapa := make(map[string]int)
	// добавляем слова из масива в мапу если такого слова нет, если слово есть прибавляем 1 к его значению
	for i := 0; i < len(words); i++ {
		word := words[i]
		if _, exists := mapa[word]; exists {
			mapa[word]++
		} else {
			mapa[word] = 1
		}
	}

	fmt.Println(mapa)
}
