package main

import (
	"fmt"
)

type Node struct {
	Value int
	Next  *Node
}

type LinkedList struct {
	Head *Node
	Size int
}

func (ll *LinkedList) InsertAtTail(value int) {
	newNode := &Node{Value: value}

	if ll.Head == nil {
		ll.Head = newNode
		ll.Size++
	} else {
		current := ll.Head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = newNode
		ll.Size++
	}
}

func (ll *LinkedList) InsertAtHead(value int) {
	newNode := &Node{Value: value, Next: ll.Head}
	ll.Head = newNode
	ll.Size++
}

func (ll *LinkedList) InsertAtPosition(idx int, value int) (*Node, error) {
	newNode := &Node{Value: value}

	if idx < 0 || idx > ll.Size {
		return nil, fmt.Errorf("Insert index out of range")
	}

	// if empty and inserting at head
	if idx == 0 {
		newNode.Next = ll.Head
		ll.Head = newNode
		ll.Size++
		return newNode, nil
	}

	// set previous to head and move forward until at idx
	previous := ll.Head
	for i := 1; i < idx; i++ {
		previous = previous.Next
	}

	newNode.Next = previous.Next
	previous.Next = newNode
	ll.Size++
	return newNode, nil
}

func (ll *LinkedList) DeleteByValue(value int) (*Node, error) {
	if ll.IsEmpty() {
		return nil, fmt.Errorf("Linked List is empty")
	} else {
		var previous *Node
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
	}
	return nil, fmt.Errorf("Value not found in list")
}

func (ll *LinkedList) DeleteFromHead() (*Node, error) {
	if ll.IsEmpty() {
		return nil, fmt.Errorf("Linked List is empty")
	} else {
		head := ll.Head
		ll.Head = ll.Head.Next
		ll.Size--
		return head, nil
	}
}

func (ll *LinkedList) DeleteFromTail() (*Node, error) {
	if ll.IsEmpty() {
		return nil, fmt.Errorf("Linked List is empty")
	}

	current := ll.Head
	if current.Next == nil {
		ll.Head = nil
		ll.Size--
		return current, nil
	}

	var previous *Node
	for current.Next != nil {
		previous = current
		current = current.Next
	}

	previous.Next = nil
	ll.Size--
	return current, nil
}

func (ll *LinkedList) IsEmpty() bool {
	return ll.Head == nil
}

func (ll *LinkedList) Length() int {
	return ll.Size
}

func (ll *LinkedList) String() string {
	var result []int
	current := ll.Head
	for current != nil {
		result = append(result, current.Value)
		current = current.Next
	}
	return fmt.Sprintf("[%v]", result)
}
