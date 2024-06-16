package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "needed longer no are #Queues"

	fmt.Println(FlipWord(str))
}

func FlipWord(str string) string {
	words := strings.Fields(str) // Разбиваем строку на слова

	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 { // Переворачиваем строку
		words[i], words[j] = words[j], words[i]
	}

	return strings.Join(words, " ") // Соединяем слова обратно в строку
}
