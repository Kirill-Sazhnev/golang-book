package btree

import (
	"fmt"
	Que "golang-book/Queue/ReverseQueue"
	"math"
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
	que := new(Que.Queue)
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

	que := new(Que.Queue)
	que2 := new(Que.Queue)

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
	que := new(Que.Queue)
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

func (t *Tree) NthPreOrder(index int) {
	var counter int
	nthPreOrder(t.root, index, &counter)
}

func nthPreOrder(node *Node, index int, counter *int) {
	//Implement your solution here
	if node != nil {
		*counter++
		if *counter == index {
			fmt.Print(node.value, " ")
			return
		}
		nthPreOrder(node.left, index, counter)
		nthPreOrder(node.right, index, counter)
	}
}

func (t *Tree) NthPostOrder(index int) {
	var counter int
	nthPostOrder(t.root, index, &counter)
}

func nthPostOrder(node *Node, index int, counter *int) {
	//Implement your solution here
	if node != nil {

		nthPostOrder(node.left, index, counter)
		nthPostOrder(node.right, index, counter)
		*counter++
		if *counter == index {
			fmt.Print(node.value, " ")
			return
		}

	}
}

func (t *Tree) NthInOrder(index int) {
	var counter int
	nthInOrder(t.root, index, &counter)
}

func nthInOrder(node *Node, index int, counter *int) {
	//Implement your solution here

	if node != nil {

		nthInOrder(node.left, index, counter)
		*counter++
		if *counter == index {
			fmt.Print(node.value, " ")
			return
		}
		nthInOrder(node.right, index, counter)

	}
}

func (t *Tree) NumNodes() int {
	return numNodes(t.root)
}

func numNodesQue(curr *Node) int {
	if curr == nil {
		return 0
	}
	que := new(Que.Queue)
	que.Enqueue(curr)
	counter := 0
	for !que.IsEmpty() {
		curr = que.Dequeue().(*Node)
		counter++
		if curr.right != nil {
			que.Enqueue(curr.right)
		}
		if curr.left != nil {
			que.Enqueue(curr.left)
		}
	}
	return counter //Kindly change the return value as per your needs
}

func numNodes(curr *Node) int {
	if curr == nil {
		return 0
	}

	return 1 + numNodes(curr.left) + numNodes(curr.right) //Kindly change the return value as per your needs
}

func (t *Tree) SumAllBT() int {
	return sumAllBT(t.root)
}

func sumAllBT(curr *Node) int {
	if curr == nil {
		return 0
	}

	return curr.value + sumAllBT(curr.left) + sumAllBT(curr.right) //Kindly change the return value as per need
}

func (t *Tree) NumLeafNodes() int {
	return numLeafNodes(t.root)
}

func numLeafNodes(curr *Node) int {
	if curr == nil {
		return 0
	}
	if curr.left == nil && curr.right == nil {
		return 1
	}
	return numLeafNodes(curr.right) + numLeafNodes(curr.left) //Kindly change the return value as per needs
}

func (t *Tree) NumFullNodesBT() int {
	return numFullNodesBT(t.root)
}

func numFullNodesBT(curr *Node) int {

	if curr == nil {
		return 0
	}

	var count int

	if curr.left != nil && curr.right != nil {
		count++
	}

	return count + numFullNodesBT(curr.left) + numFullNodesBT(curr.right)
}

func (t *Tree) SearchBT(value int) bool {
	return searchBT(t.root, value)
}

func searchBT(root *Node, value int) bool {
	if root == nil {
		return false
	}

	if root.value == value {
		return true
	}

	return searchBT(root.left, value) || searchBT(root.right, value)
}

func (t *Tree) FindMaxBT() int {
	return findMaxBT(t.root)
}

func findMaxBT(curr *Node) int {
	if curr == nil {
		return math.MinInt32
	}

	max := findMaxBT(curr.left)

	if curr.value > max {
		max = curr.value
	}
	maxRight := findMaxBT(curr.right)
	if maxRight > max {
		max = maxRight
	}

	return max
}

func (t *Tree) TreeDepth() int {
	return treeDepth(t.root)
}

func treeDepth(root *Node) int {
	if root == nil {
		return 0
	}

	leftE := treeDepth(root.left)
	rightE := treeDepth(root.right)

	if root.left != nil && leftE >= rightE {
		return 1 + leftE
	}

	if root.right != nil && rightE >= leftE {
		return 1 + rightE
	}

	return 0
}

func (t *Tree) IsEqual(t2 *Tree) bool {
	return isEqual(t.root, t2.root)
}

func isEqual(node1 *Node, node2 *Node) bool {
	if node1 == nil && node2 == nil {
		return true
	}
	if node1 == nil || node2 == nil {
		return false
	}

	return (node1.value == node2.value && isEqual(node1.left, node2.left) == isEqual(node1.right, node2.right))
}

func (t *Tree) CopyTree() *Tree {
	Tree2 := new(Tree)
	Tree2.root = copyTree(t.root)
	return Tree2
}

func copyTree(curr *Node) *Node {

	if curr == nil {
		return nil
	}

	newNode := &Node{
		value: curr.value,
		left:  copyTree(curr.left),
		right: copyTree(curr.right),
	}

	return newNode
}

func (t *Tree) CopyMirrorTree() *Tree {
	tree := new(Tree)
	tree.root = copyMirrorTree(t.root)
	return tree
}

func copyMirrorTree(curr *Node) *Node {
	if curr == nil {
		return nil
	}

	newNode := &Node{
		value: curr.value,
		left:  copyMirrorTree(curr.right),
		right: copyMirrorTree(curr.left),
	}

	return newNode
}

func (t *Tree) Free() {
	var emptyNode *Node
	t.root = emptyNode
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	tree := LevelOrderBinaryTree(arr)
	fmt.Println(tree.TreeDepth())

}
