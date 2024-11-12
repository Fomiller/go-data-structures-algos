package main

type MovingQueue struct {
	elements []int
	sum      int
	length   int
}

func (q *MovingQueue) Push(v int) {
	if len(q.elements) >= q.length {
		_ = q.Pop()
	}

	q.elements = append(q.elements, v)
	q.sum += v
}

func (q *MovingQueue) Pop() int {
	if len(q.elements) < 1 {
		return -1
	}

	element := q.elements[0]
	q.elements = q.elements[1:]

	q.sum -= element
	return element
}

func (q *MovingQueue) Average() float64 {
	return float64(q.sum / len(q.elements))
}
