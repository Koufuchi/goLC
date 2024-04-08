package longestConsecutiveSequence

func longestConsecutive(nums []int) int {
	var memo = make(map[int]interface{}, len(nums))
	var max int

	for _, v := range nums {
		memo[v] = struct{}{}
	}

	for _, v := range nums {
		var curr = v
		if _, exist := memo[curr-1]; exist {
			continue
		}

		var temp = 1
		for {
			if _, exist := memo[curr+1]; exist {
				temp++
				curr++
			} else {
				break
			}
		}
		if temp > max {
			max = temp
		}
	}
	return max
}
