// bla bla bla test
package main

import (
	"fmt"
	"image"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"sync"
	"time"
)

type command int

const (
	right = command(0)
	left  = command(1)
	back  = command(2)
	start = command(3)
	stop  = command(4)
)

var updateInterval = 500 * time.Millisecond

type RoverDriver struct {
	name      string
	commandc  chan command
	direction image.Point
	occupier  *Occupier
}

type MarsGrid struct {
	mu     sync.Mutex
	bounds image.Rectangle
	grid   [][]Occupier
	radio  RadioBuffer
}

type Occupier struct {
	busy bool
	pos  image.Point
	grid *MarsGrid
	life uint16
}

type RadioBuffer struct {
	mu        sync.Mutex
	signal    chan Message
	toEarth   chan Message
	fromEarth chan string
	buffer    []Message
}

type Message struct {
	rover       string
	life        uint16
	coordinates image.Point
}

type EarthComs struct {
	Data []Message
}

func scanPlanet(x, y int) *MarsGrid {
	mars := &MarsGrid{
		mu:     sync.Mutex{},
		bounds: image.Rectangle{Min: image.Point{X: 0, Y: 0}, Max: image.Point{X: x, Y: y}},
		grid:   make([][]Occupier, y),
		radio: RadioBuffer{
			mu:        sync.Mutex{},
			signal:    make(chan Message),
			toEarth:   make(chan Message),
			fromEarth: make(chan string),
			buffer:    []Message{},
		},
	}
	for column := range mars.grid {
		mars.grid[column] = make([]Occupier, x)
		for row := range mars.grid[column] {
			mars.grid[column][row].busy = false
			mars.grid[column][row].pos = image.Point{
				X: row,
				Y: column,
			}
			mars.grid[column][row].life = uint16(rand.Intn(999) + 1)
			mars.grid[column][row].grid = mars
		}
	}
	return mars
}

func (g *MarsGrid) Occupy(p image.Point) *Occupier {
	g.mu.Lock()
	defer g.mu.Unlock()

	if p.In(g.bounds) && !g.grid[p.Y][p.X].busy {
		g.grid[p.Y][p.X].busy = true
		return &g.grid[p.Y][p.X]
	}

	return nil
}

func (g *Occupier) Move(p image.Point) bool {

	if g.grid.Occupy(p) != nil {

		g.grid.mu.Lock()
		defer g.grid.mu.Unlock()

		g.grid.grid[g.pos.Y][g.pos.X].busy = false

		return true
	}
	return false
}

func (r *RoverDriver) Left() {
	r.commandc <- left
}

func (r *RoverDriver) Right() {
	r.commandc <- right
}

func (r *RoverDriver) Back() {
	r.commandc <- back
}

func (r *RoverDriver) Stop() {
	r.commandc <- stop
}

func (r *RoverDriver) Start() {
	go r.drive()
	log.Printf("%v has started the engine at %v with direction %v", r.name, r.occupier.pos, r.direction)
}

func (r *RoverDriver) Random() {
	r.commandc <- command(rand.Intn(3))
}

func NewRoverDriver(mars *MarsGrid, name string) *RoverDriver {
	landPoint := image.Point{
		X: rand.Intn(mars.bounds.Max.X),
		Y: rand.Intn(mars.bounds.Max.Y),
	}

	for {
		if mars.Occupy(landPoint) == nil {
			landPoint = image.Point{
				X: rand.Intn(mars.bounds.Max.X),
				Y: rand.Intn(mars.bounds.Max.Y),
			}
		} else {
			break
		}
	}

	r := &RoverDriver{
		name:      name,
		commandc:  make(chan command),
		direction: image.Point{X: 1, Y: 0},
		occupier:  &mars.grid[landPoint.Y][landPoint.X],
	}

	//show(r.occupier.grid)
	log.Printf("Rover %v has landed at %v", r.name, r.occupier.pos)

	go r.drive()
	return r
}

