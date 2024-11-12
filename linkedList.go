package main

//
// type Node struct {
// 	Value int
// 	Next  *Node
// }
//
// type LinkedList struct {
// 	Head *Node
// }
//
// func (ll *LinkedList) Insert(value int) {
// 	//create new node to insert
// 	newNode := &Node{Value: value}
// 	// if list is empty new node becomes the head
// 	if ll.Head == nil {
// 		ll.Head = newNode
// 	} else {
// 		// if list is not empty get the Head
// 		current := ll.Head
//
// 		// call current.Next until we find a Node that does not have a next
// 		for current.Next != nil {
// 			current = current.Next
// 		}
// 		// set the current.Next node to the new node
// 		current.Next = newNode
// 	}
// }
//
// func (ll *LinkedList) Length() int {
//
// 	if ll.Head == nil {
// 		return 0
// 	} else {
// 		len := 1
// 		current := ll.Head
//
// 		for current.Next != nil {
// 			len++
// 			current = current.Next
// 		}
//
// 		return len
// 	}
// }
//
// func (ll *LinkedList) Search(value int) *Node {
// 	if ll.Head != nil {
// 		current := ll.Head
// 		if current.Value == value {
// 			return current
// 		}
// 		for current.Next != nil {
// 			if current.Value == value {
// 				return current
// 			} else {
// 				current = current.Next
// 			}
// 		}
//
// 	}
//
// 	return nil
// }
