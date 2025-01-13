package main

import (
	"log"
	"slices"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// Given the head of a sorted linked list, delete all duplicates
// such that each element appears only once. Return the linked list sorted as well.
func deleteDuplicates(head *ListNode) *ListNode {
	tailFound := false
	uniqueValues := []int{head.Val}
	nodes := []ListNode{}
	nodes = append(nodes, *head)
	next := head.Next
	// we basically need the logic to exclude the node possessing a duplicate value
	// and the relink the previous to to the following node with a unique value
	if next != nil {
		for !tailFound {
			if next == nil {
				break
			}
			if !slices.Contains(uniqueValues, next.Val) {
				uniqueValues = append(uniqueValues, next.Val)
			}
			nodes = append(nodes, *next)
			next = next.Next
		}
	}

	log.Println("uniqueValues: ", uniqueValues)

	return head
}
