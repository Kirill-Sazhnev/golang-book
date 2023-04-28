package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type Book struct {
	title    string
	price    float64
	quantity int
}

func main() {
	fileHand, _ := os.Open("products.txt")
	defer fileHand.Close()

	newReader := csv.NewReader(fileHand)
	newReader.LazyQuotes = true
	newReader.Comma = ';'

	Books := []Book{}

	for {
		line, err := newReader.Read()
		if err != nil {
			fmt.Printf("%v\n", err)
			break
		}

		v1 := line[0]
		v2, _ := strconv.ParseFloat(line[1], 64)
		v3, _ := strconv.Atoi(line[2])
		Books = append(Books, Book{
			title:    v1,
			price:    v2,
			quantity: v3,
		})
	}
	fmt.Printf("%#v", Books)
}

/* CSV  File:
"The ABC of Go";25.5;1500
"Functional Programming with Go";56;280
"Go for It";45.9;356
"The Go Way";55;500
*/
