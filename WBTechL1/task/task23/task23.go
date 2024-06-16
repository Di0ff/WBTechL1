package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 10}
	i := 4
	fmt.Println(Remove(slice, i))
	fmt.Println(RemoveAppend(slice, i))
}

func Remove(slice []int, i int) []int {
	return append(slice[:i], slice[i+1:]...) // Первый аргумент — срез до i-ого элемента,
	// а второй — срез после i-ого элемента
	// Оператор "..." разворачивает второй срез так, что его элементы добавляются по отдельности
}

func RemoveAppend(slice []int, i int) []int {
	newSlice := []int{}       // Создаем результирующий срез
	for j, v := range slice { // "Бежим" по всему исходному срезу
		if j != i { // Если индексы не совпадают,
			newSlice = append(newSlice, v) // то вставляем значение в результирующий срез
		}
	}
	return newSlice // Возвращаем результирующий срез
}
