package validAnagram

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sMap := make(map[rune]int, len(s))
	for _, rune := range s {
		sMap[rune]++
	}
	for _, rune := range t {
		if times, exist := sMap[rune]; exist && times > 0 {
			sMap[rune]--
		} else {
			return false
		}
	}
	for _, times := range sMap {
		if times > 0 {
			return false
		}
	}
	return true
}
