package main

import (
	"container/heap"
	"fmt"
)

func main() {
	arr := []int{4, 7, 9}
	arr2 := []int{4, 7, 9}
	fmt.Println(kSmallestPairs(arr, arr2, 10))
}

func kthSmallestElement(matrix [][]int, k int) int {
	h := &MinHeap{}
	heap.Init(h)
	totalLen := 0
	for i, column := range matrix {
		heap.Push(h, Set{
			n1: column[0],
			n2: 0,
			n3: i,
		})
		totalLen += len(column)
	}

	if totalLen > k {
		totalLen = k
	}

	var elem Set
	for i := 0; i < totalLen; i++ {
		elem = heap.Pop(h).(Set)
		next := elem.n2 + 1
		if next >= len(matrix[elem.n3]) {
			continue
		}
		heap.Push(h, Set{
			n1: matrix[elem.n3][next],
			n2: next,
			n3: elem.n3,
		})
	}
	return elem.n1
}

func kSmallestPairs(list1 []int, list2 []int, k int) [][]int {
	totalLen := len(list1) * len(list2)
	result := make([][]int, 0, totalLen)

	h := &MinHeap{}
	heap.Init(h)
	heap.Push(h, Set{
		n1: list1[0] + list2[0], // value
		n2: list1[0],            // n1 index
		n3: list2[0],            // list
	})

	ix1, maxIndex1, ix2, maxIndex2 := 0, 0, 0, 0
	for i := 0; i < totalLen; i++ {

		next1 := list1[ix1]
		prev1, prevIx1 := next1, ix1
		if len(list1) > ix1+1 {
			ix1++
			next1 = list1[ix1]
		}

		next2 := list2[ix2]
		prev2, prevIx2 := next2, ix2
		if len(list2) > ix2+1 {
			ix2++
			next2 = list2[ix2]
		}

		if next1 < next2 && len(list1) > ix1+1 || list2[len(list2)-1] == prev2 {
			next2 = prev2
			ix2 = prevIx2
			if maxIndex1 < ix1 {
				ix2 = 0
				next2 = list2[0]
				maxIndex1 = ix1
			}
			heap.Push(h, Set{
				n1: next1 + next2,
				n2: next1,
				n3: next2,
			})
			continue
		}
		if next2 < next1 && len(list2) > ix2+1 || list1[len(list1)-1] == prev1 {
			next1 = prev1
			ix1 = prevIx1
			if maxIndex2 < ix2 {
				ix1 = 0
				next1 = list1[0]
				maxIndex2 = ix2
			}
			heap.Push(h, Set{
				n1: next1 + next2,
				n2: next1,
				n3: next2,
			})
			continue
		}

		switch {
		case maxIndex1 < ix1:
			ix2 = 0
			maxIndex1 = ix1
			next2 = list2[ix2]
		case maxIndex2 < ix2:
			ix1 = 0
			maxIndex2 = ix2
			next1 = list1[ix1]
		default:
		}
		heap.Push(h, Set{
			n1: next1 + next2,
			n2: next1,
			n3: next2,
		})
	}

	for i := 0; i < k && !h.Empty(); i++ {
		s := heap.Pop(h).(Set)
		result = append(result, []int{s.n2, s.n3})
	}
	return result
}

func (s *Set) updateSet(val, ix int) {
	s.n1 = val
	s.n2 = ix
}

func getElems(h *MinHeap) (Set, Set) {
	e := heap.Pop(h).(Set)
	if e.n3 == 1 {
		return e, heap.Pop(h).(Set)
	}
	return heap.Pop(h).(Set), e
}

func kSmallestNumber(lists [][]int, k int) int {
	h := &MinHeap{}
	heap.Init(h)

	totalLen := 0
	for i, list := range lists {
		if len(list) == 0 {
			continue
		}
		heap.Push(h, Set{
			n1: list[0],
			n2: 0,
			n3: i,
		})
		totalLen += len(list)
	}
	if totalLen == 0 {
		return 0
	}
	resultList := make([]int, 0, totalLen)
	for i := 0; i < totalLen; i++ {

		minSet := heap.Pop(h).(Set)
		resultList = append(resultList, minSet.n1)

		list := lists[minSet.n3]

		next := minSet.n2 + 1
		if len(list) <= next {
			continue
		}

		heap.Push(h, Set{
			n1: list[next],
			n2: next,
			n3: minSet.n3,
		})
	}

	if k > totalLen {
		return resultList[len(resultList)-1]
	}
	return resultList[k-1]
}
