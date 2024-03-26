package reverseLinkedList

import (
	ll "goLC/DataStructure/linkedList"
)

type ListNode = ll.ListNode

func ReverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	temp := head.Next
	head.Next = nil
	for temp != nil {
		nT := temp.Next
		temp.Next = head
		head = temp
		temp = nT
	}

	return head
}
