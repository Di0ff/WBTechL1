package main

import "fmt"

type Work interface {
	Cook() string
}

type Relax interface {
	Draw() string
}

type Profession struct {
	Chef string
}

type Hobby struct {
	Artist string
}

type Adapter struct { // Создаем адаптер
	profession Profession
	hobby      Hobby
}

func main() {
	profession := Profession{"Pizza Chef"}
	hobby := Hobby{"Painter"}

	fmt.Println(profession.Cook())
	fmt.Println(hobby.Draw())

	adapter := Adapter{profession, hobby}

	fmt.Println(adapter.CookAndDraw())
	fmt.Println(adapter.DrawAndCook())
}

func (p Profession) Cook() string {
	return "Cook a pizza"
}

func (h Hobby) Draw() string {
	return "To paint a picture"
}

func (a Adapter) CookAndDraw() string { // Реализуем методы интерфейсов Work и Relax для адаптера
	return a.profession.Cook() + " and " + a.hobby.Draw()
}

func (a Adapter) DrawAndCook() string {
	return a.hobby.Draw() + " and " + a.profession.Cook()
}
