package main

import "fmt"

type PetPlentyAndOthers struct {
	Cat   []string // Создаем подмножества для элементов множества
	Dog   []string
	Other []string
}

func main() {
	PPO := NewSet() // Инициализируем новый экземпляр структуры
	var set int

	elements := []string{"cat", "cat", "dog", "cat", "tree"}

	for _, el := range elements {
		PPO.Add(el) // Добавляем каждый элемент в свое множество
	}

	fmt.Println("1. Cat")
	fmt.Println("2. Dog")
	fmt.Println("3. Other")
	fmt.Print("Select a set: ")
	fmt.Scan(&set)

	fmt.Println("Elements:", PPO.View(set)) // Выводим элементы подмножества в консоль
	fmt.Println("Size:", PPO.Size(set))     // Выводим размер подмножества

	PPO.Clear(set) // Создаем пустое подмножество
	fmt.Println("Size after clearing:", PPO.Size(set))
}

func NewSet() *PetPlentyAndOthers {
	return &PetPlentyAndOthers{
		Cat:   []string{},
		Dog:   []string{},
		Other: []string{},
	}
}

func (p *PetPlentyAndOthers) Add(el string) {
	switch el {
	case "cat":
		p.Cat = append(p.Cat, el)
	case "dog":
		p.Dog = append(p.Dog, el)
	default:
		p.Other = append(p.Other, el)
	}
}

func (p *PetPlentyAndOthers) Size(n int) int {
	switch n {
	case 1:
		return len(p.Cat)
	case 2:
		return len(p.Dog)
	default:
		return len(p.Other)
	}
}

func (p *PetPlentyAndOthers) Clear(n int) {
	switch n {
	case 1:
		p.Cat = []string{}
	case 2:
		p.Dog = []string{}
	default:
		p.Other = []string{}
	}
}

func (p *PetPlentyAndOthers) View(n int) []string {
	switch n {
	case 1:
		return p.Cat
	case 2:
		return p.Dog
	default:
		return p.Other
	}
}
