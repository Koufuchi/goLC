package threesum

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var resultMap = [][]int{}

	for first := 0; first < len(nums)-2; first++ {
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		start := first + 1
		end := len(nums) - 1

		for start < end {
			cnt := nums[first] + nums[start] + nums[end]
			if cnt == 0 {
				resultMap = append(resultMap, []int{nums[first], nums[start], nums[end]})
				for start < end && nums[start] == nums[start+1] {
					start++
				}
				for start < end && nums[end] == nums[end-1] {
					end--
				}
				start++
				end--
			} else if cnt > 0 {
				end--
			} else {
				start++
			}
		}
	}

	return resultMap
}

// sort : [-4, -1, -1, 0, 1, 2]
