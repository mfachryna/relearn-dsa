package relearndsa

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	// defer runtime.GC()
	// if headA == nil || headB == nil {
	//     return nil
	// }
	var pointA, pointB *ListNode
	pointA, pointB = headA, headB

	for pointA != pointB {
		if pointA == nil {
			pointA = headB
		} else {
			pointA = pointA.Next
		}

		if pointB == nil {
			pointB = headA
		} else {
			pointB = pointB.Next
		}
	}
	return pointA
}
