package main

import (
	"fmt"
	"math"
)

type Point struct { // Структура Point с инкапсулированными полями
	x int
	y int
}

func main() {
	point1 := NewPoint(10, 20)
	point2 := NewPoint(30, 40)
	fmt.Printf("%.2f\n", point1.Distance(point2)) // Вывод результата с точностью до 2 знаков после запятой
}

func NewPoint(x, y int) Point { // Конструктор для структуры Point
	return Point{x: x, y: y}
}

func (p Point) Distance(other Point) float64 { // Метод для нахождения расстояния между двумя точками
	dx := float64(other.x - p.x)    // Вычисляем разность координат по оси x между двумя точками и преобразуем в float64
	dy := float64(other.y - p.y)    // Вычисляем разность координат по оси y между двумя точками и преобразуем в float64
	return math.Sqrt(dx*dx + dy*dy) // Вычисляем Евклидово расстояние между двумя точками по формуле sqrt(dx^2 + dy^2)
}
