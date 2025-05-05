package main

import "fmt"

func main() {
	word := "73"
	fmt.Println(letterCombinations(word))
}

func findAllSubsets(nums []int) []Set {

	res := make([]Set, 0, len(nums))
	res = append(res, *NewSet())

	for _, num := range nums {
		newSets := make([]Set, 0, len(res))

		for _, set := range res {
			s := NewSet()
			copyMap(s.hashMap, set.hashMap)
			s.Add(num)
			newSets = append(newSets, *s)
		}
		res = append(res, newSets...)
	}
	return res
}

func copyMap(dst map[int]bool, src map[int]bool) {
	for k, v := range src {
		dst[k] = v
	}
}

func permuteWordV2(word string) []string {

	res := make([]string, 0)

	for ix, letter := range word {
		newWords := make([]string, 0, len(word))
		if ix == 0 {
			newWords = append(newWords, string(letter))
		}
		for _, oldWord := range res {

			for i := 0; i <= len(oldWord); i++ {
				newWords = append(newWords, oldWord[:i]+string(letter)+oldWord[i:])
			}
		}
		res = newWords
	}
	return res
}

func permuteWord(word string) []string {

	res := make([]string, 0)
	res = append(res, word)

	for i := range word {
		newWords := make([]string, 0, len(word))
		for _, oldWord := range res {
			for j := i + 1; j < len(word); j++ {
				newWord := swap(i, j, oldWord)
				newWords = append(newWords, newWord)
			}
		}
		res = append(res, newWords...)
	}
	return res
}

func swap(a, b int, s string) string {
	res := []rune(s)
	res[a], res[b] = res[b], res[a]
	return string(res)
}

var letterMap = map[rune]string{
	'2': "abc",
	'3': "def",
	'4': "ghi",
	'5': "jkl",
	'6': "mno",
	'7': "pqrs",
	'8': "tuv",
	'9': "wxyz",
}

func letterCombinations(digits string) []string {
	res := make([]string, 0, len(digits))
	res = append(res, "")
	for _, digit := range digits {
		newComb := make([]string, 0, len(digits))

		letters := letterMap[digit]
		for _, letter := range letters {
			for _, comb := range res {
				comb = comb + string(letter)
				newComb = append(newComb, comb)
			}
		}
		res = newComb
	}
	return res
}
