package main

import "container/heap"

type Set struct {
	n1 int
	n2 int
	n3 int
}

type MinHeap []Set

// newMinHeap intializes an instance of MinHeap
func newMinHeap() *MinHeap {
	min := &MinHeap{}
	heap.Init(min)
	return min
}

// Len function returns the length of MinHeap
func (h MinHeap) Len() int {
	return len(h)
}

// Empty function returns true of empty, false otherwise
func (h MinHeap) Empty() bool {
	return len(h) == 0
}

// Less function compares the two elements of MinHeap given their indices
func (h MinHeap) Less(i, j int) bool {
	return h[i].n1 < h[j].n1
}

// Swap function swaps the values of the elements whose indices are given
func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// Top function returns the element at the top of the MinHeap
func (h MinHeap) Top() interface{} {
	return h[0]
}

// Push function pushes the given element into the MinHeap
func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(Set))
}

// Pop function pops the top element of MinHeap
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
