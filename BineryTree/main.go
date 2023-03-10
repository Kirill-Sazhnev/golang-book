package main

import (
	"fmt"
	Queue "golang-book/Queue/ReverseQueue"
)

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
	if start >= size {
		return nil
	}
	next := (start+1)*2 - 1
	parent := &Node{
		value: arr[start],
		left:  levelOrderBinaryTree(arr, next, size),
		right: levelOrderBinaryTree(arr, next+1, size),
	}

	return parent
}

func (t *Tree) PrintPreOrder() {
	printPreOrder(t.root)
}

func printPreOrder(n *Node) {
	fmt.Print(n.value, " ")

	if n.left != nil {
		printPreOrder(n.left)
	}
	if n.right != nil {
		printPreOrder(n.right)
	}
	//Implement your solution here
}

func (t *Tree) PrintPostOrder() {
	printPostOrder(t.root)
}

func printPostOrder(n *Node) {

	//Implement your solution here
	if n.left != nil {
		printPostOrder(n.left)
	}
	if n.right != nil {
		printPostOrder(n.right)
	}
	fmt.Print(n.value, " ")
	//Implement your solution here
}

func (t *Tree) PrintInOrder() {
	printInOrder(t.root)
}

func printInOrder(n *Node) {

	if n.left != nil {
		printInOrder(n.left)
	}
	fmt.Print(n.value, " ")
	if n.right != nil {
		printInOrder(n.right)
	}
}

func (t *Tree) PrintBreadthFirst() {
	que := new(Queue.Queue)
	que.Enqueue(t.root)
	for !que.IsEmpty() {
		node := que.Dequeue().(*Node)
		fmt.Print(node.value, " ")
		if node.left != nil {
			que.Enqueue(node.left)
		}
		if node.right != nil {
			que.Enqueue(node.right)
		}

	}

}

func (t *Tree) PrintLevelOrderLineByLine() {

	que := new(Queue.Queue)
	que2 := new(Queue.Queue)

	que.Enqueue(t.root)

	for !que.IsEmpty() || !que2.IsEmpty() {

		for !que.IsEmpty() {
			node := que.Dequeue().(*Node)
			fmt.Print(node.value, " ")

			if node.left != nil {
				que2.Enqueue(node.left)
			}
			if node.right != nil {
				que2.Enqueue(node.right)
			}
		}

		fmt.Print("; ")

		for !que2.IsEmpty() {
			node := que2.Dequeue().(*Node)
			fmt.Print(node.value, " ")

			if node.left != nil {
				que.Enqueue(node.left)
			}
			if node.right != nil {
				que.Enqueue(node.right)
			}
		}

		fmt.Print("; ")

	}
}

func (t *Tree) PrintLevelOrderLineByLine2() {
	que := new(Queue.Queue)
	que.Enqueue(t.root)

	for !que.IsEmpty() {
		counter := que.Length()
		for counter > 0 {
			node := que.Dequeue().(*Node)
			fmt.Print(node.value, " ")
			counter--

			if node.left != nil {
				que.Enqueue(node.left)
			}
			if node.right != nil {
				que.Enqueue(node.right)
			}
		}
		fmt.Print("; ")
	}
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	tree := LevelOrderBinaryTree(arr)
	tree.PrintLevelOrderLineByLine2()
}
