package main

import (
	"context"
	"fmt"
	"sync/atomic"
	"time"
)

func stopWithChannel() { // Остановка горутины с использованием канала
	stop := make(chan struct{}) // Создаем канал для отправки сигнала остановки

	go func() { // Запускаем горутину
		for {
			select {
			case <-stop: // Считываем сигнал из канала
				fmt.Println("Goroutine stopped (channel)") // Выводим сообщение об остановке
				return                                     // Завершаем выполнение горутины
			default: // Если сигнала нет, продолжаем работу
				fmt.Println("Working (channel)...")
				time.Sleep(500 * time.Millisecond) // Имитация работы
			}
		}
	}()

	time.Sleep(2 * time.Second) // Ждем некоторое время перед отправкой сигнала остановки
	stop <- struct{}{}          // Отправляем сигнал для остановки горутины
	time.Sleep(1 * time.Second) // Даем немного времени для завершения горутины
}

func stopWithContext() { // Остановка горутины с контекст
	ctx, cancel := context.WithCancel(context.Background()) // Создаем контекст с возможностью отмены

	go func(ctx context.Context) { // Запускаем горутину с передачей контекста
		for {
			select {
			case <-ctx.Done(): // Проверяем, был ли контекст отменен
				fmt.Println("Goroutine stopped (context)")
				return
			default:
				fmt.Println("Working (context)...")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}(ctx)

	time.Sleep(2 * time.Second)
	cancel() // Отправляем сигнал отмены контекста
	time.Sleep(1 * time.Second)
}

func stopWithAtomic() { // Остановка горутины с атомарной переменной
	var stop int32 // Атомарная переменная для управления остановкой

	go func() {
		for {
			if atomic.LoadInt32(&stop) == 1 { // Проверяем значение атомарной переменной
				fmt.Println("Goroutine stopped (atomic)")
				return
			}
			fmt.Println("Working (atomic)...")
			time.Sleep(500 * time.Millisecond)
		}
	}()

	time.Sleep(2 * time.Second)
	atomic.StoreInt32(&stop, 1) // Устанавливаем флаг для остановки горутины
	time.Sleep(1 * time.Second)
}

func stopWithTimer() { // Остановка горутины с использованием таймера
	timer := time.NewTimer(2 * time.Second) // Создаем новый таймер, который сработает через 2 секунды

	go func() {
		for {
			select {
			case <-timer.C: // Проверяем, истекло ли время таймера
				fmt.Println("Goroutine stopped (timer)")
				return
			default:
				fmt.Println("Working (timer)...")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	time.Sleep(3 * time.Second) // Ждем некоторое время, чтобы убедиться, что горутина остановилась
}

func main() {
	fmt.Println("Stopping goroutine with channel...")
	stopWithChannel()

	fmt.Println("\nStopping goroutine with context...")
	stopWithContext()

	fmt.Println("\nStopping goroutine with atomic variable...")
	stopWithAtomic()

	fmt.Println("\nStopping goroutine with timer...")
	stopWithTimer()
}
