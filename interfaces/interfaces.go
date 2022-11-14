package main

import "fmt"

type Reader interface {
	Read(personName string) string
}

type Book struct {
	Text string
}

func (b Book) Read(personName string) string {
	return personName + " read " + b.Text
}

type Journal struct {
	Text string
}

func (j Journal) Read(personName string) string {
	return personName + " read " + j.Text
}

func main() {
	me := "Jafree"

	var somethingToRead Reader

	shopToVisit := ShopA{}

	somethingToRead = shopToVisit.Sell()

	weirdLine := somethingToRead.Read(me)

	fmt.Println(weirdLine)
}

type ShopA struct {
	books []Book
}

func (s ShopA) Sell() Book {
	return s.books[0]
}

type ShopB struct {
	journals []Journal
}

func (s ShopB) Sell() Journal {
	return s.journals[0]
}
