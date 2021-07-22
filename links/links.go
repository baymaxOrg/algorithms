package links

type SingleLinkNode struct {
	Data int
	Next *SingleLinkNode
}

// type SingleLinkList struct {
// }

func (l *SingleLinkNode) createLinkListTail(nums []int) *SingleLinkNode {
	if nums == nil {
		return &SingleLinkNode{}
	}
	head := &SingleLinkNode{
		Data: nums[0],
	}

	tmp := head
	for i := 1; i < len(nums); i++ {
		tmp.Next = &SingleLinkNode{
			Data: nums[i],
		}
		tmp = tmp.Next
	}
	return head
}

func (l *SingleLinkNode) createLinkListHead(nums []int) *SingleLinkNode {
	if nums == nil {
		return &SingleLinkNode{}
	}
	head := &SingleLinkNode{
		// Data: nums[0],
	}

	tmp := head
	for i := 0; i < len(nums); i++ {
		tmp = &SingleLinkNode{
			Data: nums[i],
		}
		tmp.Next = head.Next
		head.Next = tmp

	}
	return head.Next
}


//func isPalindrome(head *ListNode) bool {
//	if head == nil {
//		return head
//	}
//
//	slow := halfList(head)
//	newHead := reverse(slow.Next)
//
//	p1, p2 := head, newHead
//
//	result := true
//
//	for result && p2 != nil {
//		if p1.Val != p2.Val {
//			result false
//		}
//		p1 = p1.Next
//		p2 = p2.Next
//	}
//	slow.Next = reverse(newHead)
//	return result
//}
//
//func halfList(head *ListNode) *ListNode {
//	fast, slow := head, head
//
//	for fast.Next != nil && Fast.Next.Next != nil {
//		fast = fast.Next.Next
//		slow = slow.Next
//	}
//	return slow
//}
//
//func reverse(head *ListNode) *ListNode {
//	if head == nil || head.Next == nil {
//		return head
//	}
//	var pre *ListNode = nil
//	var cur *ListNode = head
//	var tmp *ListNode
//	for cur != nil {
//		tmp = cur.Next
//		cur.Next = pre
//		pre = cur
//		cur = tmp
//	}
//	return pre
//}