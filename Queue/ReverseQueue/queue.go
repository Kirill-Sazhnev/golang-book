// some explanantion for the package
package queue

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

const capacity = 100

type Queue struct {
	size  int
	data  [capacity]interface{}
	front int
	back  int
}

func (q *Queue) Enqueue(value interface{}) {
	if q.size >= capacity {
		return
	}
	q.size++
	q.data[q.back] = value
	q.back = (q.back + 1) % (capacity - 1)
}

func (q *Queue) Dequeue() interface{} {
	var value interface{}
	if q.size <= 0 {
		return 0
	}
	q.size--
	value = q.data[q.front]
	q.front = (q.front + 1) % (capacity - 1)
	return value
}

func (q *Queue) IsEmpty() bool {
	return q.size == 0
}

func (q *Queue) Length() int {
	return q.size
}

func (q *Queue) Print() {
	if q.front == q.back {
		fmt.Print("Queue is empty.")
	}
	for i := q.front; i < q.back; i++ {
		fmt.Print(q.data[i], " ")
	}
	fmt.Println()
}

func reverseQueue(que *Queue) {
	//Implement your solution here
	stk := new(Stack)
	for que.Length() > 0 {
		stk.Push(que.Dequeue().(int))
	}
	for stk.Length() > 0 {
		que.Enqueue(stk.Pop())
	}
}

func reverseKElementInQueue(que *Queue, k int) {
	//Implement your solution here
	stk := new(Stack)
	tempque := new(Queue)
	for i := 0; i < k; i++ {
		stk.Push(que.Dequeue().(int))
	}
	for que.Length() > 0 {
		tempque.Enqueue(que.Dequeue())
	}
	for stk.Length() > 0 {
		que.Enqueue(stk.Pop())
	}
	for tempque.Length() > 0 {
		que.Enqueue(tempque.Dequeue())
	}
} /*
func main() {
	que := new(Queue)
	que.Enqueue(10)
	que.Enqueue(20)
	que.Enqueue(30)
	que.Enqueue(40)
	que.Enqueue(50)
	que.Print()
	reverseKElementInQueue(que, 3)
	que.Print()

}
*/
