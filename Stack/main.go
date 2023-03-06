package main

import (
	"fmt"
)

type StackInt struct {
	s []int
}

type Node struct {
	value int
	next  *Node
}

type StackLinkedList struct {
	head *Node
	size int
}

// isEmpty() function returns true if stack is empty or false in all other cases.
func (s *StackInt) IsEmpty() bool {
	//Implement your solution here
	return len(s.s) == 0 //Kindly make changes according to your needs
}

// length() function returns the number of elements in the stack.
func (s *StackInt) Length() int {
	//Implement your solution here

	return len(s.s) //Kindly make changes according to your needs
}

// The print function will print the elements of the array.
func (s *StackInt) Print() {
	//Implement your solution here
	fmt.Println(s.s)
}

// push() function append value to the data.
func (s *StackInt) Push(value int) {
	//Implement your solution here
	s.s = append(s.s, value)
}

/* In the pop() function, first it will check that the stack is not empty.
Then it will pop the value from the data array and return it. */

func (s *StackInt) Pop() int {
	//Implement your solution here
	if !s.IsEmpty() {
		lng := len(s.s) - 1
		val := s.s[lng]
		s.s = s.s[:lng]
		return val
	}
	fmt.Print("Stack is empty.")
	return 0 //Kindly make changes according to your needs
}

/*
top() function will first check that the stack is not empty
then returns the value stored in the top element
of stack (does not remove it).
*/
func (s *StackInt) Top() int {
	//Implement your solution here
	if !s.IsEmpty() {
		l := len(s.s) - 1
		return s.s[l]
	}
	fmt.Print("Stack is empty.")
	return 0 //Kindly make changes according to your needs
}

//==========================================================

// Size() function will return the size of the linked list.
func (s *StackLinkedList) Size() int {
	//Implement your solution here

	return s.size //Return 0 if stack is empty
}

/*
	IsEmpty() function will return true is size of the linked list is

equal to zero or false in all other cases.
*/
func (s *StackLinkedList) IsEmpty() bool {
	//Implement your solution here

	return s.size == 0 //Return true if stack is empty
}

/*
First, the Peek() function will check if the stack is empty.
If not, it will return the peek value of stack i.e., will return the
head value of the linked list.
*/
func (s *StackLinkedList) Peek() (int, bool) {
	//Implement your solution here
	if s.IsEmpty() {
		fmt.Println("The stack is empty")
		return 0, false
	}
	return s.head.value, true //Return 0,true if stack is empty
}

// Push() function  will add new value at the start of the linked list.
func (s *StackLinkedList) Push(value int) {
	//Implement your solution here
	s.head = &Node{
		value: value,
		next:  s.head,
	}
	s.size++
}

/*
In the pop() function, first it will check that the stack is not empty.
Then it will pop the value from the linked list and return it.
*/
func (s *StackLinkedList) Pop() (int, bool) {
	//Implement your solution here
	if s.IsEmpty() {
		fmt.Println("The stack is empty")
		return 0, false
	}

	head := s.head.value
	s.head = s.head.next
	s.size--
	return head, true //Return true if stack is empty
}

/* Print() function will print the elements of the linked list. */
func (s *StackLinkedList) Print() {
	//Implement your solution here
	valAdrs := s.head
	for i := 0; i < s.size; i++ {
		fmt.Printf("%v ", valAdrs.value)
		valAdrs = valAdrs.next
	}
	fmt.Println()

}

func sortedInsertMy(stk *StackInt, element int) {
	//Implement your solution here
	if stk.Top() <= element {
		stk.Push(element)
		return
	}

	for i := stk.Length() - 1; i >= 0; i-- {
		if stk.s[i] <= element {
			right := stk.s[i+1:]
			left := append([]int{}, stk.s[:i+1]...)
			left = append(left, element)
			stk.s = append(left, right...)
			return
		}

	}
}

func sortedInsert(stk *StackInt, element int) {
	var temp int
	if stk.Length() == 0 || element > stk.Top() {
		stk.Push(element)
	} else {
		temp = stk.Pop()
		sortedInsert(stk, element)
		stk.Push(temp)
	}
}

func sortStack(stk *StackInt) {
	var temp int
	if stk.Length() > 0 {
		temp = stk.Pop()
		sortStack(stk)
		sortedInsert(stk, temp)
	}
}

func sortStackMy(stk *StackInt) {
	//Implement your solution here
	if stk.Length() < 2 {
		return
	}
	for i := 0; i < 2; i++ {
		head := stk.Pop()
		if head > stk.Top() {
			sortStack(stk)
			stk.Push(head)
		} else {
			temp := stk.Pop()
			stk.Push(head)
			sortStack(stk)
			stk.Push(temp)
		}
	}
}

func bottomInsert(stk *StackInt, element int) {
	//Implement your solution here
	var temp int

	/*for i, val := range stk.s {
		temp = val
		stk.s[i] = element
		element = temp
	}
	stk.s = append(stk.s, element) */

	if stk.Length() == 0 {
		stk.Push(element)
	} else {
		temp = stk.Pop()
		bottomInsert(stk, element)
		stk.Push(temp)
	}

}

func reverseStack(stk *StackInt) {
	//Implement your solution here

	if stk.Length() == 1 {
		return
	}

	temp := stk.Pop()
	reverseStack(stk)
	bottomInsert(stk, temp)
}

