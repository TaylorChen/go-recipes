package algorithm

//import "fmt"

/*
 *  Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
 *  Output: 7 -> 0 -> 8
 *  Explanation: 342 + 465 = 807.
 *
 *
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	node := &ListNode{0, nil}
	p := node
	for nil != l1 || nil != l2 || carry > 0 {
		if l1 != nil {
			carry += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			carry += l2.Val
			l2 = l2.Next
		}

		p.Next = &ListNode{
			Val: carry % 10,
		}
		carry /= 10
		p = p.Next
	}
	return node.Next
}

/*
func main() {
	var node1 ListNode
	var node2 ListNode
	var node3 ListNode
	var node4 ListNode

	node2.Val = 4
	node2.Next = nil
	node1.Val = 2
	node1.Next = &node2
	node1.Next = &node2

	//node1.Val = 5
	//node1.Next = nil

	node4.Val = 7
	node4.Next = nil
	node3.Val = 5
	node3.Next = &node4

	//node3.Val = 5
	//node3.Next = nil
	point := addTwoNumbers(&node1, &node3)
	for nil != point {
		fmt.Println(point.Val)
		point = point.Next
	}
}
*/
