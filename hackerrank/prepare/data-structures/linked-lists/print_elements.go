package linkedlists

import "fmt"

// printLinkedList Problem: Print the Elements of a Linked List.
//
// HackerRank: https://www.hackerrank.com/challenges/print-the-elements-of-a-linked-list/problem
func printLinkedList(head *SinglyLinkedListNode) {
	itm := head
	for {
		fmt.Println(itm.data)
		itm = itm.next
		if itm == nil {
			return
		}
	}
}
