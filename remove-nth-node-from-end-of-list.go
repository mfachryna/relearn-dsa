package relearndsa

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	point := &ListNode{Val: 0, Next: head}
	calc := 0
	temp := point
	for temp != nil {
		temp = temp.Next
		calc++
	}

	temp = point
	for i := 1; i < calc-n; i++ {
		temp = temp.Next
	}
	temp.Next = temp.Next.Next
	return point.Next
}
