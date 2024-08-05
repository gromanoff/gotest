package main

import "fmt"

func main() {
	x := 5
	fmt.Println(even(x))

}

func even(x int) string {
	if x <= 0 {
		return "Число не может быть нулем или меньше нуля"
	} else if x%2 == 0 {
		return "Это число четное!"
	} else {
		return "Это число нечетное!"
	}
}
