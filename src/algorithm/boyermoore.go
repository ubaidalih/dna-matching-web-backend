package algorithm

//Mencari last occurence dari setiap karakter
func lastOccurence(s string) map[string]int {
	last := make(map[string]int)
	for i := 0; i < len(s); i++ {
		last[string(s[i])] = i
	}
	return last
}

//Boyer Moore Algorithm
//Return indeks pertama string ditemukan, jika tidak ada return -1
func BoyerMoore(s, t string) int {
	last := lastOccurence(t)
	i, j := len(t)-1, len(t)-1
	for i < len(s) && j < len(t) {
		if s[i] == t[j] {
			if j == 0 {
				return i
			} else {
				i--
				j--
			}
		} else {
			if last, ok := last[string(s[i])]; ok {
				if last+1 > j {
					i += len(t) - j
				} else {
					i += len(t) - last - 1
				}
				j = len(t) - 1
			} else {
				i += len(t)
				j = len(t) - 1
			}
		}
	}
	if j == 0 {
		return i
	}
	return -1
}
