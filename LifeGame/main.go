package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

const (
	width  = 65
	height = 15
)

type Universe [][]bool

func newUnivers() Universe {
	uni := make(Universe, height)
	for i := 0; i < height; i++ {
		uni[i] = make([]bool, width)
	}
	return uni
}

func (u Universe) Show() {
	for _, y := range u {
		for _, x := range y {
			if x {
				fmt.Printf("%v", "*")
			} else {
				fmt.Printf("%v", "_")
			}
		}
		fmt.Println()
	}
}

func (u Universe) Seed() {
	for y, line := range u {
		for x := range line {
			if rand.Intn(4) == 1 {
				u[y][x] = true
			}
		}
	}
}

func (u Universe) Alive(x, y int) bool {
	cell := u[(y+height)%height][(x+width)%width]
	return cell
}

func (u Universe) Neighbors(x, y int) int {
	counter := 0

	if u.Alive(x-1, y) {
		counter++
	}
	if u.Alive(x+1, y) {
		counter++
	}
	if u.Alive(x-1, y-1) {
		counter++
	}
	if u.Alive(x, y-1) {
		counter++
	}
	if u.Alive(x+1, y-1) {
		counter++
	}
	if u.Alive(x-1, y+1) {
		counter++
	}
	if u.Alive(x, y+1) {
		counter++
	}
	if u.Alive(x+1, y+1) {
		counter++
	}

	return counter
}

func (u Universe) Next(x, y int) bool {
	n := u.Neighbors(x, y)
	a := u.Alive(x, y)

	if a && n < 2 {
		return false
	} else if a && (n == 2 || n == 3) {
		return true
	} else if a && n > 3 {
		return false
	} else if !a && n == 3 {
		return true
	} else {
		return false
	}
}

func Step(a, b Universe) {
	for y, line := range a {
		for x := range line {
			b[y][x] = a.Next(x, y)
		}
	}
	b.Show()
}

func cls() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {
	uni := newUnivers()
	clone := newUnivers()

	cls()
	uni.Seed()
	uni.Show()

	time.Sleep(time.Second)
	cls()

	for i := 0; i < 90; i++ {

		Step(uni, clone)
		uni, clone = clone, uni
		time.Sleep(time.Second / 3)

		cls()
	}

}
