package main

import "fmt"

type Node struct {
	value       int
	left, right *Node
}

type Tree struct {
	root *Node
}

func CreateBinarySearchTree(arr []int) *Tree {
	t := new(Tree)
	size := len(arr)
	t.root = createBinarySearchTreeUtil(arr, 0, size-1)
	return t
}

func createBinarySearchTreeUtil(arr []int, start int, end int) *Node {

	if start > end {
		return nil
	}

	head := (start + end) / 2

	curr := &Node{
		value: arr[head],
		left:  createBinarySearchTreeUtil(arr, start, head-1),
		right: createBinarySearchTreeUtil(arr, head+1, end),
	}

	return curr
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

func (t *Tree) Add(value int) {
	t.root = addNode(t.root, value)
}

func addNode(n *Node, value int) *Node {
	if n == nil {
		n := &Node{
			value: value,
		}
		return n
	}

	if value < n.value {
		n.left = addNode(n.left, value)
	}
	if value > n.value {
		n.right = addNode(n.right, value)
	}

	return n
}

func (t *Tree) Find(value int) bool {

	var curr *Node = t.root
	for curr != nil {
		if curr.value == value {
			return true
		} else if curr.value > value {
			curr = curr.left
		} else {
			curr = curr.right
		}
	}
	return false
}

func (t *Tree) FindMinNode() *Node {

	var node *Node = t.root

	for node.left != nil {
		node = node.left
	}

	return node
}

func (t *Tree) FindMaxNode() *Node {
	var node *Node = t.root

	for node.right != nil {
		node = node.right
	}

	return node
}

func IsBST(root *Node) bool {
	if root == nil {
		return true
	}
	if root.left != nil && root.left.value > root.value {
		return false
	}
	if root.right != nil && root.right.value < root.value {
		return false
	}
	return IsBST(root.left) && IsBST(root.right)
}

func (t *Tree) DeleteNode(value int) {
	t.root = DeleteNode(t.root, value)
}

func DeleteNode(node *Node, value int) *Node {
	if node == nil {
		return node
	}
	switch {
	case node.value == value:
		switch {
		case node.left == nil && node.right == nil:
			return nil
		case node.left == nil:
			return node.right
		case node.right == nil:
			return node.left
		default:
			tempNode := node.left
			for tempNode.right != nil {
				tempNode = tempNode.right
			}
			node.value = tempNode.value
			node.left = DeleteNode(node.left, tempNode.value)
		}

	case node.right != nil && value > node.value:
		node.right = DeleteNode(node.right, value)
	case node.left != nil && value < node.value:
		node.left = DeleteNode(node.left, value)
	}
	return node
}

func (t *Tree) LcaBST(first int, second int) (int, bool) {
	return LcaBST(t.root, first, second)
}

func LcaBSTmy(curr *Node, first int, second int) (int, bool) {
	if curr == nil {
		return 0, false
	}
	LCA := curr.value
	nextFirst := curr
	nextSecond := curr

	switch {
	case curr.value > first:
		nextFirst = curr.left
	case curr.value < first:
		nextFirst = curr.right
	}
	switch {
	case curr.value > second:
		nextSecond = curr.left
	case curr.value < second:
		nextSecond = curr.right
	}
	if nextFirst == nextSecond && nextFirst != nil {
		LCA = nextFirst.value
	}
	for nextFirst != nil && nextSecond != nil && nextFirst.value != first && nextSecond.value != second {
		switch {
		case nextFirst != nil && nextFirst.value > first:
			nextFirst = nextFirst.left
		case nextFirst != nil && nextFirst.value < first:
			nextFirst = nextFirst.right
		}
		switch {
		case nextSecond != nil && nextSecond.value > second:
			nextSecond = nextSecond.left
		case nextSecond != nil && nextSecond.value < second:
			nextSecond = nextSecond.right
		}
		if nextFirst == nextSecond && nextFirst != nil {
			LCA = nextFirst.value
		}
	}
	return LCA, true
}

func LcaBST(curr *Node, first int, second int) (int, bool) {
	if curr == nil {
		return 0, false
	}
	if curr.value < first && curr.value < second {
		return LcaBST(curr.right, first, second)
	}
	if curr.value > first && curr.value > second {
		return LcaBST(curr.left, first, second)
	}
	return curr.value, true
}

func (t *Tree) PrintDataInRange(min int, max int) {
	printDataInRange(t.root, min, max)
}

func printDataInRange(root *Node, min int, max int) {
	if root == nil {
		return
	}
	printDataInRange(root.left, min, max)

	if root.value >= min && root.value <= max {
		fmt.Print(root.value, " ")
	}

	printDataInRange(root.right, min, max)

}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	tree := CreateBinarySearchTree(arr)

	fmt.Println(tree.LcaBST(3, 4))
}
