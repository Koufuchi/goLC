package houseRobber2

/*
 * 與 house robber 不一樣的地方在於頭尾也算相鄰。
 * 如果選頭就不能選尾，反之亦然，所以要跑兩次 house robber。
 */
func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	result := robSub(nums[1:])
	result2 := robSub(nums[:len(nums)-1])

	if result > result2 {
		return result
	}
	return result2
}

func robSub(nums []int) int {
	var max, robP, robPP int
	for _, v := range nums {
		if robPP+v > robP {
			max = robPP + v
		} else {
			max = robP
		}
		robPP = robP
		robP = max
	}
	return max
}
