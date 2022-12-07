package main

import (
	"fmt"
)

type item struct {
	name  string
	valid bool
}

type character struct {
	name     string
	leftHand *item
}

func newItem(i string) item {
	return item{i, true}
}

func (c *character) pickup(i *item) {
	if !i.valid {
		fmt.Println("There is nothing to pick up")
		return
	}
	c.leftHand = i
	fmt.Printf("%v picks up a %v\n", c.name, c.leftHand)
}

func (c *character) give(to *character) {
	to.leftHand = c.leftHand
	c.leftHand = nil
	fmt.Printf("%v gives his %v to %v\n", c.name, to.leftHand, to.name)
}

func (i item) String() string {
	if !i.valid {
		return "Nil value detected, cannot print"
	}

	return fmt.Sprintf("%s", i.name)
}

func main() {
	//sword := newItem("sword")
	sword := item{}

	artur := &character{name: "Artur"}
	knight := &character{name: "Knight"}

	artur.pickup(&sword)
	artur.give(knight)

}

/*Рыцарь встал на пути Артура. Герой безоружен, он представлен значением nil
для leftHand *item. Имплементируйте структуру character с методами вроде
pickup(i *item) и give(to *character). Потом используйте выученное в данном
уроке для написания скрипта, в котором Артур достает объект и передает его
рыцарю. Каждое действие должно отображаться с соответствующим описанием.
*/
