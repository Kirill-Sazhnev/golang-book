package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(restoreIpAddresses(
		"201023"))
}

func restoreIpAddresses(s string) []string {
	res := make([]string, 0, len(s))
	for i := 1; i <= 3; i++ {
		res = buildAddress(s, "", 0, i, 0, res)
	}
	return res
}

func buildAddress(s, adr string, ix, subIx, depth int, res []string) []string {
	if len(adr)-3 > len(s) || ix+subIx > len(s) || depth > 3 {
		return res
	}
	if subIx != 1 && s[ix] == '0' {
		return res
	}
	if ix != 0 {
		adr += "."
	}
	newSub := s[ix : ix+subIx]
	number, _ := strconv.Atoi(newSub)
	if number > 255 {
		return res
	}
	adr += newSub
	if len(adr)-3 == len(s) {
		return append(res, adr)
	}
	for i := 1; i <= 3; i++ {
		res = buildAddress(s, adr, ix+subIx, i, depth+1, res)
	}
	return res
}

func restoreIpAddressesV2(s string) []string {
	res := make([]string, 0)
	possibleDepths := make([][]int, 0, len(s))
	for i := 0; i < 81; i++ {
		d := depths(i, len(s))
		if d == nil {
			continue
		}
		possibleDepths = append(possibleDepths, d)
	}
	for _, depth := range possibleDepths {
		address := newAddress(s, depth)
		if address != "" {
			res = append(res, address)
		}
	}
	return res
}

func newAddress(s string, depths []int) string {
	res := ""
	charIx := 0
	for i, num := range depths {
		if i > 0 {
			res += "."
		}
		segment := ""
		for j := 0; j < num; j++ {
			char := s[charIx+j]

			switch {
			case j == 0 && j+1 < num && char == '0':
				return ""
			default:
				segment += string(char)
			}
		}
		number, _ := strconv.Atoi(segment)
		if len(segment) == 0 || number > 255 {
			return ""
		}
		res += segment
		charIx += num
	}
	return res
}

func depths(ix int, len int) []int {
	res := []int{1, 1, 1, 1}

	for i := 3; i >= 0; i-- {
		newDepth := ix % 3
		if newDepth > 0 {
			res[i] += newDepth
			ix -= newDepth
		}
		ix /= 3
	}
	totalLen := res[0] + res[1] + res[2] + res[3]
	if totalLen != len {
		return nil
	}
	return res
}

type TreeNode[T any] struct {
	Data  T
	Left  *TreeNode[T]
	Right *TreeNode[T]
}

func rob(root *TreeNode[int]) int {
	return maxRobbed(robNode(root))
}

func robNode(root *TreeNode[int]) (noRoot int, withRoot int) {
	if root == nil {
		return 0, 0
	}
	leftNoRoot, leftWithRoot := robNode(root.Left)                                       // 34, 27
	rightNoRoot, rightWithRoot := robNode(root.Right)                                    // 50, 25
	withRoot = root.Data + leftNoRoot + rightNoRoot                                      // 15 + 34 + 50 = 99
	noRoot = maxRobbed(leftNoRoot, leftWithRoot) + maxRobbed(rightNoRoot, rightWithRoot) // 10 + 24 = 34

	return noRoot, withRoot
}

func maxRobbed(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func wordSearch(grid [][]string, word string) bool {
	charMap := resetMap(word)

	return explore(grid, 0, 0, charMap, word)
}

func explore(grid [][]string, x, y int, chMap map[string]int, word string) bool {
	if len(chMap) == 0 {
		return true
	}
	if x < 0 || y < 0 || x >= len(grid) || y >= len(grid[x]) {
		return false
	}

	char := grid[x][y]

	_, ok := chMap[char]
	if ok {
		deleteChar(chMap, char)
	} else {
		chMap = resetMap(word)
	}
	if explore(grid, x+1, y, chMap, word) {
		return true
	}
	if explore(grid, x, y+1, chMap, word) {
		return true
	}
	if ok {
		chMap[char]++
	}
	return false
}

func resetMap(word string) map[string]int {
	charMap := make(map[string]int)
	for _, chr := range word {
		charMap[(string(chr))]++
	}
	return charMap
}

func deleteChar(chMap map[string]int, char string) {
	count, ok := chMap[char]
	if ok && count > 1 {
		chMap[char]--
	} else {
		delete(chMap, char)
	}
}
