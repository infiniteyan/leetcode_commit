package main

import "fmt"

type ListNode struct {
	next 	*ListNode
	val		int
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var length int
	ptr := head
	for ptr != nil {
		length++
		ptr = ptr.next
	}

	if length < 2 {
		return head
	}
	var newHead *ListNode = nil

	ptr = head
	var currentLen  = 0
	var nextHead *ListNode = nil
	var prev *ListNode = nil
	for currentLen + 2 <= length && ptr != nil {
		nextHead = ptr.next.next
		secPtr := ptr.next
		secPtr.next = ptr
		ptr.next = nextHead

		if newHead == nil {
			newHead = secPtr
		}
		if prev != nil {
			prev.next = secPtr
		}
		prev = ptr

		ptr = nextHead
		currentLen += 2
	}
	return newHead
}

func main() {
	first := ListNode{next:nil, val: 1}
	sec := ListNode{next:nil, val: 2}
	thir := ListNode{next:nil, val: 3}
	forth := ListNode{next:nil, val: 4}
	first.next = &sec
	sec.next = &thir
	thir.next = &forth

	newList := reverseList(&first)
	ptr := newList
	for ptr != nil {
		fmt.Printf("%d  ", ptr.val)
		ptr = ptr.next
	}
	fmt.Println()
}
