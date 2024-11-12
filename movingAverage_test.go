package main

import (
	"fmt"
	"testing"
)

func TestMovingAverage1(t *testing.T) {
	q := MovingQueue{length: 4}

	q.Push(1)
	if q.Average() != 1 {
		t.Fatalf("Incorrect: %v, %v", q.sum, q.Average())
	}

	q.Push(3)
	if q.Average() != 2 {
		t.Fatalf("Incorrect: %v, %v", q.sum, q.Average())
	}

	q.Push(5)
	if q.Average() != 3 {
		t.Fatalf("Incorrect: %v, %v", q.sum, q.Average())
	}

	q.Push(7)
	if q.Average() != 4 {
		t.Fatalf("Incorrect: %v, %v", q.sum, q.Average())
	}

	q.Push(5)
	if q.Average() != 5 {
		t.Fatalf("Incorrect: %v, %v", q.sum, q.Average())
	}

	fmt.Printf("MovingAverage1: %v, %v\n", q.sum, q.Average())
}
