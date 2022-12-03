package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type coordinate struct {
	d, m, s float64
	h       rune
}

type location struct {
	Decimal    float64 `json:"decimal"`
	Coordinate string  `json:"dms"`
	Degree     float64 `json:"degrees"`
	Minutes    float64 `json:"minutes"`
	Seconds    float64 `json:"second"`
	Hemisphere rune    `json:"hemisphere"`
}

/*
	type Marshaler interface {
		MarshalJSON() ([]byte, error)
	}

	func (l location) String() string {
		return fmt.Sprintf("%v, %v", l.lat, l.long)
	}


*/

func (c coordinate) String() string {
	return fmt.Sprintf("%v°%v'%.1f\"%c", c.d, c.m, c.s, c.h)
}

func (c coordinate) decimal() float64 {
	sign := 1.0
	switch c.h {
	case 'S', 'W', 's', 'w':
		sign = -1
	}
	return sign * (c.d + c.m/60 + c.s/3600)
}

func (c coordinate) MarshalJSON() ([]byte, error) {

	return json.Marshal(location{
		Decimal:    c.decimal(),
		Coordinate: c.String(),
		Degree:     c.d,
		Minutes:    c.m,
		Seconds:    c.s,
		Hemisphere: c.h,
	})
}

func main() {

	elysium := coordinate{135, 54, 0.0, 'E'}
	json, err := json.MarshalIndent(elysium, "", "  ")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(string(json))
}

/*
{
    "decimal": 135.9,
    "dms": "135°54'0.0\" E",
    "degrees": 135,
    "minutes": 54,
    "seconds": 0,
    "hemisphere": "E"
}
*/
