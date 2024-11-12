package main

type binaryNode struct {
	Value int
	Left  *binaryNode
	Right *binaryNode
}

func (n *binaryNode) Insert(value int) {
	if value < n.Value {
		if n.Left == nil {
			n.Left = &binaryNode{Value: value}
		} else {
			n.Left.Insert(value)
		}
	} else {
		if n.Right == nil {
			n.Right = &binaryNode{Value: value}
		} else {
			n.Right.Insert(value)
		}
	}
}

func (n *binaryNode) InOrder() []int {
	if n == nil {
		return nil
	}
	output := []int{}
	output = append(output, n.Left.InOrder()...)
	output = append(output, n.Value)
	output = append(output, n.Right.InOrder()...)

	return output
}
