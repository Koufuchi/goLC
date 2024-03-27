package houseRobber

/*
* 定 f(i) 為 num[0] ~ nums[i] 中的最大解：
* nums = [1,4,2,3]
*
  - 1,       f(0) = 1
  - 1 4,     f(1) = max(nums[1], f(0)) = max(4, 1) = 4
  - 1 4 2,   f(2) = max(nums[2]+f(0), f(1)) = max(2+1, 4) = 4
  - 1 4 2 3, f(3) = max(nums[3]+f(1), f(2)) = max(3+4, 4) = 7
    拆解：f(3) = max(nums[3]+ max(nums[1], f(0)), max(nums[2]+f(0), max(nums[1], f(0))))

* 得公式 f(n) = max(nums[n]+f(n-2), f(n-1))
*/
func rob(nums []int) int {
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

func robWithMem(nums []int) int {
	var mem = make(map[int]int, len(nums))

	for i, v := range nums {
		prevPrev, exist := mem[i-2]
		if !exist {
			prevPrev = 0
		}
		prev, exist := mem[i-1]
		if !exist {
			prev = 0
		}
		if prevPrev+v > prev {
			mem[i] = prevPrev + v
		} else {
			mem[i] = prev
		}
	}

	return mem[len(mem)-1]
}
