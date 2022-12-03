package main

import (
	"fmt"
	"math/rand"
	"time"
)

type farmer interface {
	move() string
	feed() string
}

type dog struct {
	name string
}

func (d dog) String() string {
	return fmt.Sprintf("%v", d.name)
}

func (d dog) move() string {
	return fmt.Sprintf("%v runs", d)
}

func (d dog) feed() string {
	switch rand.Intn(3) {
	case 0:
		return fmt.Sprintf("%v eats bone", d)
	case 1:
		return fmt.Sprintf("%v eats dry food", d)
	case 2:
		return fmt.Sprintf("%v eats meat", d)
	default:
		return " "
	}
}

type horse struct {
	name string
}

func (h horse) String() string {
	return fmt.Sprintf("%v", h.name)
}

func (h horse) move() string {
	return fmt.Sprintf("%v gallops", h)
}

func (h horse) feed() string {
	switch rand.Intn(3) {
	case 0:
		return fmt.Sprintf("%v eats grass", h)
	case 1:
		return fmt.Sprintf("%v eats apple", h)
	case 2:
		return fmt.Sprintf("%v eats treat", h)
	default:
		return " "
	}
}

type bird struct {
	name string
}

func (b bird) String() string {
	return fmt.Sprintf("%v", b.name)
}

func (b bird) move() string {
	return fmt.Sprintf("%v flyes", b)
}

func (b bird) feed() string {
	switch rand.Intn(2) {
	case 0:
		return fmt.Sprintf("%v eats insekt", b)
	case 1:
		return fmt.Sprintf("%v eats seed", b)
	default:
		return " "
	}
}

type fish struct {
	name string
}

func (f fish) String() string {
	return fmt.Sprintf("%v", f.name)
}

func (f fish) move() string {
	return fmt.Sprintf("%v swims", f)
}

func (f fish) feed() string {
	switch rand.Intn(3) {
	case 0:
		return fmt.Sprintf("%v eats worm", f)
	case 1:
		return fmt.Sprintf("%v eats planktone", f)
	case 2:
		return fmt.Sprintf("%v eats bread crump", f)
	default:
		return " "
	}
}
func main() {
	felix := dog{"Felix"}
	swift := horse{"Swift"}
	hawk := bird{"Hawk"}
	neo := fish{"Neo"}

	var animal farmer

	for i := 0; i < 72; i++ {
		j := i % 24
		fmt.Printf("%02d:00 ", j)

		if i/8%3 == 0 {
			fmt.Println("All animals asleep")
		} else {

			switch rand.Intn(4) {
			case 0:
				animal = felix
			case 1:
				animal = swift
			case 2:
				animal = hawk
			case 3:
				animal = neo
			}

			switch rand.Intn(2) {
			case 0:
				fmt.Println(animal.move())
			case 1:
				fmt.Println(animal.feed())
			}
		}
		time.Sleep(time.Second / 2)
	}
}
