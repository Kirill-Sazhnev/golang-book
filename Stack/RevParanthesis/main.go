package main

import (
	"fmt"
	"math"
)

type Stack struct {
	s []rune
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

func (s *Stack) Push(value rune) {
	s.s = append(s.s, value)
}

func (s *Stack) Pop() rune {
	if s.IsEmpty() {
		fmt.Print("Stack in empty.")
		return ' '
	}
	length := len(s.s)
	res := s.s[length-1]
	s.s = s.s[:length-1]
	return res
}

func (s *Stack) Top() rune {
	if s.IsEmpty() {
		fmt.Print("Stack in empty.")
		return ' '
	}
	length := len(s.s)
	res := s.s[length-1]
	return res
}

//You can "math" library to solve the problem

func reverseParenthesis(expn string, size int) int {

	//Uncomment this line of code and use ch to iterate in a string

	stk := new(Stack)
	var ch rune
	countLeft := 0
	countRight := 0

	for _, ch = range expn {

		switch {
		case isOpen(ch):
			stk.Push(ch)
		case !stk.IsEmpty() && Pair(ch) == stk.Top():
			stk.Pop()
		case ch == ')':
			countRight++
		}

	}
	countLeft = stk.Length()

	switch {
	case (countLeft+countRight)%2 != 0:
		return -1
	case countLeft == countRight:
		return countLeft + countRight
	case countLeft != countRight:
		return int(math.Ceil(float64(countLeft)/2) + math.Ceil(float64(countRight)/2))
	}

	return 0 //Return 0 if string is empty
	//Return -1 if string's length is odd
}

func isOpen(val rune) bool {
	switch val {
	case '(', '{', '[':
		return true
	default:
		return false

	}
}

func Pair(val rune) rune {
	switch val {
	case ')':
		return '('
	case '}':
		return '{'
	case ']':
		return '['
	default:
		return 0
	}
}

func main() {
	str := ")(())(((“"
	size := len(str)

	fmt.Println(reverseParenthesis(str, size))
	fmt.Println(size)
}
