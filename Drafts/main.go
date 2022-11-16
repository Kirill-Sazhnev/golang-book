package main

import (
	"fmt"
)

type Planets []string

func (p Planets) terraform() {

	for i, planet := range p {
		p[i] = "New " + planet
	}
}

func main() {
	planets := Planets{"Mars", "Uranus", "Neptun"}

	planets.terraform()

	fmt.Println(planets)
}

/*
Напишите программу для преобразования слайса строки через добавление слова
"Новый " перед названием планеты. Используйте программу для изменения
 названий планет Марс, Уран и Нептун.

В первой итерации может использоваться функция terraform, но в конечной
реализации должен быть введен тип Planets с методом terraform, похожим на
sort.StringSlice.
*/
