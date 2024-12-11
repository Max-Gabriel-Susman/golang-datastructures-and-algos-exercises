package main

// Given the head of a singly linked list, return the middle node of the linked list.
// If there are two middle nodes, return the second middle node.
func middleNode(head *ListNode) *ListNode {
	tailFound := false
	nodes := []ListNode{}
	nodes = append(nodes, *head)
	next := head.Next
	if next != nil {
		for !tailFound {
			if next == nil {
				break
			}
			nodes = append(nodes, *next)
			next = next.Next
		}
	}

	length := len(nodes)
	if length < 2 {
		return &nodes[0]
	}

	half := length / 2
	if length%half == 1 {
		return &nodes[half]
	}

	return &nodes[half]
}
