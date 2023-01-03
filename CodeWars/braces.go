package main

func ValidBraces(str string) bool {

	braces := make(map[byte][]int, 3)

	for i, sign := range str {
		if sign == '(' || sign == '[' || sign == '{' {
			braces[str[i]] = append(braces[str[i]], i+1)
		} else {
			if openBr, ok := braces[pair(str[i])]; ok && len(openBr) > 0 {
				maxCounter := openBr[len(openBr)-1]
				for _, signCounter := range braces {
					if len(signCounter) > 0 && maxCounter < signCounter[len(signCounter)-1] {

						return false

					}
				}
				braces[pair(str[i])] = removeLast(braces[pair(str[i])])
			} else {
				return false
			}
		}
	}

	trueCounter := 0
	for _, slice := range braces {
		if len(slice) == 0 {
			trueCounter++
		}
	}

	return trueCounter == len(braces)
}

func removeLast(s []int) []int {
	return s[:len(s)-1]
}

func pair(sign byte) byte {
	switch sign {
	case ')':
		return '('
	case ']':
		return '['
	case '}':
		return '{'
	default:
		return ' '
	}
}

/*
func main() {
	fmt.Println(ValidBraces("(){}[]"))
}



type braced map[string][]int

func ValidBraces(str string) bool {

	braces := make(braced, len(str))

	for i := 0; i < len(str); i++ {
		braces[string(str[i])] = append(braces[string(str[i])], i)
	}

	return braces.isBraced("(", ")") && braces.isBraced("[", "]") && braces.isBraced("{", "}") {
	}
}

func (b braced) isBraced(left, right string) bool {

	switch {
	case b[left] == nil && b[right] == nil:
		return true
	case b[left] != nil && b[right] != nil:
		return len(b[left]) == len(b[right]) && sum(b[left]) < sum(b[right])
	default:
		return false
	}
Println()
}

func sum(s []int) int {
	var sum int
	for _, v := range s {
		sum += v
	}
	return sum
}

func main() {
	fmt.Println(ValidBraces("[(])"))
}

*/
