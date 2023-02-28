package main

import (
	"container/list"
	"fmt"

	"github.com/golang-collections/collections/queue"
	"github.com/golang-collections/collections/stack"
)

// setting set as map
type Set map[interface{}]bool

// Add value in set
func (s *Set) Add(key interface{}) {
	(*s)[key] = true
}

// Remove value from set
func (s *Set) Remove(key interface{}) {
	delete((*s), key)
}

// Find is key available in set or not
func (s *Set) Find(key interface{}) bool {
	return (*s)[key]
}

// setting counter as map
type Counter map[interface{}]int

func (s *Counter) Add(key interface{}) {
	(*s)[key] += 1
}
func (s *Counter) Find(key interface{}) bool {
	_, ok := (*s)[key]
	return ok
}
func (s *Counter) Get(key interface{}) (int, bool) {
	val, ok := (*s)[key]
	return val, ok
}

func main() {
	s := stack.Stack{}
	s.Push(1)
	s.Push(2)
	s.Push(3)
	for s.Len() != 0 {
		val := s.Pop()
		fmt.Print(val, " ")
	}

	fmt.Println()

	q := queue.Queue{}
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	for q.Len() != 0 {
		val := q.Dequeue().(int)
		fmt.Print(val, " ")
	}
	fmt.Println()

	hash := map[string][]int{}

	hash["val"] = []int{3, 5, 7}
	hash["bal"] = []int{4, 6, 8}
	delete(hash, "val")
	hash["bal"] = append(hash["bal"], 10)

	fmt.Println(hash)

	fmt.Println("SET")

	mp := make(Set)
	mp.Add(1)
	mp.Add(2)
	fmt.Println(mp.Find(1))
	fmt.Println(mp.Find(3))
	mp.Remove(1)
	fmt.Println(mp.Find(1))

	fmt.Println("COUNTER")

	ct := make(Counter)
	ct.Add("a")
	ct.Add("b")
	ct.Add("a")
	fmt.Println(ct.Find("a"))
	fmt.Println(ct.Find("b"))
	fmt.Println(ct.Find("c"))
	fmt.Println(ct.Get("a"))
	fmt.Println(ct.Get("b"))
	fmt.Println(ct.Get("c"))

	// Create a new linked doubly list
	fmt.Println("LINKED LIST")
	myList := list.New()

	// Add elements to the list
	myList.PushBack(1)
	myList.PushBack(2)
	myList.PushBack(3)

	// Traverse the list and print its contents
	for e := myList.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

	// Remove an element from the list
	if element := myList.Front(); element != nil {
		myList.Remove(element)
	}

	// Traverse the list again and print its contents
	for e := myList.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
