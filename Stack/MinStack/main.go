package main

import (
	"fmt"
)

type Stack struct {
	s []int
}

func (s *Stack) IsEmpty() bool {
	length := len(s.s)
	return length == 0
}

func (s *Stack) Length() int {
	length := len(s.s)
	return length
}

func (s *Stack) Print() {
	length := len(s.s)
	for i := 0; i < length; i++ {
		fmt.Print(s.s[i], " ")
	}
	fmt.Println()
}

func (s *Stack) Push(value int) {
	s.s = append(s.s, value)
}

func (s *Stack) Pop() int {
	if s.IsEmpty() == true {
		fmt.Print("Stack in empty.")
		return 0
	}
	length := len(s.s)
	res := s.s[length-1]
	s.s = s.s[:length-1]
	return res
}

func (s *Stack) Top() int {
	if s.IsEmpty() == true {
		fmt.Print("Stack in empty.")
		return 0
	}
	length := len(s.s)
	res := s.s[length-1]
	return res
}

type MinStack struct {
	maxSize int
	stk     Stack
	// Initialize your data structures here
}

// removes and returns value from stack
func (m *MinStack) Pop() int {
	// write your code here
	m.maxSize--
	return m.stk.Pop()
}

// Pushes value into the stack
func (m *MinStack) Push(value int) {
	m.maxSize++
	m.stk.Push(value)
}

// returns minimum value in O(1)
func (m *MinStack) Min() int {
	// write your code here
	tempStack := new(Stack)

	min := m.stk.Top()
	for m.maxSize > 1 {
		tempStack.Push(m.Pop())
		temp := m.stk.Top()
		if temp < min {
			min = temp
		}

	}
	for !tempStack.IsEmpty() {
		m.Push(tempStack.Pop())
	}

	return min
}

func MinStacker(arr []int) *MinStack {
	stk := &MinStack{
		maxSize: len(arr),
		stk:     Stack{arr},
	}

	return stk
}

func main() {
	arr := []int{5, 10, 4, 9, 3}
	fmt.Println(MinStacker(arr).Min())

}
