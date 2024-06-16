package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func produce(ch chan<- int, done <-chan struct{}) { // Непрерывно генерируем данные
	// и отправляем их в канал ch
	i := 0
	for {
		select {
		case <-done:
			fmt.Println("Producer stopped") // Получен сигнал завершения,
			// закрываем канал и выходим из функции
			close(ch)
			return
		case ch <- i:
			fmt.Printf("Produced: %d\n", i) // Отправляем сгенерированное значение в канал
			i++
			time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond) // Имитация работы с задержкой
			// от 0 до 500 миллисекунд
		}
	}
}

func worker(id int, ch <-chan int, wg *sync.WaitGroup) { // Прекращаем работу при закрытии канала ch
	defer wg.Done() // Уменьшаем счетчик WaitGroup при завершении функции
	for val := range ch {
		fmt.Printf("Worker %d: Consumed %d\n", id, val) // Читаем и выводим данные из канала
	}
	fmt.Printf("Worker %d stopped\n", id) // Выводим сообщение о завершении работы воркера
}

func main() {
	var numWorkers int
	fmt.Print("Enter number of workers: ")
	fmt.Scan(&numWorkers)

	ch := make(chan int) // Создаем канал для передачи данных

	done := make(chan struct{}) // Создаем канал для уведомления о завершении работы

	var wg sync.WaitGroup

	go produce(ch, done) // Запуск основного производителя данных в отдельной горутине

	for i := 0; i < numWorkers; i++ { // Запуск воркеров
		wg.Add(1) // Увеличиваем счетчик WaitGroup
		go worker(i, ch, &wg)
	}

	sigChan := make(chan os.Signal, 1) // Создаем канал для обработки сигналов завершения (Ctrl+C)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	<-sigChan // Ожидание сигнала завершения
	fmt.Println("Received termination signal, shutting down...")

	close(done) // Закрываем канал done

	wg.Wait() // Ожидание завершения всех воркеров
	fmt.Println("All workers stopped. Program terminated.")
}
