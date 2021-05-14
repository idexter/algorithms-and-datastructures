package linked_list

import "strconv"

// isPalindrome checks if LinkedList is palindrome.
// See: https://leetcode.com/explore/interview/card/top-interview-questions-easy/93/linked-list/772/
func isPalindrome(head *ListNode) bool {
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
