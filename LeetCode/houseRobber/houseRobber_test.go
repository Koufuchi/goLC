package houseRobber

import "testing"

func Test_rob(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "nums = \"[2,3,2]\"",
			args: args{nums: []int{2, 3, 2}},
			want: 4,
		},
		{
			name: "nums = \"[1,2,3,1]\"",
			args: args{nums: []int{1, 2, 3, 1}},
			want: 4,
		},
		{
			name: "nums = \"[2,7,9,3,1]\"",
			args: args{nums: []int{2, 7, 9, 3, 1}},
			want: 12,
		},
		{
			name: "nums = \"[1]\"",
			args: args{nums: []int{1}},
			want: 1,
		},
		{
			name: "nums = \"[1,2,1,1]\"",
			args: args{nums: []int{1, 2, 1, 1}},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rob(tt.args.nums); got != tt.want {
				t.Errorf("rob() = %v, want %v", got, tt.want)
			}
		})
		t.Run(tt.name, func(t *testing.T) {
			if got := robWithMem(tt.args.nums); got != tt.want {
				t.Errorf("robWithMem() = %v, want %v", got, tt.want)
			}
		})
	}
}
