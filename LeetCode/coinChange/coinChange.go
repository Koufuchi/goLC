package coinChange

import (
	"math"
	"sort"
)

/*  對於 f(n) 來說，選擇面額 A 需要多取 f(n-A) 次。
 *               1        5       11
 *  f(0) = MIN(  X     ,  X,       X      ) = -1
 *  f(1) = MIN(  1     ,  X,       X      ) =  1
 *  f(2) = MIN(  1+f(1),  X,       X      ) =  2
 *  f(3) = MIN(  1+f(2),  X,       X      ) =  3
 *  f(4) = MIN(  1+f(3),  X,       X      ) =  4
 *  f(5) = MIN(  1+f(4),  1,       X      ) =  1
 *  f(6) = MIN(  1+f(5),  1+f(1),  X      ) =  2
 *  f(7) = MIN(  1+f(6),  1+f(2),  X      ) =  3
 *   ...
 *  f(15) = MIN( 1+f(14), 1+f(10), 1+f(4) ) =  3
 */
func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}

	var memo = make(map[int]int, amount+1)

	for i := 1; i <= amount; i++ {
		memo[i] = math.MaxInt32
		for _, v := range coins {
			temp := i - v
			if temp >= 0 && memo[i] > memo[temp]+1 {
				memo[i] = memo[temp] + 1
			}
		}
	}
	if memo[amount] == math.MaxInt32 {
		return -1
	}
	return memo[amount]
}

/*
coins=[1, 5, 11], amount=15, ans=3
無法靠 greedy 從最大面額窮盡來得到最佳解，所以純 DFS 會超時。
*/
func coinChangeDFS(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	sort.Ints(coins)
	min := coinChangeDFSSub(coins, len(coins)-1, amount, 0, math.MaxInt)
	if min != math.MaxInt {
		return min
	}
	return -1
}

func coinChangeDFSSub(coins []int, index, remain, used, min int) int {
	if remain == 0 {
		return used
	}
	if remain < 0 || index < 0 {
		return -1
	}
	for i := index; i >= 0; i-- {
		if remain >= coins[i] {
			ans := coinChangeDFSSub(coins, i, remain-coins[i], used+1, min)
			if ans < min {
				min = ans
			}
		}
	}

	return min
}
