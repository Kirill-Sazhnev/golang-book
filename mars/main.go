package main

import (
  "fmt"
  "math/rand"
)

func main() {
  var station string
  var tripType string
  const distance  = 62100
  
  fmt.Printf("%-15v %-3v %-10v %-5v\n", "Spaceline", "Days", "Trip type", "Price")
  fmt.Println("=====================================")

for i := 0 ; i < 10 ; i++ {
    var typeId = rand.Intn(2)+1
    var speed = rand.Intn(15)+16
    var price = speed+20
    var time = distance/speed/24

    switch stationId := rand.Intn(4); stationId {
    case 0: station = "Space"
    case 1: station = "Adventures"
    case 2: station = "SpaceX"
    case 3: station = "Virgin Galactic"
    }

    switch typeId {
    case 1: tripType = "One-way"
    case 2: tripType = "Round-trip"
    }

    fmt.Printf("%-15v %4v %-10v $ %3v\n", station, time, tripType, price*typeId)
  }
}
