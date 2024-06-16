package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "aaaa"
	str1 := "abcd"
	str2 := "abCdefAaf"
	str3 := "aabcd"
	fmt.Println(CFU(str))
	fmt.Println(CFU(str1))
	fmt.Println(CFU(str2))
	fmt.Println(CFU(str3))
}

func CFU(str string) bool { // Checking for uniqueness
	if len(str) == 0 { // Если на вход подается пустая строка, она априори считается уникальной
		return true
	}

	lowerStr := strings.ToLower(str) // По тз функция должна быть регистронезависимой,
	// поэтому необходимо изменить регистр всех символов в строке
	seen := make(map[rune]bool) // Создаем пустую мапу для хранения символов

	for _, char := range lowerStr { // Бежим по каждому символу строки
		if seen[char] { // Если такой символ есть в мапе, возвращаем false и завершаем выполнение программы
			return false
		}
		seen[char] = true // Иначе добавляем символ в мапу
	}
	return true // Если все символы уникальны, возвращаем true
}
