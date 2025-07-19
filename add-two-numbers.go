package relearndsa

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l3 := &ListNode{}
	current := l3
	temp := 0
	for l1 != nil || l2 != nil || temp != 0 {
		sum := temp

		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		temp = sum / 10
		current.Next = &ListNode{Val: sum % 10}
		current = current.Next
	}

	return l3.Next
}
