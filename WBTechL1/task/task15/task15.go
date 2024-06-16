/*
var justString string <- Глобальная переменная это не есть хорошо, могут возникнуть проблемы
func someFunc() { <- Функция ничего не возвращает
v := createHugeString(1 << 10) <- Функция не объявлена, нужно её создать
justString = v[:100] <- Если строка меньше 100 символов, может возникнуть ошибка
}

func main() {
someFunc() <- Просто выполняется, никуда не выводится
}
*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

const CHARS = "abcdefghijklmnopqrstuvwxyz" // Алфавит

func Create(n int) string { // Создаем строку из случайных рун заданного размера
	rand.Seed(time.Now().UnixNano())
	runes := make([]rune, n)
	for i := range runes {
		runes[i] = rune(CHARS[rand.Intn(len(CHARS))])
	}
	return string(runes)
}

func Trim(n int, str string) string { // Обрезаем строку до заданной длины
	if len(str) < n {
		return str
	}
	return str[:n]
}

func main() {
	v := Create(1 << 10)
	justString := Trim(100, v)
	fmt.Println(justString)
}
