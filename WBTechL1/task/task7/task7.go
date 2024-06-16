package main

import (
	"fmt"
	"sync"
)

type SafeMap struct { // Потокобезопасная мапа
	mu sync.Mutex
	m  map[string]interface{}
}

func main() {
	safeMap := NewSafeMap() // Создаем экземпляр структуры
	wg := sync.WaitGroup{}  // Создаем WaitGroup для ожидания завершения всех горутин

	for i := 0; i < 10; i++ { // Цикл для заполнения мапы значениями от 0 до 9
		wg.Add(1)
		go func(i int) { // Запускаем горутину для конкурентной записи в мапу
			defer wg.Done() // Уменьшаем счетчик WaitGroup при завершении горутины
			key := fmt.Sprintf("№%d", i)
			safeMap.Add(key, i)                            // Добавляем данные в мапу
			value, _ := safeMap.Get(key)                   // Получаем данные из мапы
			fmt.Printf("Key: %s, Value: %v\n", key, value) // Выводим данные в консоль
		}(i) // Передаем текущий индекс в горутину
	}

	wg.Wait() // Ожидаем завершения всех горутин
}

func NewSafeMap() *SafeMap { // Для создания нового экземпляра структуры
	return &SafeMap{
		m: make(map[string]interface{}),
	}
}

func (s *SafeMap) Add(key string, value interface{}) { // Для добавления элемента в мапу
	s.mu.Lock()         // Блокируем мьютекс для безопасного доступа к мапе
	defer s.mu.Unlock() // Разблокируем мьютекс при завершении функции
	s.m[key] = value    // Записываем значение в мапу
}

func (s *SafeMap) Get(key string) (interface{}, bool) { // Для получения значения по ключу
	s.mu.Lock()
	defer s.mu.Unlock()
	value, ok := s.m[key] // Проверяем наличие ключа в карте и получаем значение
	return value, ok      // Возвращаем значение и признак его наличия
}
