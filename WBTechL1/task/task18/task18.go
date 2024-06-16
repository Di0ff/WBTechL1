package main

import (
	"fmt"
	"sync"
)

type Counter struct { // Структура-счетчик
	mu    sync.Mutex // Используем мьютекс для предотвращения гонки данных (data race)
	value int
}

func (c *Counter) Increment() { // Увеличиваем значение счетчика на 1
	c.mu.Lock()         // Блокируем мьютекс, чтобы доступ был только у одного потока
	defer c.mu.Unlock() // Разблокируем мьютекс после выполнения функции
	c.value++           // Увеличиваем значение поля структуры на 1
}

func (c *Counter) Value() int { // Возвращаем текущее значение счетчика
	c.mu.Lock()         // Блокируем мьютекс, чтобы безопасно прочитать значение
	defer c.mu.Unlock() // Разблокируем мьютекс после выполнения функции
	return c.value
}

func main() {
	var wg sync.WaitGroup // Создаем WaitGroup для синхронизации горутин
	counter := &Counter{} // Создаем указатель на экземпляр структуры Counter

	numGoroutines := 1000 // Количество горутин

	for i := 0; i < numGoroutines; i++ { // Инкрементируем счетчик в нескольких горутинах
		wg.Add(1) // Увеличиваем счетчик WaitGroup
		go func() {
			defer wg.Done()     // Уменьшаем счетчик WaitGroup при завершении горутины
			counter.Increment() // Вызываем метод Increment
		}()
	}

	wg.Wait() // Ждем завершения всех горутин

	fmt.Println(counter.Value()) // Выводим итоговое значение счетчика
}
