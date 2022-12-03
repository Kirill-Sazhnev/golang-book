package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Пчела
type honeyBee struct {
	name string
}

func (hb honeyBee) String() string {
	return hb.name
}

func (hb honeyBee) move() string {
	switch rand.Intn(2) {
	case 0:
		return "жужжит"
	default:
		return "летает и веселится"
	}
}

func (hb honeyBee) eat() string {
	switch rand.Intn(2) {
	case 0:
		return "пыльцу"
	default:
		return "нектар"
	}
}

// Суслик
type gopher struct {
	name string
}

func (g gopher) String() string {
	return g.name
}

func (g gopher) move() string {
	switch rand.Intn(2) {
	case 0:
		return "гулят и изучает территорию"
	default:
		return "прячется в норку"
	}
}

func (g gopher) eat() string {
	switch rand.Intn(5) {
	case 0:
		return "морковку"
	case 1:
		return "салат-латук"
	case 2:
		return "редиску"
	case 3:
		return "кукурузу"
	default:
		return "корнеплоды"
	}
}

type animal interface {
	move() string
	eat() string
}

func step(a animal) {
	switch rand.Intn(2) {
	case 0:
		fmt.Printf("%v %v.\n", a, a.move())
	default:
		fmt.Printf("%v кушает %v.\n", a, a.eat())
	}
}

const sunrise, sunset = 8, 18

func main() {
	rand.Seed(time.Now().UnixNano())

	animals := []animal{
		honeyBee{name: "Шмель Базз"},
		gopher{name: "Суслик Го"},
	}

	var sol, hour int

	for {
		fmt.Printf("%2d:00 ", hour)
		if hour < sunrise || hour >= sunset {
			fmt.Println("Животные спят.")
		} else {
			i := rand.Intn(len(animals))
			step(animals[i])
		}

		time.Sleep(500 * time.Millisecond)

		hour++
		if hour >= 24 {
			hour = 0
			sol++
			if sol >= 3 {
				break
			}
		}
	}
}
