package reverseLinkedList

import (
	ll "goLC/DataStructure/linkedList"
	"reflect"
	"testing"
)

func TestReverseList(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "head = \"[1,2,3,4,5]\"",
			args: args{head: ll.GenLinkedListFromInts([]int{1, 2, 3, 4, 5})},
			want: ll.GenLinkedListFromInts([]int{5, 4, 3, 2, 1}),
		},
		{
			name: "head = \"[1,2]\"",
			args: args{head: ll.GenLinkedListFromInts([]int{1, 2})},
			want: ll.GenLinkedListFromInts([]int{2, 1}),
		},
		{
			name: "head = \"[]\"",
			args: args{head: ll.GenLinkedListFromInts([]int{})},
			want: ll.GenLinkedListFromInts([]int{}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseList(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReverseList() = %v, want %v", got, tt.want)
			}
		})
	}
}
