package algorithm

//Border untuk KMP
func border(s string) []int {
	n := len(s) - 1
	next := make([]int, n)
	next[0] = -1
	j := -1
	for i := 1; i < n; i++ {
		for j != -1 && s[i] != s[j+1] {
			j = next[j]
		}
		if s[i] == s[j+1] {
			j++
		}
		next[i] = j
	}
	return next
}

//KMP Algorithm
//Return indeks pertama string ditemukan, jika tidak ada return -1
func KMP(s, t string) int {
	next := border(t)
	i, j := 0, 0
	for i < len(s) && j < len(t) {
		if s[i] == t[j] {
			i++
			j++
		} else {
			if j == 0 {
				i++
			} else {
				j = next[j-1] + 1
			}
		}
	}
	if j == len(t) {
		return i - j
	}
	return -1
}
