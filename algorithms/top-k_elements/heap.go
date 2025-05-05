package main

import (
	"container/heap"
	"strconv"
)

type Set struct {
	count int
	char  rune
}

// MaxHeap structure intialization
type MaxHeap []Set

// newMaxHeap function initializes an instance of MaxHeap
func newMaxHeap() *MaxHeap {
	max := &MaxHeap{}
	heap.Init(max)
	return max
}

// Len function returns the length of MaxHeap
func (h MaxHeap) Len() int {
	return len(h)
}

// Empty function returns true if the MaxHeap empty, false otherwise
func (h MaxHeap) Empty() bool {
	return len(h) == 0
}

// Less function compares two elements of MaxHeap given their indices
func (h MaxHeap) Less(i, j int) bool {
	return h[i].count > h[j].count
}

// Swap function swaps the value of the elements whose indices are given
func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// Top function returns the element at the top of the MaxHeap
func (h MaxHeap) Top() interface{} {
	return h[0]
}

// Push function pushes an element into the MaxHeap
func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(Set))
}

// Pop function pops the element at the top of the MaxHeap
func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type Point struct {
	x, y int
}

// InitPoint will be used to make a Point type object
func InitPoint(x, y int) Point {
	p := Point{}
	p.x = x
	p.y = y
	return p
}

// Less is used for max-heap
func (p Point) Less(other Point) bool {
	return p.DistFromOrigin() < other.DistFromOrigin()
}

// String is used to print the x and y values
func (p Point) String() string {
	return "[" + strconv.Itoa(p.x) + ", " + strconv.Itoa(p.y) + "]"
}

// DistFromOrigin calculates the distance using x, y coordinates
func (p Point) DistFromOrigin() int {
	// ignoring sqrt to calculate the distance
	return (p.x * p.x) + (p.y * p.y)
}

// Structure for MaxHeap
type MaxHeapV2 []Point

// newMaxHeap function intializes an instance of the MaxHeap
func newMaxHeapV2() *MaxHeapV2 {
	max := &MaxHeapV2{}
	heap.Init(max)
	return max
}

// Len function returns the length of the MaxHeap
func (h MaxHeapV2) Len() int {
	return len(h)
}

// Empty returns true if the MaxHeap is empty, false otherwise
func (h MaxHeapV2) Empty() bool {
	return len(h) == 0
}

// Greater returns true if the first of the given elements is greater than the second one
func (h MaxHeapV2) Less(i, j int) bool {
	return h[i].DistFromOrigin() > h[j].DistFromOrigin()
}

// Swap function swaps the values at the given indices
func (h MaxHeapV2) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// Top function returns the element at the top of the MaxHeap
func (h MaxHeapV2) Top() any {
	return h[0]
}

// Push function inserts the element in the MaxHeap
func (h *MaxHeapV2) Push(x any) {
	*h = append(*h, x.(Point))
}

// Pop function pops the element at the top of the MaxHeap
func (h *MaxHeapV2) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type Pair struct {
	first, second int
}

// MinHeap structure initialization
type MinHeap []Pair

// newMinHeap function initializes an instance of MinHeap
func newMinHeap() *MinHeap {
	min := &MinHeap{}
	heap.Init(min)
	return min
}

// Len function returns the length of min heap
func (h MinHeap) Len() int {
	return len(h)
}

// Empty function returns true if empty, false otherwise
func (h MinHeap) Empty() bool {
	return len(h) == 0
}

// Less function compares two elements of the MinHeap given their indexes
func (h MinHeap) Less(i, j int) bool {
	return (h[i].second < h[j].second)
}

// Swap function swaps the values of the elements whose indices are given
func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// Top function returns the element at the top of the MinHeap
func (h MinHeap) Top() Pair {
	return h[0]
}

// Push function pushes an element into the MinHeap
func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(Pair))
}

// Pop function pops the element at the top of the MinHeap
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
