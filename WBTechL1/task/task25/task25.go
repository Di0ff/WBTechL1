package main

import (
	"fmt"
	"time"
)

func CustomSleep(d time.Duration) { // Функция принимает значения типа time.Duration
	// time.Duration - кол-во наносекунд
	start := time.Now() // В переменной start мы храним время с момента старта программы
	for {
		if time.Since(start) >= d { // Выполняется до тех пока time.Since не будет больше либо равно d
			// time.Since - возвращает временной интервал, прошедший с указанного момента времени
			break
		}
	}
}

func main() {
	CustomSleep(2 * time.Second) // Приостановка на 2 секунды
	fmt.Println("2 second later")
}
