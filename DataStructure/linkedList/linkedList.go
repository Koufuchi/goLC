package linkedList

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func GenLinkedListFromInts(ints []int) *ListNode {
	if len(ints) == 0 {
		return nil
	}

	head := &ListNode{Val: ints[0]}
	current := head

	for i := 1; i < len(ints); i++ {
		node := &ListNode{Val: ints[i]}
		current.Next = node
		current = node
	}

	return head
}
