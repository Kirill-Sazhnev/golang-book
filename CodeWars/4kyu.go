package main

func Snail(snaipMap [][]int) (snail []int) { // 4 kyu
	if len(snaipMap) == 0 {
		return snail
	}

	snail = append(snail, snaipMap[0]...)
	snaipMap = snaipMap[1:]
	for i, val := range snaipMap {
		snail = append(snail, val[len(val)-1])
		snaipMap[i] = snaipMap[i][:len(val)-1]
	}

	if len(snaipMap) == 0 {
		return snail
	}

	lastSl := snaipMap[len(snaipMap)-1]
	for i := len(lastSl) - 1; i >= 0; i-- {
		snail = append(snail, lastSl[i])
	}
	snaipMap = snaipMap[:len(snaipMap)-1]
	for i := len(snaipMap) - 1; i >= 0; i-- {
		snail = append(snail, snaipMap[i][0])
		snaipMap[i] = snaipMap[i][1:]
	}
	snail = append(snail, Snail(snaipMap)...)
	return snail
}

func Solution1(sl []int) int { // 4 kyu

	ar := make([]int, len(sl))
	copy(ar, sl)
	cntr := 0
	i, j := len(ar)-2, len(ar)-1

	for cntr != len(ar)-1 {
		switch {
		case ar[i] > ar[j]:
			if ar[i]%ar[j] == 0 {
				ar[i] = ar[j]
			} else {
				ar[i] %= ar[j]
			}
			cntr = 0
		case ar[j] > ar[i]:
			if ar[j]%ar[i] == 0 {
				ar[j] = ar[i]
			} else {
				ar[j] %= ar[i]
			}
			cntr = 0
		default:
			cntr++
			if i == 0 {
				i = len(ar) - 1
			}
			i--
			j = i + 1
		}
	}
	return ar[0] * len(ar)
}
