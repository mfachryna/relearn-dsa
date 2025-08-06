package relearndsa

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	temp := head.Next
	head.Next = head.Next.Next
	temp.Next = head
	head = temp

	head.Next.Next = swapPairs(head.Next.Next)

	return head
}
