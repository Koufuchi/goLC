package containerwithmostwater

func maxArea(height []int) int {
	var max = 0
	var left = 0
	var right = len(height) - 1

	for left < right {
		currArea := 0
		if height[left] > height[right] {
			currArea = (right - left) * height[right]
			right--
		} else {
			currArea = (right - left) * height[left]
			left++
		}
		if currArea > max {
			max = currArea
		}
	}

	return max
}
