package main

import (
	"container/heap"
	"fmt"
)

func main() {
	//str := "bbqnowgiznovswjffpbfvnnppnzblhnzxdgzdfgbulrjbrmhuvhwtvhlztqwgpyjjokpawuosinrdedntkycihcpybvkqtqnnmgsgbooenqeufwkynuxgkmmxuqmuvvxyxlbckdkmkhqdpjrxogywjjbjuigbkjhdmpurftagpbaeztgsdsljjgluppjzhgfthmmimngqdvejorrowmdoecqevfpoopfveqceodmworrojevdqgnmimmhtfghzjppulgjjlsdsgtzeabpgatfrupmdhjkbgiujbjjwygoxrjpdqhkmkdkcblxyxvvumquxmmkgxunykwfueqneoobgsgmnnqtqkvbypchicyktndedrnisouwapkojjypgwqtzlhvtwhvuhmrbjrlubgfdzgdxznhlbznppnnvfbpffjwsvonzigwonqbb"
	fmt.Println(reorganizeString("g"))
}

func topKFrequent(nums []int, k int) []int {
	numMap := make(map[int]int)
	for _, ch := range nums {
		numMap[ch]++
	}
	h := newMinHeap()
	steps := 0
	for num, counter := range numMap {
		steps++
		if steps <= k {
			heap.Push(h, Pair{num, counter})
			continue
		}
		minNum := h.Top()
		if counter > minNum.second {
			heap.Pop(h)
			heap.Push(h, Pair{num, counter})
		}
	}
	result := make([]int, 0, k)
	for !h.Empty() {
		num := heap.Pop(h).(Pair)
		result = append(result, num.first)
	}
	return result
}

func kClosest(points []Point, k int) []Point {
	h := newMaxHeapV2()
	for i := 0; i < k; i++ {
		heap.Push(h, points[i])
	}

	for i := k; i < len(points); i++ {
		maxPoint := h.Top().(Point)
		currentPoint := points[i]
		if currentPoint.DistFromOrigin() < maxPoint.DistFromOrigin() {
			heap.Pop(h)
			heap.Push(h, currentPoint)
		}
	}
	res := make([]Point, 0, k)
	for !h.Empty() {
		point := heap.Pop(h).(Point)
		res = append(res, point)
	}

	return res
}

func reorganizeString(str string) string {
	if len(str) < 2 {
		return str
	}
	charMap := make(map[rune]int)
	for _, ch := range str {
		charMap[ch]++
	}

	h := newMaxHeap()
	for char, count := range charMap {
		heap.Push(h, Set{
			char:  char,
			count: count,
		})
	}
	if h.Len() < 1 {
		return ""
	}

	var lastSet Set
	var res []rune
	for h.Len() > 1 {
		set1 := heap.Pop(h).(Set)
		set2 := heap.Pop(h).(Set)

		char1 := set1.char
		set1.count--
		char2 := set2.char
		set2.count--

		res = append(res, char1, char2)

		if set1.count > 0 {
			heap.Push(h, set1)
		}
		if set2.count > 0 {
			heap.Push(h, set2)
		}
	}
	if h.Empty() {
		return string(res)
	}

	lastSet = heap.Pop(h).(Set)

	if lastSet.count > 1 {
		return ""
	}

	lastChar := res[len(res)-1]
	if lastChar == lastSet.char {
		return ""
	}
	return string(append(res, lastSet.char))
}
