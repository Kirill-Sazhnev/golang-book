package main

import "fmt"

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
	length := len(s.s)
	res := s.s[length-1]
	s.s = s.s[:length-1]
	return res
}

func (s *Stack) Top() int {
	length := len(s.s)
	res := s.s[length-1]
	return res
}

type QueueUsingStack struct {
	stk1 Stack
	// add more values here if need be
	size int
	head int
	tail int
}

func (que *QueueUsingStack) Add(value int) {
	//Implement your solution here
	que.stk1.Push(value)
	if que.IsEmpty() {
		que.head = 0
		que.tail = 0
	} else {
		que.tail++
	}

	que.size++

}

func (que *QueueUsingStack) Remove() int {
	//Implement your solution here
	if que.IsEmpty() {
		fmt.Println("The queue is empty")
		return 0
	}
	value := que.stk1.s[que.head]
	que.size--
	que.head++
	return value
}

func (que *QueueUsingStack) Length() int {
	//Implement your solution here

	return que.size
}

func (que *QueueUsingStack) IsEmpty() bool {
	//Implement your solution here

	return que.size == 0
}

func main() {
	q := new(QueueUsingStack)
	q.Add(10)
	q.Add(20)
	q.Add(30)
	q.stk1.Print()
	fmt.Println(q.head)
	fmt.Println(q.tail)

	fmt.Println(q.IsEmpty())
	fmt.Println(q.Length())
	fmt.Println(q.Remove())
	q.stk1.Print()
	fmt.Println(q.Remove())
	fmt.Println(q.Remove())
}
