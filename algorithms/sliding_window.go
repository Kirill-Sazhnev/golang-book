package main

func longestRepeatingCharacterReplacement(s string, k int) int {
	maxLen := 0
	for i, ch := range s {
		steps := k
		prevChar := ch
		length := 0
		subString := s[i:]
		for _, nextCh := range subString {
			if prevChar == nextCh {
				length++
				continue
			}
			if steps == 0 {
				break
			}
			steps--
			length++
			prevChar = nextCh
		}

		if length > maxLen {
			maxLen = length
		}
		if maxLen == len(s) {
			break
		}
	}
	return maxLen
}

type chRate struct {
	chars map[uint8]int
	max   int
	maxCh uint8
}

func newChRate() *chRate {
	return &chRate{
		chars: make(map[uint8]int),
	}
}

func longestRepeatingCharacterReplacementV2(s string, k int) int {
	charRate := newChRate()
	right, left := 0, 0
	maxLen := 0

	for right < len(s) {
		ch := s[right]
		maxRate := charRate.increase(ch)

		for (right-left+1)-maxRate > k {
			maxRate = charRate.decrease(s[left])
			left++
			if (right-left+1)-maxRate <= k {
				break
			}
		}
		if (right - left + 1) > maxLen {
			maxLen = right - left + 1
		}
		right++
	}

	return maxLen
}

func (cr *chRate) decrease(char uint8) int {
	cr.chars[char]--
	if cr.maxCh != char {
		return cr.max
	}
	cr.max = 0
	cr.maxCh = 0
	for c, v := range cr.chars {
		if v > cr.max {
			cr.maxCh = c
			cr.max = v
		}
	}
	return cr.max
}

func (cr *chRate) increase(char uint8) int {
	cr.chars[char]++
	if cr.chars[char] > cr.max {
		cr.maxCh = char
		cr.max = cr.chars[char]
	}
	return cr.max
}

func findLongestSubstring(str string) int {
	chars := make(map[rune]int)
	maxLen := 0
	start := 0
	for end, char := range str {
		if start == end {
			chars[rune(str[start])]++
			continue
		}

		rate := chars[char]
		if rate > 0 {
			chars[rune(str[start])]--
			start++
			continue
		}

		currLen := end - start + 1
		if currLen > maxLen {
			maxLen = currLen
		}
		chars[char]++
	}

	return maxLen
}

func findLongestSubstringV2(str string) int {
	chars := make(map[rune]int)
	maxLen := 0
	start := 0
	chars[rune(str[start])] = 1
	for end, char := range str {
		if start == end {
			currLen := end - start + 1
			if currLen > maxLen {
				maxLen = currLen
			}
			continue
		}

		place := chars[char]
		if place > start {
			start = place
		}

		currLen := end - start + 1
		if currLen > maxLen {
			maxLen = currLen
		}
		chars[char] = end + 1
	}

	return maxLen
}
