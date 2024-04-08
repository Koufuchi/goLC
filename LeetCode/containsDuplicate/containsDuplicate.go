package containsDuplicate

func containsDuplicate(nums []int) bool {
	var visited = make(map[int]interface{}, len(nums))
	for _, v := range nums {
		if _, exist := visited[v]; exist {
			return true
		} else {
			visited[v] = struct{}{}
		}
	}
	return false
}
