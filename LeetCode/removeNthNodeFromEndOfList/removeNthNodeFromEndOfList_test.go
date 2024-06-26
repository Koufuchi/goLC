package removeNthNodeFromEndOfList

import (
	ll "goLC/DataStructure/linkedList"
	"reflect"
	"testing"
)

func Test_removeNthFromEnd(t *testing.T) {
	type args struct {
		head *ll.ListNode
		n    int
	}
	tests := []struct {
		name string
		args args
		want *ll.ListNode
	}{
		{
			name: "head = \"[1,2,3,4,5]\", n = 2",
			args: args{head: ll.GenLinkedListFromInts([]int{1, 2, 3, 4, 5}), n: 2},
			want: ll.GenLinkedListFromInts([]int{1, 2, 3, 5}),
		},
		{
			name: "head = \"[1]\", n = 1",
			args: args{head: ll.GenLinkedListFromInts([]int{1}), n: 1},
			want: ll.GenLinkedListFromInts([]int{}),
		},
		{
			name: "head = \"[1,2]\", n = 1",
			args: args{head: ll.GenLinkedListFromInts([]int{1, 2}), n: 1},
			want: ll.GenLinkedListFromInts([]int{1}),
		},
		{
			name: "head = \"[1,2]\", n = 2",
			args: args{head: ll.GenLinkedListFromInts([]int{1, 2}), n: 2},
			want: ll.GenLinkedListFromInts([]int{2}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeNthFromEnd(tt.args.head, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeNthFromEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}
