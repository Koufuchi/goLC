package longestsubstringwithoutrepeatingcharacters

import "testing"

func BenchmarkTest(b *testing.B) {
	s := "abcabcbb"
	for idx := 0; idx < b.N; idx++ {
		lengthOfLongestSubstring(s)
	}
}
func Test_lengthOfLongestSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "s = \"abcabcbb\"",
			args: args{s: "abcabcbb"},
			want: 3,
		},
		{
			name: "s = \"bbbbb\"",
			args: args{s: "bbbbb"},
			want: 1,
		},
		{
			name: "s = \"pwwkew\"",
			args: args{s: "pwwkew"},
			want: 3,
		},
		{
			name: "s = \"\"",
			args: args{s: ""},
			want: 0,
		},
		{
			name: "s = \"au\"",
			args: args{s: "au"},
			want: 2,
		},
		{
			name: "s = \"dvdf\"",
			args: args{s: "dvdf"},
			want: 3,
		},
		{
			name: "s = \"aabaab!bb\"",
			args: args{s: "aabaab!bb"},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
