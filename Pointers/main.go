package main

import (
	"fmt"
)

type turtle struct {
	x, y int
}

func (t *turtle) moveUp(s int) {
	t.y -= s
}

func (t *turtle) moveDown(s int) {
	t.y += s
}

func (t *turtle) moveLeft(s int) {
	t.x -= s
}

func (t *turtle) moveRight(s int) {
	t.x += s
}

func main() {
	tortilla := turtle{0, 0}
	tortilla.moveDown(3)
	tortilla.moveLeft(4)
	tortilla.moveRight(9)
	fmt.Println(tortilla)
}
