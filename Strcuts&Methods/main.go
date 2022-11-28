package main

import (
	"fmt"
	"math"
)

type coordinate struct {
	d, m, s float64
	h       rune
}

type location struct {
	lat, long float64
}

func (c coordinate) decimal() float64 {
	sign := 1.0
	switch c.h {
	case 'S', 'W', 's', 'w':
		sign = -1
	}
	return sign * (c.d + c.m/60 + c.s/3600)
}

func newLocation(lat, long coordinate) location {
	return location{lat.decimal(), long.decimal()}
}

type world struct {
	radius float64
}

func rad(deg float64) float64 {
	return deg * math.Pi / 180
}

func (w world) distance(p1, p2 location) float64 {
	s1, c1 := math.Sincos(rad(p1.lat))
	s2, c2 := math.Sincos(rad(p2.lat))
	clong := math.Cos(rad(p1.long - p2.long))
	return w.radius * math.Acos(s1*s2+c1*c2*clong) // Использует поле радиуса world
}

func main() {

	spirit := newLocation(coordinate{14, 36, 6.2, 'S'}, coordinate{175, 28, 21.5, 'E'})
	opportunity := newLocation(coordinate{1, 56, 46.4, 'S'}, coordinate{354, 28, 24.2, 'E'})
	curiosity := newLocation(coordinate{4, 35, 22.2, 'S'}, coordinate{136, 27, 30.1, 'E'})
	inSight := newLocation(coordinate{4, 30, 0.0, 'N'}, coordinate{135, 54, 0.0, 'E'})
	london := newLocation(coordinate{51, 30, 0.0, 'N'}, coordinate{0, 8, 0, 'W'})
	paris := newLocation(coordinate{48, 51, 0.0, 'N'}, coordinate{2, 21, 0, 'E'})

	var mars = world{radius: 3389.5}
	var earth = world{6371.0}

	howers := []location{spirit, opportunity, curiosity, inSight}
	howerNames := []string{"Spirit", "Opportunity", "Curiosity", "InSight"}

	fmt.Printf("Spirit: %.2f\nOpportunity: %.2f\nCuriosity: %.2f\nInSight: %.2f\n", spirit, opportunity, curiosity, inSight)

	for i, hower1 := range howers {

		for j, hower2 := range howers {
			if hower1 != hower2 && j > i {
				fmt.Printf("Distance between %v and %v is %.2f km\n", howerNames[i], howerNames[j], mars.distance(hower1, hower2))
			}
		}
	}

	fmt.Printf("Distance between London and Paris is %.2f km", earth.distance(london, paris))
}
