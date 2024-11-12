package main

import "fmt"

type Queue struct {
	elements []int
}

func (q *Queue) Enqueue(value int) {
	q.elements = append(q.elements, value)
}

func (q *Queue) Dequeue() (int, error) {
	if q.IsEmpty() {
		return 0, fmt.Errorf("Queue is empty")
	} else {
		value := q.elements[0]
		q.elements = q.elements[1:]
		return value, nil
	}
}

func (q *Queue) Front() (int, error) {
	if q.IsEmpty() {
		return 0, fmt.Errorf("Queue is empty")
	} else {
		return q.elements[0], nil
	}
}

func (q *Queue) IsEmpty() bool {
	return len(q.elements) == 0
}

func (q *Queue) Size() int {
	return len(q.elements)
}
