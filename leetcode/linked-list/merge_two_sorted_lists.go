package linkedlist

// mergeTwoLists merges two sorted lists.
// See: https://leetcode.com/explore/interview/card/top-interview-questions-easy/93/linked-list/771/
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	} else if l2 == nil {
		return l1
	} else if l1 == nil {
		return l2
	}

	if l1.Next == nil && l2.Next == nil {
		if l1.Val > l2.Val {
			l2.Next = l1
			return l2
		}
		l1.Next = l2
		return l1
	}

	first := 0
	if l1.Val > l2.Val {
		first = l2.Val - 1
	} else {
		first = l1.Val - 1
	}
	newHead := &ListNode{Val: first, Next: nil}
	newList := newHead

	for {
		if l1 == nil && l2 == nil {
			break
		}

		if l1 == nil {
			tmp := *l2
			tmp.Next = nil
			l2 = l2.Next
			newHead.Next = &tmp
			newHead = newHead.Next
			continue
		}

		if l2 == nil {
			tmp := *l1
			tmp.Next = nil
			l1 = l1.Next
			newHead.Next = &tmp
			newHead = newHead.Next
			continue
		}

		if l1.Val < l2.Val {
			tmp := *l1
			tmp.Next = nil
			l1 = l1.Next
			newHead.Next = &tmp
			newHead = newHead.Next
			continue
		}

		if l2.Val < l1.Val {
			tmp := *l2
			tmp.Next = nil
			l2 = l2.Next
			newHead.Next = &tmp
			newHead = newHead.Next
			continue
		}

		if l1.Val == l2.Val {
			tmp := *l1
			tmp.Next = nil
			l1 = l1.Next
			newHead.Next = &tmp
			newHead = newHead.Next

			tmp2 := *l2
			tmp2.Next = nil
			l2 = l2.Next
			newHead.Next = &tmp2
			newHead = newHead.Next
			continue
		}
	}

	return newList.Next
}
