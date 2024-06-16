package main

import (
	"fmt"
	"time"
)

func main() {
	const N = 5                // Время в секундах, через которое программа завершится
	dataChan := make(chan int) // Создаем канал для передачи данных

	go func() { // Горутинa для отправки значений в канал
		for i := 0; i < 10; i++ {
			fmt.Printf("Sending: %d\n", i)
			dataChan <- i
			time.Sleep(500 * time.Millisecond) // Имитация работы
		}
		close(dataChan) // Закрываем канал после отправки всех значений
	}()

	go func() { // Горутинa для чтения из канала
		for value := range dataChan {
			fmt.Printf("Received: %d\n", value)
		}
	}()

	timer := time.NewTimer(time.Duration(N) * time.Second) // Таймер на N секунд
	select {
	case <-timer.C:
		fmt.Println("Time's up! Exiting...")
	}

	time.Sleep(500 * time.Millisecond) // Дополнительная пауза для обработки всех оставшихся сообщений в канале
	fmt.Println("Program has finished.")
}
