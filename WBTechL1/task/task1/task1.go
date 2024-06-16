package main

import "fmt"

type Human struct { // Создаем структуру Human
	Name    string
	Surname string
	Age     int
}

type Action struct { // Создаем структуру Action
	Human
	Type string
}

func main() {
	human := Human{"Cassandra", "Smith", 19} // Создаем экземпляр структуры Human

	action := Action{human, "working"} // Создаем экземпляр структуры Action

	action.Working()                                         // Вызываем метод Working через экземпляр Action
	fmt.Printf("\n%s is %s now\n", action.Name, action.Type) // Получаем данные из полей структуры Human
	// через экземпляр Action
}

func (h Human) Working() { // Метод структуры Human
	fmt.Printf("%s %s works from 6 a.m. to 6 a.m.", h.Name, h.Surname)
}
