package main

import (
	"fmt"
	"sort"
)

func main() {
	slice1 := []int{12, 45, 7, 22, 19, 30, 3, 8, 15, 27, 11, 36, 23, 1, 40, 33, 17, 9, 5, 28}
	slice2 := []int{14, 3, 18, 21, 6, 37, 10, 25, 29, 4, 32, 19, 8, 2, 31, 20, 11, 38, 27, 16, 3}
	result := IntersectionAndSort(slice1, slice2)
	fmt.Println(Cut(result))
	fmt.Println(IntersectionMap(slice1, slice2))
}

func IntersectionAndSort(set1, set2 []int) []int { // Решение задачи по принципам функционального программирования
	sort.Ints(set1) // Сортируем оба слайса это нам пригодится чуть позже
	sort.Ints(set2)
	result := make([]int, 0) // Создаем результирующий слайс

	for i := 0; i < len(set1); i++ { // Бежим по всем элементам слайса set1
		for j := 0; j < len(set2); j++ { // и set2
			if set1[i] == set2[j] { // Если встречается пересечение множеств,
				result = append(result, set1[i]) // то записываем их в результирующий слайс
				j = len(set2) - 1                // и заканчиваем итерацию внутреннего цикла
			}
		}
	}
	return result // Возвращаем результирующий слайс
}

func Cut(slice []int) []int { // Удаляем повторяющиеся элементы из результирующего слайса
	if len(slice) == 0 { // Проверяем длину слайса, если нет пересечений в двух множествах,
		return slice // то результирующий слайс пустой и вырезать нечего
	}

	result := make([]int, 0) // Создаем конечный слайс

	for i := 0; i < len(slice)-1; i++ { // Проходим по всем элементам результирующего слайса, кроме последнего
		if slice[i] != slice[i+1] { // Если текущий элемент не равен следующему,
			result = append(result, slice[i]) // добавляем его в конечный слайс
		}
	}
	result = append(result, slice[len(slice)-1]) // Добавляем последний элемент

	return result // Возвращаем конечный слайс
}

func IntersectionMap(set1, set2 []int) []int { // Решение задачи мапой
	setMap := make(map[int]bool) // Создаем мапу для хранения элементов из первого множества
	for _, v := range set1 {     // Заполняем мапу элементами из set1
		setMap[v] = true
	}

	result := make([]int, 0) // Создаем результирующий слайс
	for _, v := range set2 { // Проверяем каждый элемент из set2
		if setMap[v] { // Если элемент существует в мапе,
			result = append(result, v) // добавляем его в результирующий слайс
			delete(setMap, v)          // Удаляем элемент из мапы, чтобы избежать дубликатов в результате
		}
	}

	return result // Возвращаем результирующий слайс
}
