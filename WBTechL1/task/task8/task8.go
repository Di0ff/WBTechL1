package main

import (
	"fmt"
)

func setBit(n int64, i uint, bit bool) int64 { // setBit устанавливает i-й бит в значение 1 или 0
	if bit {
		return n | (1 << i) // Установить i-й бит в 1
	} else {
		return n &^ (1 << i) // Установить i-й бит в 0
	}
}

func main() {
	var n int64 = 42
	var i uint = 3
	var bit = true

	n = setBit(n, i, bit)
	fmt.Println(n)

	bit = false
	n = setBit(n, i, bit)
	fmt.Println(n)
}
