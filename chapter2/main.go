package main

import ("fmt"; "math")

type Circle struct {
  x, y, r float64
}

func (c *Circle) area() float64 {
  return math.Pi * c.r*c.r
}

type Rectangle struct {
  x1, y1, x2, y2 float64
}

func distance(x1, y1, x2, y2 float64) float64 {
  a := x2 - x1
  b := y2 - y1
  return math.Sqrt(a*a + b*b)
}

func (r *Rectangle) area() float64 {
  l := distance(r.x1, r.y1, r.x1, r.y2)
  w := distance(r.x1, r.y1, r.x2, r.y1)
  return l * w
}

type Shape interface {
  area() float64
}

type MultiShape struct {
  shapes []Shape
}

func (m *MultiShape) area() float64 {
  var area float64
  for _, s := range m.shapes {
    area += s.area()
  }
  return area
}

func main() {
  multiShape := MultiShape{ //array of structs
  shapes: []Shape{ //array of interfaces
    &Circle{0,0,5}, // why do we assign values to the interface? Interface is list of methods not values
    &Rectangle{0,0,10,10},
    },
  }
  fmt.Println(multiShape.area())
}


//  multiShape := MultiShape{ //array of structs
//    shapes: []Shape{ //array of interfaces
//      Circle{0,0,5}, // why do we assign values to the interface? Interface is list of methods not values
//      Rectangle{0,0,10,10},
//    },
//  }

//  r := Rectangle{0, 0, 10, 10}
//  c := Circle{x: 0, y: 0, r: 5}

//  fmt.Println(r.area())
//  fmt.Println(c.area())
//  fmt.Println(multiShape.shapes.area())


/*
func fib(num int) (fnum int){
  switch num {
  case 0: fnum = 0
  case 1: fnum = 1
  default: fnum = fib(num-1) + fib(num-2)
  }
  return
}

func main() {
  fmt.Println("Enter a number: ")
  var input int
  fmt.Scanf("%d", &input)
  fmt.Println("Your input is", input)
  fmt.Println(fib(input))
  }



/*
func isEven(x int) (int, bool){
  if x % 2 == 0 {
    return x/2, true
  } else {
    return x/2, false
  }
}

func main() {
  fmt.Println("Enter a number: ")
  var input int
  fmt.Scanf("%d", &input)
  fmt.Println("Your input is", input)
  fmt.Println(isEven(input))
}
*/

/*
  x := []int{
    48,96,86,68,
    57,82,63,70,
    37,34,83,27,
    19,97, 9,17,
  }

  y := x[0]

  for i := 1 ; i < len(x) ; i++ {
    if y > x[i] {
      y = x[i]
    }
  }
  fmt.Println("The smallest number is",y)
*/
