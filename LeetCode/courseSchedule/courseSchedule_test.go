package courseSchedule

import "testing"

func Test_canFinish(t *testing.T) {
	type args struct {
		numCourses    int
		prerequisites [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "numCourses = 2, prerequisites = \"[[1,0]]\"",
			args: args{numCourses: 2, prerequisites: [][]int{{1, 0}}},
			want: true,
		},
		{
			name: "numCourses = 2, prerequisites = \"[[1,0],[0,1]]\"",
			args: args{numCourses: 2, prerequisites: [][]int{{1, 0}, {0, 1}}},
			want: false,
		},
		{
			name: "numCourses = 4, prerequisites = \"[[1,0],[2,1],[3,2]]\"",
			args: args{numCourses: 4, prerequisites: [][]int{{1, 0}, {2, 1}, {3, 2}}},
			want: true,
		},
		{
			name: "numCourses = 5, prerequisites = \"[[0,1],[0,2],[1,3],[1,4],[3,4]]\"",
			args: args{numCourses: 4, prerequisites: [][]int{{0, 1}, {0, 2}, {1, 3}, {1, 4}, {3, 4}}},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canFinish(tt.args.numCourses, tt.args.prerequisites); got != tt.want {
				t.Errorf("canFinish() = %v, want %v", got, tt.want)
			}
		})
	}
}
