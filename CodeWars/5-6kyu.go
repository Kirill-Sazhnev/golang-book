package main

import (
	"fmt"
	"strconv"
	"strings"
)

func MoveZeros(arr []int) []int { //5kyu
	var intArr []int
	var zeroArr []int

	for _, v := range arr {
		switch v != 0 {
		case true:
			intArr = append(intArr, v)
		case false:
			zeroArr = append(zeroArr, v)
		}
	}
	return append(intArr, zeroArr...)
}

func SpinWords(str string) string { //6kyu
	words := strings.Fields(str)
	sdrow := make([]string, len(words))
	for i, word := range words {
		if len(word) > 4 {
			drow := make([]byte, len(word))
			for j, letter := range word {
				drow[len(word)-j-1] = byte(letter)
			}
			sdrow[i] = string(drow)
		} else {
			sdrow[i] = word
		}
	}
	return strings.Join(sdrow, " ")
} // SpinWords

func ArrayDiff(a, b []int) []int { //6kyu

	for _, valB := range b {
		difA := make([]int, 0)
		for _, valA := range a {
			if valB != valA {
				difA = append(difA, valA)
			}
		}
		a = difA
	}
	return a
}

func FindNb(m int) int { //6kyu
	currNb, cubes := 0, 0

	for i := 1; currNb < m; i++ {
		currNb += i * i * i
		cubes++
	}

	if currNb == m {
		return cubes
	}
	return -1
}

func Revrot(s string, n int) string { // 6kyu
	if n < 1 {
		return ""
	}
	subStrgs := make([]string, 0)

	for i := 0; i <= len(s)-n; i += n {
		subStrgs = append(subStrgs, s[i:i+n])
	}

	for i, subStr := range subStrgs {
		if cubeSum(subStr)%2 == 0 {
			subStrgs[i] = reverse(subStr)
		} else {
			bytes := []byte(subStrgs[i])
			bytes = append(bytes[1:], bytes[0])
			subStrgs[i] = string(bytes)
		}
	}
	return strings.Join(subStrgs, "")
}

func cube(n int) int {
	return n * n * n
}

func cubeSum(nums string) int {
	sum := 0
	for _, val := range nums {
		num, _ := strconv.Atoi(string(val))
		sum += cube(num)
	}
	return sum
}

func reverse(s string) string {
	ln := len(s)
	str := []byte(s)
	for i := 0; i < ln/2; i++ {
		str[i], str[ln-i-1] = str[ln-i-1], str[i]
	}
	return string(str)
}

func Millipede(words []string) bool { // 6kyu
	for i, word := range words {

		wordSl := reSlice(i, words)

		if MPrecur(word, wordSl) {
			return true
		}
	}
	return false
}

func MPrecur(word string, words []string) bool {
	if len(words) == 1 && llo(word) == flo(words[0]) {
		return true
	}

	for i, current := range words {

		if llo(word) == flo(current) {
			wordSl := reSlice(i, words)

			if MPrecur(current, wordSl) {
				return true
			}
		}
	}
	return false
}

func reSlice(i int, words []string) []string {
	wordSl := make([]string, len(words))
	copy(wordSl, words)
	return append(wordSl[:i], wordSl[i+1:]...)
}

func flo(s string) byte {
	return s[0]
}

func llo(s string) byte {
	return s[len(s)-1]
}

func DirReduc(arr []string) []string { // 5kyu
	for i := 0; i < len(arr)-1; i++ {
		if isOpposite(arr[i], arr[i+1]) {
			arr = Slice(arr, i)
			i = -1
		}
	}
	return arr
}

func isOpposite(dir1, dir2 string) bool {
	switch dir1 {
	case "NORTH":
		return dir2 == "SOUTH"
	case "SOUTH":
		return dir2 == "NORTH"
	case "WEST":
		return dir2 == "EAST"
	case "EAST":
		return dir2 == "WEST"
	default:
		panic("Unknown direction")
	}
}

func Slice(slice []string, i int) []string {
	return append(slice[:i], slice[i+2:]...)
}

func PrimeFactors(n int) string { // 5kyu

	primes := make([]int, 0)
	for i := 2; n >= i; i++ {
		if n%i == 0 {
			primes = append(primes, i)
			n = n / i
			i = 1
		}
	}

	result := ""
	powerCntr := 0
	for i := 0; i < len(primes); i++ {
		switch {
		case i > 0 && primes[i-1] == primes[i] && powerCntr == 0:
			result += "**"
			powerCntr = 2
		case powerCntr > 0 && primes[i-1] == primes[i]:
			powerCntr++
		case powerCntr > 0:
			result += strconv.Itoa(powerCntr)
			powerCntr = 0
			fallthrough
		case i > 0 && powerCntr == 0:
			result += ")"
			fallthrough
		default:
			result += "(" + strconv.Itoa(primes[i])
		}
	}
	if powerCntr > 0 {
		result += strconv.Itoa(powerCntr) + ")"
	} else {
		result += ")"
	}
	return result
}

func Tribonacci(signature [3]float64, n int) []float64 { // 6 kyu
	return TribonacciRecur(signature[:], n)
}

func TribonacciRecur(slice []float64, n int) []float64 {
	if n <= 3 {
		initialSl := make([]float64, n)
		for i := 0; i < n; i++ {
			initialSl[i] = slice[i]
		}
		return initialSl
	}

	newSl := TribonacciRecur(slice, n-1)
	ln := len(newSl)
	return append(newSl, newSl[ln-1]+newSl[ln-2]+newSl[ln-3])
}

func Crossover(ns []int, xs []int, ys []int) ([]int, []int) { // 6 kyu
	// Your code here
	isCross := 1
	newXs, newYs := make([]int, len(xs)), make([]int, len(ys))
	for i := range xs {
		if crossOver(ns, i, &isCross) {
			newXs[i], newYs[i] = ys[i], xs[i]
		} else {
			newXs[i], newYs[i] = xs[i], ys[i]
		}
	}
	return newXs, newYs
}

func crossOver(ns []int, ix int, isCross *int) bool {
	for _, val := range ns {
		if val == ix {
			*isCross += 1
			break
		}
	}
	return *isCross%2 == 0
}

func MultiplicationTable(size int) [][]int { // 6 kyu
	table := make([][]int, size)
	for i := 1; i <= size; i++ {
		newRow := make([]int, size)
		for j := 1; j <= size; j++ {
			newRow[j-1] = i * j
		}
		table[i-1] = newRow
	}
	return table
}

func HumanReadableTime(seconds int) string { // 5 kyu
	hours := seconds / 3600
	mins := seconds % 3600 / 60
	secs := (seconds % 3600) % 60
	return fmt.Sprintf("%02d:%02d:%02d", hours, mins, secs)
}

func main() {
	fmt.Println(FindNb(100))
}

// ([]int{1, 2, 0, 1, 0, 1, 0, 3, 0, 1}) // returns []int{ 1, 2, 1, 1, 3, 1, 0, 0, 0, 0 }
