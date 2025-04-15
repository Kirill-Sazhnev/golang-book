package main

import (
	"container/heap"
	"fmt"
	"sort"
)

func main() {
	//arr := [][]int{{1,7},{8,13},{5,6},{10,14},{6,7}}
	//arr := [][]int{{0, 5}, {1, 6}, {6, 7}, {7, 8}, {8, 9}}
	//arr := [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}, {4, 5}}
	arr := [][]int{{0, 4}, {1, 3}, {2, 4}, {3, 5}, {4, 6}, {5, 7}}
	fmt.Println(mostBooked(arr, 4))
}

func mostBooked(meetings [][]int, rooms int) int {
	if len(meetings) == 0 {
		return 0
	}
	h := &PairHeap{}
	hr := &MinHeap{}
	heap.Init(h)
	heap.Init(hr)
	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i][0] < meetings[j][0]
	})
	for i := 0; i < rooms; i++ {
		heap.Push(hr, i)
	}

	for _, m := range meetings {
		start, end := m[0], m[1]

		if h.Empty() {
			roomNum := heap.Pop(hr).(int)
			heap.Push(h, Pair{
				endTime: int64(end),
				room:    roomNum,
				counter: 1,
			})
			continue
		}

		nearestMeeting := heap.Pop(h).(Pair)
		for nearestMeeting.endTime < int64(start) && !h.Empty() {
			heap.Push(hr, nearestMeeting.room)
			nearestMeeting = heap.Pop(h).(Pair)
		}
		nearestEnd := int(nearestMeeting.endTime)

		if hr.Len() > 0 && hr.Top().(int) < nearestMeeting.room {
			heap.Push(h, nearestMeeting)
			roomNum := heap.Pop(hr).(int)
			heap.Push(h, Pair{
				endTime: int64(end),
				room:    roomNum,
				counter: 1,
			})
			continue
		}
		if nearestEnd == start {
			nearestMeeting.endTime = int64(end)
			nearestMeeting.counter++
			heap.Push(h, nearestMeeting)
			continue
		}
		if hr.Len() > 0 {
			heap.Push(h, nearestMeeting)
			roomNum := heap.Pop(hr).(int)
			heap.Push(h, Pair{
				endTime: int64(end),
				room:    roomNum,
				counter: 1,
			})
			continue
		}

		delta := nearestEnd - start
		nearestMeeting.endTime = int64(end + delta)
		nearestMeeting.counter++
		heap.Push(h, nearestMeeting)
	}

	resultRoom, maxCntr := -1, 0
	for !h.Empty() {
		room := heap.Pop(h).(Pair)
		if maxCntr < room.counter {
			maxCntr = room.counter
			resultRoom = room.room
		}
		if maxCntr == room.counter {
			resultRoom = minRoom(room.room, resultRoom)
		}
	}
	return resultRoom
}

func minRoom(i, j int) int {
	if j < 0 {
		return i
	}
	if i < j {
		return i
	}
	return j
}

// MinHeap structure initialization
type MinHeap []int

func (h MinHeap) Len() int {
	return len(h)
}

// Top function returns the element at the top of the heap
func (h MinHeap) Top() interface{} {
	if len(h) == 0 {
		return -1
	}
	return h[0]
}

func (h MinHeap) Empty() bool {
	return len(h) == 0
}

// Less function compares two elements of MinHeap given their indices
func (h MinHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

// Swap function swaps the value of the elements whose indices are given
func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// Push function pushes an element into the MinHeap
func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

// Pop function pops the element at the top of the MinHeap
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// Pair struct to hold the end time and room number
type Pair struct {
	endTime int64
	room    int
	counter int
}

// PairHeap is a MinHeap of Pairs
type PairHeap []Pair

func (h PairHeap) Len() int {
	return len(h)
}
func (h PairHeap) Empty() bool {
	return len(h) == 0
}

func (h PairHeap) Less(i, j int) bool {
	return h[i].endTime < h[j].endTime || (h[i].endTime == h[j].endTime && h[i].room < h[j].room)
}
func (h PairHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *PairHeap) Push(x interface{}) {
	*h = append(*h, x.(Pair))
}

func (h *PairHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
