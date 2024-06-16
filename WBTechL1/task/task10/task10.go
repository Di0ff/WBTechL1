package main

import (
	"fmt"
	"sort"
)

func main() {
	temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	TempMap(temperatures)
	Temp(temperatures)
}

func TempMap(slice []float64) {
	groups := make(map[int][]float64) // Создаем карту для группировки температур

	for _, temp := range slice { // Группировка значений
		key := int(temp/10.0) * 10              // Получаем десятки
		groups[key] = append(groups[key], temp) // Вносим значения в свои группы
	}

	keys := make([]int, 0, len(groups)) // Так как мы используем мапы,
	// нам нужно отсортировать ключи,
	// чтобы значения всегда выводились в ожидаемом порядке
	for key := range groups {
		keys = append(keys, key)
	}
	sort.Ints(keys)

	for _, key := range keys { // Выводим результат
		fmt.Printf("%d: %v\n", key, groups[key])
	}
}

func Temp(slice []float64) { // Более глупый подход
	a := make([]float64, 0)
	b := make([]float64, 0)
	c := make([]float64, 0)
	d := make([]float64, 0)

	for _, i := range slice {
		switch {
		case i >= -30 && i < -20:
			a = append(a, i)
		case i >= 10 && i < 20:
			b = append(b, i)
		case i >= 20 && i < 30:
			c = append(c, i)
		case i >= 30 && i < 40:
			d = append(d, i)
		}
	}
	fmt.Printf("-20: %v\n", a)
	fmt.Printf("10: %v\n", b)
	fmt.Printf("20: %v\n", c)
	fmt.Printf("30: %v", d)
}
