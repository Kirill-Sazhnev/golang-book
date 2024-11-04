package main

import (
	"fmt"
	"math"
	"regexp"
	"sort"
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

func HumanReadableTime(seconds int) string { // 5 kyu
	hours := seconds / 3600
	mins := seconds % 3600 / 60
	secs := (seconds % 3600) % 60
	return fmt.Sprintf("%02d:%02d:%02d", hours, mins, secs)
}

func alphanumeric1(str string) bool { // 5 kyu
	valid := regexp.MustCompile(`^[a-zA-Z0-9]+$`)
	return valid.MatchString(str)
}

func RectIntoRects(l, w int) []string { // 5 kyu
	if l < w {
		l, w = w, l
	}

	sl := make([]int, 0)
	strSl := make([]string, 0)

	for i, j := l, w; i > 0 && j > 0; {
		i -= j
		sl = append(sl, j)
		if i < j {
			i, j = j, i
		}
	}

	for i := 0; i < len(sl)-1; i++ {
		for j := i + 1; j < len(sl); j++ {
			if sl[i] == sl[j] {
				strSl = append(strSl, fmt.Sprintf("(%v*%v)", sl[i]*(j-i+1), sl[i]))
			} else {
				strSl = append(strSl, fmt.Sprintf("(%v*%v)", sl[i]*(j-i)+sl[j], sl[i]))
				break
			}
		}
	}
	return strSl
}

func FindMissingNumber(seq []int) int { // 5 kyu
	sort.Ints(seq)
	temp, step := 0, 0
	for i := 1; i < len(seq); i++ {
		step = seq[i] - seq[i-1]
		if temp == step {
			break
		} else {
			temp = step
		}
	}

	for i := 1; i < len(seq); i++ {
		if seq[i-1]+step != seq[i] {
			return seq[i-1] + step
		}
	}
	return 1
}

func RGB(r, g, b int) string {
	r = int(math.Min(math.Max(float64(r), 0), 255))
	red := strconv.FormatInt(int64(r), 16)
	if len(red) == 1 {
		red = "0" + red
	}
	g = int(math.Min(math.Max(float64(g), 0), 255))
	green := strconv.FormatInt(int64(g), 16)
	if len(green) == 1 {
		green = "0" + green
	}
	b = int(math.Min(math.Max(float64(b), 0), 255))
	blue := strconv.FormatInt(int64(b), 16)
	if len(blue) == 1 {
		blue = "0" + blue
	}
	return strings.ToUpper(red + green + blue)
}
