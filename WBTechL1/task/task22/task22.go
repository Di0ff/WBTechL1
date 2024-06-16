package main

import (
	"fmt"
	"log"
)

func main() {
	var a, b int64          // Объявляем переменные
	a, b = 2097153, 1048577 // Задаем значения

	if !checkCondition(a) || !checkCondition(b) { // Проверяем, соответствуют ли значения переменных условию задачи
		// Если хотя бы одна переменная не соответствует условию,
		// программа завершает свою работу и сообщает об ошибке
		log.Fatalf("Values of a and b must be greater than 2^20")
	}

	// Выводим результаты операций
	fmt.Println(Sum(a, b))
	fmt.Println(Difference(a, b))
	fmt.Println(Multiplication(a, b))
	fmt.Println(Division(a, b))
}

func checkCondition(x int64) bool { // Функция для проверки значений
	const minVal int64 = 1 << 20 // Операция битового сдвига влево
	// Число 1 сдвигается влево на 20 бит, что эквивалентно 2^20 (1048576)

	if x < minVal { // Если значение меньше 2^20, возвращаем false
		return false
	}
	return true // В противном случае возвращаем true
}

func Sum(a, b int64) int64 { // Функция сложения чисел
	return a + b
}

func Difference(a, b int64) int64 { // Функция вычитания чисел
	return a - b
}

func Multiplication(a, b int64) int64 { // Функция умножения чисел
	return a * b
}

func Division(a, b int64) int64 { // Функция деления чисел
	return a / b
}
