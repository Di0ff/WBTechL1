package main

import (
	"fmt"
)

func main() {
	str := "wbtech4"
	str1 := "вбтех4"

	fmt.Println(FlipSimple(str))
	fmt.Println(FlipSimple(str1))
	fmt.Println(FlipConcat(str))
	fmt.Println(FlipConcat(str1))
	fmt.Println(FlipAppend(str))
	fmt.Println(FlipAppend(str1))
	fmt.Println(FlipRecursive(str))
	fmt.Println(FlipRecursive(str1))
}

func FlipSimple(str string) string {
	runes := []rune(str)                                  // Преобразуем строку в срез рун, т.к. символы могут быть Unicode
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 { // Запускаем два "указателя" с начала и конца
		runes[i], runes[j] = runes[j], runes[i] // Меняем символы местами
	}
	return string(runes) // Возвращаем результирующую строку, преобразовав срез рун в строку
}

func FlipConcat(str string) string {
	runes := []rune(str)
	var result string // Инициализируем результирующую переменную
	for i := len(runes) - 1; i >= 0; i-- {
		result += string(runes[i]) // На каждой итерации цикла преобразуем руну в строку и добавляем в результирующую
	}
	return result
}

func FlipAppend(str string) string {
	runes := []rune(str)
	var result []rune
	for i := len(runes) - 1; i >= 0; i-- {
		result = append(result, runes[i]) // Используем append для добавления рун в слайс
		// Append работает быстрее, чем конкатенация строк, потому что append использует динамическое выделение памяти и
		// добавление элементов к слайсу, тогда как конкатенация строк создает новую строку на каждой итерации,
		// что приводит к многократному выделению и копированию памяти
	}
	return string(result)
}

func FlipRecursive(str string) string {
	if len(str) == 0 { // Если строка пустая, возвращаем её
		return str
	}
	runes := []rune(str)
	return string(runes[len(runes)-1]) + FlipRecursive(string(runes[:len(runes)-1]))
	// Рекурсивно вызываем FlipRecursive для подстроки без последней руны
	// и добавляем последнюю руну к результату
}
