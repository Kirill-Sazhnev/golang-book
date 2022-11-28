package main

import (
	"fmt"
	"math"
)

type rover struct {
	name string
	gps
}

type gps struct {
	current     location
	destination location
	planet      world
}

type location struct {
	name      string
	lat, long float64
}

type world struct {
	radius float64
}

func (l location) description() string {
	return fmt.Sprintf("%v lat: %.2f, long: %.2f", l.name, l.lat, l.long)
}

func rad(deg float64) float64 {
	return deg * math.Pi / 180
}

func (g gps) distance(p1, p2 location) float64 {
	return g.planet.distance(p1, p2)
}

func (w world) distance(p1, p2 location) float64 {
	s1, c1 := math.Sincos(rad(p1.lat))
	s2, c2 := math.Sincos(rad(p2.lat))
	clong := math.Cos(rad(p1.long - p2.long))
	return w.radius * math.Acos(s1*s2+c1*c2*clong) // Использует поле радиуса world
}

func (g gps) message() string {
	return fmt.Sprintf("%.2f km left to destination: %v", g.distance(g.current, g.destination), g.destination.description())
}

func main() {
	mars := gps{
		current:     location{"Bradbury Landing", -4.5895, 137.4417},
		destination: location{"Elysium Planitia", 4.5, 135.9},
		planet:      world{radius: 3389.5},
	}

	curiosity := rover{"Curiosity", mars}
	fmt.Println(curiosity.message())
}
