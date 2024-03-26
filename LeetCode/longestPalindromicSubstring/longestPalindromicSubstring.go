package longestpalindromicsubstring

import (
	"strings"
)

func longestPalindrome(s string) string {
	var (
		max      = 0
		ansLeft  = 0
		ansRight = 0
	)
	var findFromMid = func(left, right int) {
		for left >= 0 && right < len(s) && s[left] == s[right] {
			left--
			right++
		}
		left++
		right--
		if right-left > max {
			max = right - left
			ansLeft = left
			ansRight = right
		}
	}
	for k := range s {
		findFromMid(k, k)
		findFromMid(k, k+1)
	}

	return s[ansLeft : ansRight+1]
}

func longestPalindromeSpecial(s string) string {
	if len(s) == 1 {
		return s
	}
	var maxLen = 0
	var ansLeft, ansRight = 0, 0

	var newS = "#" + strings.Join(strings.Split(s, ""), "#") + "#"

	for k := range newS {
		var left = k - 1
		var right = k + 1
		for left >= 0 && right < len(newS) && newS[left] == newS[right] {
			left--
			right++
		}
		left++
		right--
		if right-left > maxLen {
			maxLen = right - left
			if left == 1 {
				ansLeft = 0
			} else {
				ansLeft = left / 2
			}
			ansRight = right / 2
		}
	}

	return s[ansLeft:ansRight]
}
