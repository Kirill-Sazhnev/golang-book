package main

import "fmt"

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

func FormatDuration(seconds int64) string { // 4 kyu
	if seconds == 0 {
		return "now"
	}
	time := make(map[string]int64, 6)
	switch {
	case seconds >= 31536000:
		time["year"] = seconds / 31536000
		time["remndr"] = seconds % 31536000
		fallthrough
	case seconds >= 86400:
		calculate(seconds, 86400, "day", time)
		fallthrough
	case seconds >= 3600:
		calculate(seconds, 3600, "hour", time)
		fallthrough
	case seconds >= 60:
		calculate(seconds, 60, "minute", time)
		fallthrough
	default:
		calculate(seconds, 1, "second", time)
	}
	delete(time, "remndr")
	res := ""
	ln := len(time)
	for i := 0; i < ln; i++ {
		switch {
		case time["second"] > 0:
			res = fmt.Sprintf(" and %d second", time["second"])
			if time["second"] > 1 {
				res += "s"
			}
			delete(time, "second")
		case time["minute"] > 0:
			res = stringTime("minute", res, time)
		case time["hour"] > 0:
			res = stringTime("hour", res, time)
		case time["day"] > 0:
			res = stringTime("day", res, time)
		case time["year"] > 0:
			res = stringTime("year", res, time)
		}
	}
	if ln == 1 {
		return res[5:]
	}
	return res[2:]
}

func calculate(ttlSec, formSec int64, frmt string, time map[string]int64) {
	if time["remndr"] >= formSec {
		time[frmt] = time["remndr"] / formSec
		time["remndr"] = time["remndr"] % formSec
	} else if rem, ok := time["remndr"]; rem < formSec && ok {
		return
	} else {
		time[frmt] = ttlSec / formSec
		time["remndr"] = ttlSec % formSec
	}
}

func stringTime(frmt, res string, time map[string]int64) string {
	if len(res) == 0 {
		res = fmt.Sprintf(" and %d %s", time[frmt], frmt)
		if time[frmt] > 1 {
			res += "s"
		}
	} else {
		temp := fmt.Sprintf(", %d %s", time[frmt], frmt)
		if time[frmt] > 1 {
			temp += "s"
		}
		res = temp + res
	}
	delete(time, frmt)
	return res
}
