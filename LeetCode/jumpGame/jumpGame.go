package jumpGame

// 每一輪跑最大jump數
func canJump(nums []int) bool {
	var max int
	for k, v := range nums {
		if v > max {
			max = v
		}
		if max == 0 && k != len(nums)-1 {
			return false
		}
		max--
	}

	return true
}
