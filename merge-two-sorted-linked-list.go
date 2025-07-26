package relearndsa

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	temp := &ListNode{}
	op := temp

	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			op.Next = list1
			list1 = list1.Next
		} else {
			op.Next = list2
			list2 = list2.Next
		}
		op = op.Next
	}
	if list1 != nil {
		op.Next = list1
	} else {
		op.Next = list2
	}
	return temp.Next
}
