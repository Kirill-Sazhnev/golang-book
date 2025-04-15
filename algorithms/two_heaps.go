package main

import (
	"container/heap"
	"fmt"
	"sort"
)

func main() {

	//arr := [][]int{{1,7},{8,13},{5,6},{10,14},{6,7}}
	arr := [][]int{{12, 13}, {13, 15}, {17, 20}, {13, 14}, {19, 21}, {18, 20}}
	fmt.Println(minimumMachines(arr))
}

func connectSticks(sticks []int) int {
	h := &MinHeap{}
	heap.Init(h)
	for _, stick := range sticks {
		heap.Push(h, stick)
	}

	cost := 0
	for h.Len() > 1 {
		stick1 := heap.Pop(h).(int)
		stick2 := heap.Pop(h).(int)
		newStick := stick1 + stick2
		cost += newStick
		heap.Push(h, newStick)
	}

	return cost
}

func minimumMachines(tasks [][]int) int {
	if len(tasks) == 0 {
		return 0
	}
	h := newHeap()
	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i][0] < tasks[j][0]
	})
	start, end := 0, 0
	counter := 0
	for _, t := range tasks {
		start = t[0]
		heap.Push(h, t)

		nearestTask := h.Top().([]int)
		end = nearestTask[1]

		if start < end {
			counter++
			continue
		}

		_ = heap.Pop(h).([]int)
	}

	return counter
}

// Heap structure intialization
type Heap [][]int

// newHeap function initializes an instance of the heap
func newHeap() *Heap {
	min := &Heap{}
	heap.Init(min)
	return min
}

// Len function returns the length of the heap
func (h Heap) Len() int {
	return len(h)
}

// Empty function returns true if empty, false otherwise
func (h Heap) Empty() bool {
	return len(h) == 0
}

// Less function compares two elements of the heap given their indexes
func (h Heap) Less(i, j int) bool {
	return h[i][1] < h[j][1]
}

// Swap function swaps the values of the elements whose indices are given
func (h Heap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// Top function returns the element at the top of the heap
func (h Heap) Top() interface{} {
	return h[0]
}

// Push function pushes an element into the heap
func (h *Heap) Push(x interface{}) {
	*h = append(*h, x.([]int))
}

// Pop function pops the element at the top of the heap
func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *Heap) Heapify(array [][]int) {
	for _, arr := range array {
		heap.Push(h, arr)
	}
}

func minEnd(i, j int) int {
	if i < j {
		return i
	}
	return j
}

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
