package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	target := 5

	result := BinarySearch(arr, target)
	fmt.Println(result)
}

func BinarySearch(arr []int, target int) int {
	low := 0 // Обозначаем рамки начала и конца бинарного поиска
	high := len(arr) - 1

	for low <= high {
		mid := low + (high-low)/2 // Находим индекс среднего элемента
		if arr[mid] == target {   // Сравниваем средний и нужный элементы,
			return mid // чтобы понять, в какую сторону сдвигать указатели
		}
		if arr[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1 // Возвращаем -1, если элемент не найден
}
