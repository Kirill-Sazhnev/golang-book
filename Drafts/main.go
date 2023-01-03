package main

import (
	"image"
	"log"
	"sync"
	"time"
)

type command int

const (
	right   = command(0)
	left    = command(1)
	back    = command(-1)
	stop    = command(2)
	start   = command(3)
	rows    = 5
	columns = 5
)

type Cell struct {
	digit int8
	empty bool
}

// Grid является сеткой Судоку
type Grid [rows][columns]Cell

type MarsGrid struct {
	mu   sync.Mutex
	grid Grid
}

func (g *MarsGrid) Occupy(p image.Point) *Occupier {
	return nil
}

// Occupier представляет занятую клетку сетки
// Может использоваться конкурентно другой сеткой
type Occupier struct {
	// Готово
}

type RoverDriver struct {
	commandc  chan command
	pos       image.Point
	direction image.Point
}

func NewRoverDriver() *RoverDriver {
	r := &RoverDriver{
		commandc:  make(chan command),
		pos:       image.Point{X: 0, Y: 0},
		direction: image.Point{X: 1, Y: 0},
	}
	go r.drive()
	return r
}

// Left поворачивает марсоход налево (90° против часовой стрелки).
func (r *RoverDriver) Left() {
	r.commandc <- left
}

// Right поворачивает марсоход направо (90° по часовой стрелке).
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
}

func (r *RoverDriver) drive() {

	updateInterval := 250 * time.Millisecond
	nextMove := time.After(updateInterval)
active:
	for {
		select {
		case c := <-r.commandc: // Ждет новых команд на командном канале
			switch c {
			case left: // поворот направо
				r.direction = image.Point{
					X: -r.direction.Y,
					Y: r.direction.X,
				}
			case right: // поворот налево
				r.direction = image.Point{
					X: r.direction.Y,
					Y: -r.direction.X,
				}
			case back:
				r.direction = image.Point{
					X: -r.direction.X,
					Y: r.direction.Y,
				}
			case stop:
				break active
			}
			log.Printf("new direction %v", r.direction)
		case <-nextMove:
			r.pos = r.pos.Add(r.direction)
			log.Printf("moved to %v", r.pos)
			nextMove = time.After(updateInterval)

		}
	}
}

func main() {
	r := NewRoverDriver()
	time.Sleep(2 * time.Second)
	r.Left()
	time.Sleep(2 * time.Second)
	r.Right()
	r.Stop()
	time.Sleep(2 * time.Second)
	r.Start()
	time.Sleep(2 * time.Second)
	r.Back()
	time.Sleep(2 * time.Second)
}
