package kata

func FindOdd(seq []int) int { // 6 kyu
	// your code here
	nums := make(map[int]int)
	for i := 0; i < len(seq); i++ {
		if _, ok := nums[seq[i]]; ok {
			continue
		}
		nums[seq[i]]++
		for j := i + 1; j < len(seq); j++ {
			if seq[i] == seq[j] {
				nums[seq[i]]++
			}
		}
		if nums[seq[i]]%2 != 0 {
			return seq[i]
		}
	}
	return -1
}
