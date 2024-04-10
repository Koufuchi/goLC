package combinationSum

import (
	"sort"
)

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	results := make([][]int, 0)
	combinationSumSub(candidates, 0, target, []int{}, &results)

	return results
}

func combinationSumSub(candidates []int, index int, target int, memo []int, result *[][]int) {
	if target == 0 {
		valid := make([]int, len(memo))
		copy(valid, memo)
		*result = append(*result, valid)
	} else if target > 0 {
		for curr := index; curr < len(candidates); curr++ {
			combinationSumSub(candidates, curr, target-candidates[curr], append(memo, candidates[curr]), result)
		}
	}
}
