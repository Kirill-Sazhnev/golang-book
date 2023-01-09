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
	right   = command(0)
	left    = command(1)
	back    = command(2)
	start   = command(3)
	stop    = command(4)
	rows    = 8
	columns = 8
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
	grid   [rows][columns]Occupier
}

type Occupier struct {
	busy bool
	pos  image.Point
	grid *MarsGrid
}

func (g *MarsGrid) Occupy(p image.Point) *Occupier {
	g.mu.Lock()
	defer g.mu.Unlock()

	if p.In(g.bounds) && !g.grid[p.Y][p.X].busy {
		g.grid[p.Y][p.X].pos = p
		g.grid[p.Y][p.X].busy = true
		g.grid[p.Y][p.X].grid = g
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
		X: rand.Intn(columns),
		Y: rand.Intn(rows),
	}

	for {
		if mars.Occupy(landPoint) == nil {
			landPoint = image.Point{
				X: rand.Intn(columns),
				Y: rand.Intn(rows),
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

	log.Printf("Rover %v has landed at %v", r.name, r.occupier.pos)
	show(r.occupier.grid)

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
				fmt.Printf("%v is turning left\n", r.name)
				r.direction = image.Point{
					X: r.direction.Y,
					Y: -r.direction.X,
				}
			case right: // поворот налево
				fmt.Printf("%v is turning right\n", r.name)
				r.direction = image.Point{
					X: -r.direction.Y,
					Y: r.direction.X,
				}
			case back:
				fmt.Printf("%v is turning back\n", r.name)
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
				show(r.occupier.grid)
				fmt.Printf("%v moved to %v\n", r.name, newPos)
			} else {
				go r.Random()
			}

			newPos = r.occupier.pos.Add(r.direction)
			nextMove = time.After(updateInterval)
		}
		//updateInterval += (50 * time.Millisecond)
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
		time.Sleep(updateInterval)
		cls()
		fmt.Println()
	}
}

func cls() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {
	mars := &MarsGrid{
		mu: sync.Mutex{},
		bounds: image.Rectangle{
			Min: image.Point{
				X: 0,
				Y: 0,
			},
			Max: image.Point{
				X: columns,
				Y: rows,
			},
		},
		grid: [columns][rows]Occupier{},
	}

	r := NewRoverDriver(mars, "Curiosity")

	//go display(mars)

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

	p.Start()

	time.Sleep(15 * time.Second)
	show(mars)
}
