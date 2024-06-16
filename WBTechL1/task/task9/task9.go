package main

import (
	"fmt"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5} // Исходный массив чисел

	ch1 := make(chan int) // Создаем каналы
	ch2 := make(chan int)

	go func() { // Горутин для записи чисел в первый канал
		for _, num := range numbers {
			ch1 <- num
		}
		close(ch1)
	}()

	go func() { // Горутин для чтения чисел из первого канала
		// и записи во второй канал
		for num := range ch1 {
			ch2 <- num * 2
		}
		close(ch2)
	}()

	for result := range ch2 { // Чтение из второго канала и вывод в stdout
		fmt.Println(result)
	}
}
