# Challenge #45: Linked List Cycle Detection

## Objective
Write a Go program to detect if a given singly linked list contains a cycle. A cycle exists if a node in the list can be reached again by continuously following the `Next` pointer.

## Difficulty
Hard

## Concepts Tested
* Linked Lists (Structs and Pointers)
* Algorithmic Thinking (Floyd's Tortoise and Hare Algorithm)
* Pointers

## Rules/Logic
1.  Define a struct named `ListNode` with two fields: `Val` (int) and `Next` (a pointer to another `ListNode`, `*ListNode`).
2.  Create a function `hasCycle(head *ListNode) bool`.
3.  The input `head` is a pointer to the first node of the linked list.
4.  The function should return `true` if the list contains a cycle, and `false` otherwise.
5.  **The Challenge**: You should aim for an efficient solution that uses **O(1)** (constant) memory. This means you cannot use a map or set to store the nodes you've already visited.
6.  **Hint (Floyd's Tortoise and Hare Algorithm)**:
    * Initialize two pointers, `slow` and `fast`, both starting at the `head`.
    * Use a loop that continues as long as `fast` and `fast.Next` are not `nil`.
    * Inside the loop, move `slow` one step forward (`slow = slow.Next`).
    * Move `fast` two steps forward (`fast = fast.Next.Next`).
    * If `slow` and `fast` ever point to the same node (`slow == fast`), it means there is a cycle in the list. Return `true`.
7.  If the loop finishes (meaning `fast` or `fast.Next` reached `nil`), it means the end of the list was reached without finding a cycle. Return `false`.
8.  In your `main` function, create a sample linked list *with* a cycle and another one *without* a cycle, and test your `hasCycle` function on both.

## Your Tasks
1.  Create a new directory: `problems/challenges/45_linked_list_cycle/`.
2.  Inside this directory, create `README.md` and `main.go`.
3.  Copy and paste this problem statement into your `README.md`.
4.  Write your efficient, O(1) memory solution in `main.go`.

## Sample Input/Output

```go
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

// Expected Terminal Output:
List 1 has cycle: true
List 2 has cycle: false