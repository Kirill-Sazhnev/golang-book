package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	return ioutil.WriteFile(p.Title+".txt", p.Body, 0664)
}

func load(title string) (*Page, error) {
	page := new(Page)
	var err error
	page.Title = title
	page.Body, err = ioutil.ReadFile(title + ".txt")

	return page, err
}

func main() {
	page := &Page{
		Title: "PD is Hard",
		Body:  []byte("Learning PD is a hard work requiring motivation, discipline and dedication. \nIt can be frustrating and confusing one day and exiting and fascinating the other. \nAs long as you move forward even with a small steps you are closer to the day of becoming a PD"),
	}

	err := page.save()

	if err != nil {
		log.Fatal(err)
	}

	text, errText := load(page.Title)

	if errText != nil {
		log.Fatal(errText)
	}
	fmt.Printf("Page: \"%v\"\n%v", text.Title, string(text.Body))
}
