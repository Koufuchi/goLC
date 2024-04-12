package courseSchedule

import "fmt"

/*
 * numCourses:[[0,1],[0,2],[1,3],[1,4],[3,4]], prerequisites:5
 *  0 -> 1 -> 3
 *  |    |    |
 *  v    |    v
 *  2     --> 4
 *
 *  每堂課對應其對誰來說是必修課
 *  preCourseMap: map[0:[1,2], 1:[3,4], 3:[4]]
 *
 *  對每堂課往後接續的課程做 DFS 確認是否有環，並將確認沒問題的節點洗掉避免重複。
 */
func canFinish(numCourses int, prerequisites [][]int) bool {
	preCourseMap := make(map[int][]int, numCourses)
	visit := make(map[int]struct{})
	for _, dependency := range prerequisites {
		preCourseMap[dependency[0]] = append(preCourseMap[dependency[0]], dependency[1])
	}
	fmt.Println("preCourseMap", preCourseMap)

	var dfs func(course int) bool
	dfs = func(course int) bool {
		// 曾找過代表存在環
		if _, exist := visit[course]; exist {
			return false
		}
		// 已達終端節點
		if len(preCourseMap[course]) == 0 {
			fmt.Println("已到終點")
			return true
		}
		// 紀錄當前已查找課程
		visit[course] = struct{}{}
		fmt.Println("當前訪問", course, "已紀錄 visit", visit)
		// 對當前課程往後接續的所有課程做遞迴
		for _, preCourse := range preCourseMap[course] {
			fmt.Println("呼叫下層遞迴 dfs(", preCourse, ") 尚有", preCourseMap)
			if !dfs(preCourse) {
				return false
			}
		}
		// 確認往後接都正常，刪除查找紀錄
		delete(visit, course)
		fmt.Println("刪除", course, "剩餘", visit)
		// 刪除接續課程紀錄，之後造訪到這裡就是終端節點
		preCourseMap[course] = []int{}
		fmt.Println("當前課程", course, "變為終端節點，剩餘", preCourseMap)
		return true
	}

	for idx := 0; idx < numCourses; idx++ {
		if !dfs(idx) {
			return false
		}
	}
	return true
}
