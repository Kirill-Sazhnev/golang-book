package main

func main() {
	arr := [][]string{{"E"}, {"C"}, {"E"}, {"G"}, {"D"}}
	word := "ECEGD"
	if wordSearch(arr, word) {
		println("Found")
	} else {
		println("Not Found")
	}
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
