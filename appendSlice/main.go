package main

import (
	"fmt"
)

func main() {
	planets := []string{
		"Меркурий", "Венера", "Земля", "Марс",
		"Юпитер", "Сатурн", "Уран", "Нептун", "Mercury", "Moon",
	}
	solarSys := make([]string, 3, 3)

	for i, planet := range planets {
		if i >= cap(solarSys) {
			fmt.Println(solarSys, len(solarSys), cap(solarSys))
		}

		if i < len(solarSys) {
			solarSys[i] = planet

		} else {
			solarSys = append(solarSys, planet)
		}
	}

	s := []string{}
	lastCap := cap(s)

	for i := 0; i < 10000; i++ {
		s = append(s, "An element")
		if cap(s) != lastCap {
			fmt.Println(cap(s))
			lastCap = cap(s)
		}
	}
}

/*
Напишите программу, что использует цикл для продолжающегося добавления
элементов в срез. Каждый раз при изменении вместимости среза выводится
новое значение. Всегда ли append удваивает вместимость при завершении места
в базовом массиве?
*/