func reverseKElementInStack(stk *StackInt, k int) {
	//Implement your solution here
	tempStk := &StackInt{}

	for i := 0; i < k; i++ {
		tempStk.Push(stk.Pop())
	}

	reverseStack(tempStk)

	for tempStk.Length() > 0 {
		stk.Push(tempStk.Pop())
	}
}

func reverseKElementInStackQueue(stk *StackInt, k int) {
	//Implement your solution here
	tempQue := &Queue{}

	for i := 0; i < k && !stk.IsEmpty(); i++ {
		tempQue.Enqueue(stk.Pop())
	}

	for !tempQue.IsEmpty() {
		stk.Push(tempQue.Dequeue().(int))
	}
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
		return ' '
	}
	length := len(s.s)
	res := s.s[length-1]
	s.s = s.s[:length-1]
	return res
}

func (s *Stack) Top() rune {
	if s.IsEmpty() {
		return ' '
	}
	length := len(s.s)
	res := s.s[length-1]
	return res
}

func IsBalancedParenthesis(expn string) bool {
	//Implement your solution here
	stk := new(Stack)

	for _, val := range expn {
		switch {
		case isOpen(val):
			stk.Push(val)
		case !isOpen(val) && Pair(val) == stk.Top():
			stk.Pop()
		default:
			return false
		}
	}

	return stk.IsEmpty() //Return true if parentheses are balanced
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

func longestContBalParen(str string, size int) int {
	//Implement your solution here
	stk := new(StackInt)
	stk.Push(-1)
	length := 0
	for i := 0; i < size; i++ {
		if str[i] == '(' {
			stk.Push(i)
		} else // string[i] == ')'
		{
			stk.Pop()
			if stk.Length() != 0 {
				if length < i-stk.Top() {
					length = i - stk.Top()
				}
			} else {
				stk.Push(i)
			}
		}
	}
	return length
}

func findDuplicateParenthesis(expn string, size int) bool {

	stk := new(StackInt)
	pair := 0
	for i := 0; i < size; i++ {
		switch {
		case expn[i] == '(':
			stk.Push(i)
		case expn[i] == ')' && !stk.IsEmpty():
			stk.Pop()
			pair++
			if i-stk.Top() < 3 {
				return true
			}
		}

	}

	if !stk.IsEmpty() || pair > 1 {
		return true
	}
	return false
}

func printParenthesisNumber(expn string, size int) {
	//Uncomment the ch variable and use it iterate through the string
	var ch byte
	stk := new(StackInt)
	index := 0
	output := ""
	for i := 0; i < size; i++ {
		ch = expn[i]
		switch ch {
		case '(':
			index++
			stk.Push(index)
			output += fmt.Sprintf("%v", index)
		case ')':
			output += fmt.Sprintf("%v", stk.Pop())
		}
	}

	//use output variable to save and print the output string

	//Implement your solution here

	fmt.Println(output)
}

func InfixToPostfix(expn string) string {

	//kindly replace the output string with empty string while
	//implementing the solution
	output := "No Output"
	stkS := new(StackInt)
	input := ""
	for _, rval := range expn {
		val := int(rval)
		switch val {
		case '(':
			stkS.Push(val)
		case '+', '-', '*', '/', '^', '%':
			for !stkS.IsEmpty() && priority(stkS.Top()) >= priority(val) {
				input += fmt.Sprintf(" %c", stkS.Pop())
			}
			stkS.Push(val)
			input += " "
		case ')':
			for stkS.Top() != '(' {
				input += fmt.Sprintf(" %c", stkS.Pop())
			}
			stkS.Pop()
		default:
			input += fmt.Sprintf("%c", val)
		}
	}
	for !stkS.IsEmpty() {
		input += fmt.Sprintf(" %c", stkS.Pop())
	}

	if input != "" {
		output = input
	}
	return output
}

func priority(sign int) int {
	prt := 0
	switch sign {
	case '^':
		prt = 3
	case '*', '/', '%':
		prt = 2
	case '+', '-':
		prt = 1
	}
	return prt
}

func InfixToPrefix(expn string) string {
	//Implement your solution here

	stk := new(StackInt)
	input := ""
	output := ""

	for i := len(expn) - 1; i >= 0; i-- {
		val := int(expn[i])
		switch val {
		case ')':
			stk.Push(val)
		case '(':
			for !stk.IsEmpty() && stk.Top() != ')' {
				input += fmt.Sprintf(" %c", stk.Pop())
			}
			stk.Pop()

		case '+', '-', '*', '/', '^', '%':
			for !stk.IsEmpty() && priority(stk.Top()) >= priority(val) {
				input += fmt.Sprintf(" %c", stk.Pop())
			}
			stk.Push(val)
			input += " "

		default:
			input += fmt.Sprintf("%c", val)
		}
	}

	for !stk.IsEmpty() {
		input += fmt.Sprintf("%c ", stk.Pop())
	}

	if input != "" {
		for i := len(input) - 1; i >= 0; i-- {
			output += string(input[i])
		}
		return output
	}

	return " No Output "
}

func main() {
	str := "10+((3))*5/(16-4)"
	fmt.Println(InfixToPrefix(str))

	/*
		stack := new(StackLinkedList)

		fmt.Println(stack.IsEmpty())
		stack.Push(1)
		stack.Push(2)
		fmt.Println(stack.Size())
		stack.Print()
		stack.Push(3)
		stack.Print()
		fmt.Println(stack.Pop())
		stack.Print()
		fmt.Println(stack.Peek())
	*/
}
