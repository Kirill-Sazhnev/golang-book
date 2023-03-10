package main

import "fmt"

type node struct {
	nvalue *Node
	next   *node
}

type Stack struct {
	head *node
	size int
}

func (s *Stack) Length() int {
	return s.size
}

func (s *Stack) IsEmpty() bool {
	return s.size == 0
}

func (s *Stack) Peek() (*Node, bool) {
	if s.IsEmpty() {
		fmt.Println("StackEmptyException")
		return nil, false
	}
	return s.head.nvalue, true
}

func (s *Stack) Push(nvalue *Node) {
	s.head = &node{&Node{nvalue.value, nvalue.left, nvalue.right}, s.head}
	s.size++
}

func (s *Stack) Pop() *Node {
	if s.IsEmpty() {
		fmt.Println("StackEmptyException")
		return nil
	}
	temp := new(Node)
	temp.value = s.head.nvalue.value
	temp.left = s.head.nvalue.left
	temp.right = s.head.nvalue.right
	s.head = s.head.next
	s.size--
	return temp
}

type Node struct {
	value       int
	left, right *Node
}

type Tree struct {
	root *Node
}

func LevelOrderBinaryTree(arr []int) *Tree {
	tree := new(Tree)
	tree.root = levelOrderBinaryTree(arr, 0, len(arr))
	return tree
}

func levelOrderBinaryTree(arr []int, start int, size int) *Node {
	curr := &Node{arr[start], nil, nil}
	left := 2*start + 1
	right := 2*start + 2
	if left < size {
		curr.left = levelOrderBinaryTree(arr, left, size)
	}
	if right < size {
		curr.right = levelOrderBinaryTree(arr, right, size)
	}
	return curr
}

func (t *Tree) PrintDepthFirst() {

	stk := new(Stack)
	if t.root != nil {
		stk.Push(t.root)
	}
	for !stk.IsEmpty() {
		node := stk.Pop()
		if node.right != nil {
			stk.Push(node.right)
		}
		if node.left != nil {
			stk.Push(node.left)
		}
		fmt.Print(node.value, " ")
	}
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	tree := LevelOrderBinaryTree(arr)
	tree.PrintDepthFirst()
}
