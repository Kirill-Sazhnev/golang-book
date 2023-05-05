package main

import (
	"fmt"
	"math"
)

type polar struct {
	radius, angle float64
}

type cartesian struct {
	x, y float64
}

func main() {
	channel1 := make(chan polar)
	channel2 := createSolver(channel1)
	interact(channel1, channel2)

}

func createSolver(questions chan polar) chan cartesian {
	cartesianCh := make(chan cartesian)
	go func() {
		for pol := range questions {
			fmt.Printf("Entered Points are: %v\n", pol)
			a := pol.angle * math.Pi / 180.0
			cart := cartesian{
				x: pol.radius * math.Cos(a),
				y: pol.radius * math.Sin(a),
			}
			cartesianCh <- cart
		}
	}()
	return cartesianCh
}

func interact(questions chan polar, answers chan cartesian) {
	pol := new(polar)

	for {
		fmt.Print("Please enter a Polar Points (radius and angle): ")
		n, _ := fmt.Scanf("%f %f", &pol.radius, &pol.angle)
		if n != 2 {
			panic("Number of values is incorrect")
		}
		questions <- *pol

		cart := <-answers
		fmt.Printf("Conversion result is: x %.2f y %.2f\n", cart.x, cart.y)
		fmt.Scanln()
	}
}
