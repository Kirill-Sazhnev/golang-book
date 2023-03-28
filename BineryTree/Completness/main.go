package main

import "fmt"

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

func (t *Tree) NumNodes() int {
	return numNodes(t.root)
}

func numNodes(curr *Node) int {
	if curr == nil {
		return 0
	}
	return (1 + numNodes(curr.left) + numNodes(curr.right))
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

func (t *Tree) IsCompleteTree() bool {
	return isCompleteTree(t.root)
}

func isCompleteTree(root *Node) bool {
	var que *Queue
	que1 := new(Queue)
	que2 := new(Queue)

	if root != nil {
		que1.Enqueue(root)
	}
	que1Length := que1.Length()
	que2Length := 0

	for {
		for !que1.IsEmpty() {

			root = que1.Dequeue().(*Node)

			if root.left != nil {
				que2.Enqueue(root.left)
			}
			if root.right != nil {
				que2.Enqueue(root.right)
			}
		}
		que2Length = que2.Length()

		if que2Length != que1Length*2 {
			que = que2
			break
		}

		for !que2.IsEmpty() {

			root = que2.Dequeue().(*Node)

			if root.left != nil {
				que1.Enqueue(root.left)
			}
			if root.right != nil {
				que1.Enqueue(root.right)
			}
		}

		que1Length = que1.Length()
		if que1Length != que2Length*2 {
			que = que1
			break
		}
	}
	for !que.IsEmpty() {
		root = que.Dequeue().(*Node)

		if root.right != nil || root.left != nil {
			return false
		}
	}

	return true
}

func (t *Tree) IsCompleteTree2() bool {
	count := t.NumNodes()
	fmt.Println(count)
	return isCompleteTreeUtil(t.root, 0, count)
}

func isCompleteTreeUtilMy(curr *Node, index int, count int) bool {
	//Implement your solution here
	var isComp bool
	if curr != nil {
		parent := (count - 2) / 2

		if (count-1)%2 != 0 {
			if parent == index {
				return curr.left != nil && curr.right == nil
			}
		} else {
			if parent == index {
				return curr.left != nil && curr.right != nil
			}
		}

		if curr.left != nil {
			isComp = isCompleteTreeUtilMy(curr.left, (index+1)*2-1, count)
		}

		if curr.right != nil && !isComp {
			isComp = isCompleteTreeUtilMy(curr.right, (index+1)*2, count)
		}

	}
	return isComp
}

func isCompleteTreeUtil(curr *Node, index int, count int) bool {
	if curr == nil {
		return true
	}
	if index > count {
		return false
	}
	return isCompleteTreeUtil(curr.left, index*2+1, count) &&
		isCompleteTreeUtil(curr.right, index*2+2, count)
}

func (t *Tree) IsHeap() bool {
	parentValue := -99999999
	return t.IsCompleteTree() && isHeapUtil(t.root, parentValue)
}

func isHeapUtil(curr *Node, parentValue int) bool {
	//Implement your solution here
	if curr == nil {
		return true
	}
	if curr.value < parentValue {
		return false
	}

	return isHeapUtil(curr.left, curr.value) && isHeapUtil(curr.right, curr.value)
}

func isHeapUtilQue(curr *Node, parentValue int) bool {
	//Implement your solution here
	que := new(Queue)
	que.Enqueue(curr)

	for !que.IsEmpty() {

		node := que.Dequeue().(*Node)
		parentValue := node.value

		if node.left != nil {
			que.Enqueue(node.left)
			if parentValue > node.left.value {
				return false
			}

		}
		if node.right != nil {
			que.Enqueue(node.right)
			if parentValue > node.right.value {
				return false
			}
		}
	}

	return true
}

func (t *Tree) TreeToListRec() *Node {
	return treeToListRec(t.root)
}

func treeToListRec(curr *Node) *Node {
	if curr == nil {
		return nil
	}
	if curr.right == nil && curr.left == nil {
		return curr
	}

	listTail := treeToListRec(curr.right)

	listHead := treeToListRec(curr.left)

	listNode := &Node{
		value: curr.value,
		left:  listHead,
		right: listTail,
	}

	if listHead.right == nil {
		listHead.right = listNode
	} else {
		node := listHead.right
		for node.right != nil {
			node = node.right
		}
		node.right = listNode
		listNode.left = node
	}

	return listHead
}

func (root *Node) PrintDLL() {
	if root == nil {
		fmt.Print("Empty List")
		return
	}
	curr := root
	tail := curr.left

	for curr != tail {
		fmt.Print(curr.value, " ")
		curr = curr.right
	}
	if curr != nil {
		fmt.Print(curr.value)
	}
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17}
	tree := LevelOrderBinaryTree(arr)

	list := tree.TreeToListRec()

	list.PrintDLL()
}
