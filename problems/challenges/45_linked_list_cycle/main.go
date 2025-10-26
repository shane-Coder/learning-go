package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}

func main() {
	// Sample code in your main function:

	// List 1: 1 -> 2 -> 3 -> 4 -> (points back to 2)
	node1 := &ListNode{Val: 1}
	node2 := &ListNode{Val: 2}
	node3 := &ListNode{Val: 3}
	node4 := &ListNode{Val: 4}
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node2 // Creates the cycle

	fmt.Println("List 1 has cycle:", hasCycle(node1))

	// List 2: 1 -> 2 -> 3 -> 4 -> nil (no cycle)
	nodeA := &ListNode{Val: 1}
	nodeB := &ListNode{Val: 2}
	nodeC := &ListNode{Val: 3}
	nodeD := &ListNode{Val: 4}
	nodeA.Next = nodeB
	nodeB.Next = nodeC
	nodeC.Next = nodeD

	fmt.Println("List 2 has cycle:", hasCycle(nodeA))

}
