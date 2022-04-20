package algorithm

func hammingDistance(s, t string) int {
	count := 0
	max := 0
	for i := 0; i+len(t) <= len(s); i++ {
		count = 0
		for j := 0; j < len(t); j++ {
			if s[i+j] == t[j] {
				count++
			}
		}
		if count > max {
			max = count
		}
	}
	return 100 * max / len(t)
}