func (r *RoverDriver) drive() {

	nextMove := time.After(updateInterval)
	newPos := r.occupier.pos.Add(r.direction)

	for {
		select {
		case c := <-r.commandc: // Ждет новых команд на командном канале

			switch c {
			case left: // поворот направо
				log.Printf("%v is turning left\n", r.name)
				r.direction = image.Point{
					X: r.direction.Y,
					Y: -r.direction.X,
				}
			case right: // поворот налево
				log.Printf("%v is turning right\n", r.name)
				r.direction = image.Point{
					X: -r.direction.Y,
					Y: r.direction.X,
				}
			case back:
				log.Printf("%v is turning back\n", r.name)
				r.direction = image.Point{
					X: -r.direction.X,
					Y: -r.direction.Y,
				}
			case stop:
				log.Printf("%v has stoped at %v with direction %v", r.name, r.occupier.pos, r.direction)
				return
			}

			log.Printf("%v is taking new direction %v", r.name, r.direction)
			newPos = r.occupier.pos.Add(r.direction)

		case <-nextMove:
			if r.occupier.Move(newPos) {
				r.occupier = &r.occupier.grid.grid[newPos.Y][newPos.X]
				//show(r.occupier.grid)
				log.Printf("%v moved to %v\n", r.name, newPos)

				if r.occupier.life > 800 {
					log.Printf("Sending signal of founded life %v\n", r.occupier.life)
					r.radioSend(r.occupier)
					r.occupier.grid.grid[newPos.Y][newPos.X].life = 0
				}

			} else {
				go r.Random()
			}

			newPos = r.occupier.pos.Add(r.direction)
			nextMove = time.After(updateInterval)
		}
		//updateInterval += (50 * time.Millisecond)
	}
}

func (r *RoverDriver) radioSend(cell *Occupier) {
	signal := Message{
		rover:       r.name,
		life:        cell.life,
		coordinates: cell.pos,
	}
	cell.grid.radio.signal <- signal
}

func radio(radio *RadioBuffer) {

	for {
		select {
		case message := <-radio.signal:
			radio.buffer = append(radio.buffer, message)
		case coms := <-radio.fromEarth:
			if coms == "open" {
				radio.mu.Lock()
				log.Println("Earth signal received, establishing transmition")
				for _, msg := range radio.buffer {
					radio.toEarth <- msg
				}
				radio.mu.Unlock()
				radio.buffer = []Message{}
			}
		}
	}

}

func (earth *EarthComs) ActivateCom(planet *MarsGrid, interval int64, duration int64) {
	for {
		time.Sleep(time.Duration(interval) * 1000000000)
		planet.radio.fromEarth <- "open"
		transmition := time.After(time.Duration(duration) * 1000000000)
	transmission:
		for {
			select {
			case msg := <-planet.radio.toEarth:
				earth.Data = append(earth.Data, msg)

			case signal := <-planet.radio.signal:
				earth.Data = append(earth.Data, signal)

			case <-transmition:
				break transmission
			}
		}
	}
}

var frame sync.Mutex

func show(g *MarsGrid) {
	frame.Lock()
	defer frame.Unlock()

	cls()
	for _, row := range g.grid {
		for _, cell := range row {
			if cell.busy {
				fmt.Print("[o-o]")
			} else {
				fmt.Print("[   ]")
			}
		}
		fmt.Println()
	}
}

func display(g *MarsGrid) {
	for {
		show(g)
		time.Sleep(updateInterval - 10)
	}
}

func cls() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {
	mars := scanPlanet(7, 9)
	earth := &EarthComs{
		Data: []Message{},
	}

	go radio(&mars.radio)
	go earth.ActivateCom(mars, 15, 3)

	go display(mars)

	r := NewRoverDriver(mars, "Curiosity")

	time.Sleep(6 * time.Second)
	r.Stop()

	p := NewRoverDriver(mars, "Perseverance")

	time.Sleep(4 * time.Second)
	r.Start()
	time.Sleep(8 * time.Second)

	NewRoverDriver(mars, "Love")

	r.Right()
	p.Stop()

	time.Sleep(6 * time.Second)

	NewRoverDriver(mars, "Moon")

	p.Start()

	time.Sleep(15 * time.Second)

	//showing results:
	show(mars)

	fmt.Println()
	fmt.Println("Messages left on Mars Buffer:")
	fmt.Println()

	for _, life := range mars.radio.buffer {
		fmt.Printf("%+v\n", life)
	}

	fmt.Println()
	fmt.Println("Received by Earth Com Center:")
	fmt.Println()
	for _, life := range earth.Data {
		fmt.Printf("%+v\n", life)
	}
	time.Sleep(15 * time.Second)
}
