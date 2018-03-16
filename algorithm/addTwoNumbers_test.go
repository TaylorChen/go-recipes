package main

import "testing"

func TestAddTwoNumbers(t *testing.T) {
	var node1 ListNode
	var node2 ListNode
	var node3 ListNode
	var node4 ListNode

	node2.Val = 4
	node2.Next = nil
	node1.Val = 2
	node1.Next = &node2

	node4.Val = 7
	node4.Next = nil
	node3.Val = 5
	node3.Next = &node4

	point := addTwoNumbers(&node1, &node3)
	for nil != point {
		if !(point.Val == 7 || point.Val == 1) {
			t.Errorf("Expected %d , %d , %d but actual got %d!", 7, 1, 1, point.Val)
		}
		point = point.Next
	}
}
