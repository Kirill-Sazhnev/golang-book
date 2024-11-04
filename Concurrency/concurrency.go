package main

import (
	"fmt"
	"sync"
	"time"
)

func Merge(left, right []int) []int {
	merged := make([]int, 0, len(left)+len(right))
	for len(left) > 0 || len(right) > 0 {
		if len(left) == 0 {
			return append(merged, right...)
		} else if len(right) == 0 {
			return append(merged, left...)
		} else if left[0] < right[0] {
			merged = append(merged, left[0])
			left = left[1:]
		} else {
			merged = append(merged, right[0])
			right = right[1:]
		}
	}
	return merged
}

func MergeSort(data []int) []int {
	if len(data) <= 1 {
		return data
	}

	mid := len(data) / 2
	leftCh := make(chan []int)

	go func() {
		leftCh <- MergeSort(data[:mid])
	}()

	left := <-leftCh
	right := MergeSort(data[mid:])
	close(leftCh)

	return Merge(left, right)
}

func printTable(n int, wg *sync.WaitGroup) {
	for i := 1; i <= 12; i++ {
		fmt.Printf("%d x %d = %d\n", i, n, n*i)
	}
	wg.Done()
}

type Node struct {
	Data  int
	Sleep time.Duration
	Left  *Node
	Right *Node
}

var treeTraversal []int

var wg sync.WaitGroup

func (n *Node) ProcessNodeParallel() {

	defer wg.Done()

	for i := 0; i < 10000; i++ {
		time.Sleep(n.Sleep)

	}
	treeTraversal = append(treeTraversal, n.Data)

}

func (n *Node) TreeTraversalParallel() {

	defer wg.Done()

	//Write your code here
	if n.Left != nil {
		wg.Add(1)
		go n.Left.TreeTraversalParallel()
	}
	wg.Add(1)
	n.ProcessNodeParallel()
	if n.Right != nil {
		wg.Add(1)
		go n.Right.TreeTraversalParallel()
	}

}
