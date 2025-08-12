package relearndsa

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	turtle, hare := head, head
	for hare != nil && hare.Next != nil {
		turtle = turtle.Next
		hare = hare.Next.Next

		if turtle == hare {
			return true
		}
	}
	return false
}
