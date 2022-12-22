package main

import (
	"image"
	"log"
	"time"
)

type command int

const (
	right = command(0)
	left  = command(1)
	back  = command(-1)
	stop  = command(2)
	start = command(3)
)

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

/*

func worker() {
	pos := image.Point{X: 10, Y: 10}     // Текущая позиция (изначально [10, 10])
	direction := image.Point{X: 1, Y: 0} // Текущее направление (изначально [1, 0])
	next := time.After(time.Second)      // Создаем начальный канал таймера
	for {
		select {
		case <-next: // Ожидает истечение срока таймера
			pos = pos.Add(direction)
			fmt.Println("текущая позиция ", pos) // Выводит текущую позицию
			next = time.After(time.Second)       // Создает другой канал таймера для другого события
			fmt.Println("!")
		}
	}
}

func main() {
	worker()
}
*/
