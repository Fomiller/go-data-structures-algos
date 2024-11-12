package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestBinaryTree(t *testing.T) {
	values := []int{3, 4, 5, 3, 15, 8, 132, 18}
	expected := []int{3, 3, 4, 5, 8, 10, 15, 18, 132}

	root := &binaryNode{Value: 10}

	for _, v := range values {
		root.Insert(v)
	}

	result := root.InOrder()

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Incorrect: %v", result)
	} else {
		fmt.Printf("binaryTree1: %v\n", result)
	}
}
