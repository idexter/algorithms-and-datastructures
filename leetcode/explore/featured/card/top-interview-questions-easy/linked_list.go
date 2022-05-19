package top_interview_questions_easy

import "strconv"

// ListNode describe LinkedList node.
type ListNode struct {
	Val  int
	Next *ListNode
}

// deleteNode Problem: Delete Node in a Linked List.
//
// LeetCode: https://leetcode.com/explore/featured/card/top-interview-questions-easy/93/linked-list/553/
func deleteNode(node *ListNode) {
	*node = *node.Next
}

// removeNthFromEnd Problem: Remove Nth Node From End of List.
//
// LeetCode: https://leetcode.com/explore/featured/card/top-interview-questions-easy/93/linked-list/603/
//
// Hint 1:
//	Maintain two pointers and update one with a delay of n steps.
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil && n == 1 {
		return nil
	}

	last := head
	elementToDelete := head
	beforeLast := head
	k := 0
	for {
		if k >= n {
			elementToDelete = elementToDelete.Next
		}
		if last.Next != nil && last.Next.Next == nil {
			beforeLast = last
		}
		if last.Next == nil {
			break
		}
		last = last.Next
		k++
	}

	if elementToDelete == head {
		return elementToDelete.Next
	}

	if n == 1 && k == 1 {
		head.Next = nil
		return head
	}

	if elementToDelete == last {
		beforeLast.Next = nil
		return head
	}

	*elementToDelete = *elementToDelete.Next
	return head
}

// reverseList Problem: Reverse Linked List.
//
// LeetCode: https://leetcode.com/explore/featured/card/top-interview-questions-easy/93/linked-list/560/
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var prev *ListNode
	// Go through whole list
	for head.Next != nil {
		next := head.Next // Save next
		head.Next = prev  // Set next to previous
		prev = head       // set new previous to current
		head = next       // set head to next
	}
	head.Next = prev // for last case set next to previous
	return head
}

// mergeTwoLists Problem: Merge Two Sorted Lists.
//
// LeetCode: https://leetcode.com/explore/featured/card/top-interview-questions-easy/93/linked-list/771/
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

// isPalindromeList Problem: Palindrome Linked List.
//
// LeetCode: https://leetcode.com/explore/featured/card/top-interview-questions-easy/93/linked-list/772/
func isPalindromeList(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	var s1, s2 string
	for head != nil {
		s1 += strconv.Itoa(head.Val)
		s2 = strconv.Itoa(head.Val) + s2
		head = head.Next
	}

	return s1 == s2
}

// hasCycle Problem: Linked List Cycle.
//
// LeetCode: https://leetcode.com/explore/featured/card/top-interview-questions-easy/93/linked-list/773/
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	existsMap := make(map[*ListNode]struct{})
	for head.Next != nil {
		if _, ok := existsMap[head]; ok {
			return true
		}
		existsMap[head] = struct{}{}
		head = head.Next
	}
	return false
}
