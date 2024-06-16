package main

import "fmt"

func main() {
	// Самый простой способ
	a, b := 0, 1 // Инициализируем переменные
	a, b = b, a  // Меняем местами значения переменных
	fmt.Println(a, b)

	// Когда-то увидел в ютубе
	c, d := 10, 6
	c = c + d
	d = c - d
	c = c - d
	fmt.Println(c, d)
}
