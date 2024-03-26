package longestsubstringwithoutrepeatingcharacters

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	if len(s) == 1 {
		return 1
	}

	var (
		left = 0
		max  = 0
		sMap = make(map[string]struct{})
	)

	for _, v := range s {
		_, exist := sMap[string(v)]
		if exist {
			if len(sMap) > max {
				max = len(sMap)
			}
			for string(s[left]) != string(v) {
				delete(sMap, string(s[left]))
				left++
			}
			left++
			continue
		}
		sMap[string(v)] = struct{}{}
	}

	if len(sMap) > max {
		return len(sMap)
	}

	return max
}
