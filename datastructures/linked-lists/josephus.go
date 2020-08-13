package main

// https://www.geeksforgeeks.org/josephus-circle-using-circular-linked-list/
import "fmt"

type ListNode struct { // defines a ListNode in a singly linked list
	data interface{} // the datum
	next *ListNode   // pointer to the next ListNode
}

// Create a new node of circular linked list
func NewListNode(data int) *ListNode {
	temp := &ListNode{}
	temp.next = temp
	temp.data = data
	return temp
}

// Function to find the only person left after one in every m-th node is killed in a circle of n nodes
func getJosephusPosition(m, n int) {
	// Create a circular linked list of size N.
	head := NewListNode(1)
	prev := head
	for i := 2; i <= n; i++ {
		prev.next = NewListNode(i)
		prev = prev.next
	}
	prev.next = head // Connect last node to first

	// while only one node is left in the linked list
	ptr1, ptr2 := head, head
	for ptr1.next != ptr1 {
		// Find m-th node
		count := 1
		for count != m {
			ptr2 = ptr1
			ptr1 = ptr1.next
			count++
		}

		// Remove the m-th node
		ptr2.next = ptr1.next
		ptr1 = ptr2.next
	}
	fmt.Println("Last person left standing ", "(Josephus Position) is ", ptr1.data)
}

// Driver program to test above functions
func main() {
	n, m := 14, 2
	getJosephusPosition(m, n)
}
