package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 5, 6, 6, 7, 8, 9, 0, -1, 10}
	n := len(arr)

	QuickSort(arr, 0, n-1)
	fmt.Println(arr)
}

func medianOfThree(arr []int, low, mid, high int) int { // Выбираем медианный элемент
	if arr[low] > arr[mid] {
		arr[low], arr[mid] = arr[mid], arr[low]
	}
	if arr[mid] > arr[high] {
		arr[mid], arr[high] = arr[high], arr[mid]
	}
	if arr[low] > arr[mid] {
		arr[low], arr[mid] = arr[mid], arr[low]
	}
	return mid
	// Чтобы избежать случайностей при выборе опорного элемента
	// и не потерять в производительности, если массив уже частично отсортирован,
	// лучше всего выбрать медиану из трех: начала, середины и конца
}

func partition(arr []int, low, high int) int { // Находим индекс опорного элемента
	mid := low + (high-low)/2                               // Середина массива
	pivotIndex := medianOfThree(arr, low, mid, high)        // Индекс опорного элемента
	pivot := arr[pivotIndex]                                // Опорный элемент
	arr[pivotIndex], arr[high] = arr[high], arr[pivotIndex] // Перемещаем опорный элемент в конец
	i := low - 1                                            // Индекс меньшего элемента

	for j := low; j < high; j++ {
		if arr[j] < pivot { // Если текущий элемент меньше опорного,
			i++
			arr[i], arr[j] = arr[j], arr[i] // то меняем элементы местами
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1] // Обмен опорного элемента с элементом после i
	return i + 1                              // возвращаем индекс опорного элемента
}

func QuickSort(arr []int, low, high int) {
	if low < high {
		pivot := partition(arr, low, high)

		QuickSort(arr, low, pivot-1)  // Рекурсивно сортируем элементы до опорного
		QuickSort(arr, pivot+1, high) // и после опорного
	}
}
