package main

import (
	"fmt"
)

type Node3 struct {
	Value int
	Next  *Node3
}

type LinkedList3 struct {
	Head *Node3
	Size uint
}

func (ll *LinkedList3) Insert(value int) {
	newNode := &Node3{Value: value}
	current := ll.Head

	for current.Next != nil {
		current = current.Next
	}

	current.Next = newNode
	ll.Size++
}

func (ll *LinkedList3) Length() uint {
	return ll.Size
}

func (ll *LinkedList3) IsEmpty() bool {
	return ll.Size > 0
}

func (ll *LinkedList3) DeleteFromHead() {
	if ll.IsEmpty() {
		return
	}

	ll.Head = ll.Head.Next
	ll.Size--
}

func (ll *LinkedList3) DeleteFromTail() {
	current := ll.Head
	if ll.IsEmpty() {
		return
	}

	var previous *Node3
	for current.Next != nil {
		previous = current
		current = current.Next
	}

	previous.Next = nil
	ll.Size--
}

func (ll *LinkedList3) DeleteByValue(value int) (*Node3, error) {
	if ll.IsEmpty() {
		return nil, fmt.Errorf("Linked List is empty")
	}

	var previous *Node3
	current := ll.Head

	// while next is not nil
	for current != nil {

		// if the current nodes value matches the value to be deleted
		if current.Value == value {
			// if we are at the head of the list
			if previous == nil {
				ll.Head = current.Next
			} else {
				// if we are in the middle of the list
				previous.Next = current.Next
			}

			ll.Size--
			return current, nil
		}

		previous = current
		current = current.Next
	}

	return nil, fmt.Errorf("Value not found in list")
}
