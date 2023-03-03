package main

import "fmt"

type Stack struct {
	s []byte
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

func (s *Stack) Push(value byte) {
	s.s = append(s.s, value)
}

func (s *Stack) Pop() byte {
	if s.IsEmpty() == true {
		return ' '
	}
	length := len(s.s)
	res := s.s[length-1]
	s.s = s.s[:length-1]
	return res
}

func (s *Stack) Top() byte {
	if s.IsEmpty() == true {
		return ' '
	}
	length := len(s.s)
	res := s.s[length-1]
	return res
}

func maxDepthParenthesis(expn string, size int) int {
	//Implement your solution here
	stk := new(Stack)
	//Uncomment the ch variable and use it to move in an expression array
	var ch byte
	count := 0
	for i := range expn {
		ch = expn[i]
		switch {
		case isOpen(ch):
			stk.Push(ch)
		case !isOpen(ch) && Pair(ch) == stk.Top():
			if count < stk.Length() {
				count = stk.Length()
			}
			stk.Pop()
		}
	}

	return count - stk.Length() //Return 0 if depth of parenthesis is zero
}

func isOpen(val byte) bool {
	switch val {
	case '(', '{', '[':
		return true
	default:
		return false

	}
}

func Pair(val byte) byte {
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
	str := "(((((()))()[[[[]]]]]"
	size := len(str)
	fmt.Println(maxDepthParenthesis(str, size))
}
