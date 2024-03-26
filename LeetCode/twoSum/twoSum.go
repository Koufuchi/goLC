package twoSum

func twoSum(nums []int, target int) []int {
	if len(nums) < 2 {
		return []int{}
	}

	var vmap = make(map[int]int)
	for k, v := range nums {
		ans, ok := vmap[target-v]
		if ok {
			return []int{ans, k}
		}
		vmap[v] = k
	}

	return []int{}
}
