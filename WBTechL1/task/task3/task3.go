package main

import (
	"fmt"
	"sync"
)

func main() {
	slice := []int{2, 4, 6, 8, 10}      // Создаем слайс чисел из ТЗ
	intCh := make(chan int, len(slice)) // Создаем буферизированный канал для чисел
	sum := 0

	var wg sync.WaitGroup // Создаем wait group для синхронизации горутин
	for _, i := range slice {
		wg.Add(1)        // Увеличиваем счетчик wait group
		go func(d int) { // Анонимная функция для возведения чисел из слайса в квадрат
			defer wg.Done() // Уменьшаем счетчик wait group при завершении горутины
			intCh <- d * d  // Передаем результат в канал
		}(i) // Передаем в анонимную функцию текущее значение i
	}

	go func() { // Анонимная функция для закрытия канала после завершения всех горутин
		wg.Wait()    // Ждем завершения всех горутин
		close(intCh) // Закрываем канал
	}()

	for result := range intCh { // Читаем данные из канала пока он не будет закрыт
		sum += result // Суммируем результаты
	}
	fmt.Println(sum) // Выводим данные из канала в консоль
}
