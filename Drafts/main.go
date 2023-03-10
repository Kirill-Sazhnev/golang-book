package main

import "fmt"

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

func (t *Tree) PrintLevelOrderLineByLine() {

	que1 := new(Queue)
	que2 := new(Queue)
	var temp *Node
	if t.root != nil {
		que1.Enqueue(t.root)
	}

	for que1.Length() != 0 || que2.Length() != 0 {

		for que1.Length() != 0 {
			temp2 := que1.Dequeue()
			temp = temp2.(*Node)
			fmt.Print(temp.value, " ")
			if temp.left != nil {
				que2.Enqueue(temp.left)
			}
			if temp.right != nil {
				que2.Enqueue(temp.right)
			}
		}

		fmt.Println(" ")

		for que2.Length() != 0 {
			temp2 := que2.Dequeue()
			temp = temp2.(*Node)
			fmt.Print(temp.value, " ")
			if temp.left != nil {
				que1.Enqueue(temp.left)
			}
			if temp.right != nil {
				que1.Enqueue(temp.right)
			}
		}

		fmt.Println(" ")

	}
}

func main() {

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	t := LevelOrderBinaryTree(arr)
	t.PrintLevelOrderLineByLine()
}
