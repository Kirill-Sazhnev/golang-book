package main

import (
	"fmt"
)

type item struct {
	name string
}

type character struct {
	name     string
	leftHand *item
}

func (c *character) pickup(i *item) {
	if c == nil || i == nil {
		return
	}
	fmt.Printf("%v поднимает %v\n", c.name, i.name)
	c.leftHand = i
}

func (c *character) give(to *character) {
	if c == nil || to == nil {
		return
	}
	if c.leftHand == nil {
		fmt.Printf("%v ничего не может дать\n", c.name)
		return
	}
	if to.leftHand != nil {
		fmt.Printf("%v с занятыми руками\n", to.name)
		return
	}
	to.leftHand = c.leftHand
	c.leftHand = nil
	fmt.Printf("%v дает %v %v\n", c.name, to.name, to.leftHand.name)
}

func (c character) String() string {
	if c.leftHand == nil {
		return fmt.Sprintf("%v ничего не несет", c.name)
	}
	return fmt.Sprintf("%v несет %v", c.name, c.leftHand.name)
}

func main() {
	arthur := &character{name: "Артур"}

	shrubbery := &item{name: "мечь"}
	arthur.pickup(shrubbery) // Выводит: Артур поднимает кустарник

	knight := &character{name: "Рыцарь/ю"}
	arthur.give(knight) // Выводит: Артур дает Рыцарь/ю кустарник

	fmt.Println(arthur) // Выводит: Артур ничего не несет
	fmt.Println(knight) // Выводит: Рыцарь/ю несет кустарник
}
