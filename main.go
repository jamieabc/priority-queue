package main

import "fmt"

type MinPriorityQueue interface {
	Poll() interface{}
	Remove(int)
	Insert(int)
}

type minPriorityQueue struct {
	heap []int
}

func swap(q *minPriorityQueue, src, dst int) {
	q.heap[src], q.heap[dst] = q.heap[dst], q.heap[src]
}

func leftChild(q *minPriorityQueue, index int) interface{} {
	childIndex := 2*index + 1
	if len(q.heap) <= childIndex {
		return nil
	}

	return childIndex
}

func rightChild(q *minPriorityQueue, index int) interface{} {
	childIndex := 2*index + 2
	if len(q.heap) <= childIndex {
		return nil
	}

	return childIndex
}

func parent(index int) int {
	if index == 0 {
		return 0
	}

	return (index - 1) / 2
}

func smallestIndex(q *minPriorityQueue, idx1, idx2 interface{}) int {
	if idx1 == nil {
		return idx2.(int)
	}

	if idx2 == nil {
		return idx1.(int)
	}

	if q.heap[idx1.(int)] <= q.heap[idx2.(int)] {
		return idx1.(int)
	} else {
		return idx2.(int)
	}
}

func (q *minPriorityQueue) Poll() interface{} {
	if len(q.heap) == 0 {
		return nil
	}

	target := q.heap[0]
	swap(q, 0, len(q.heap)-1)
	q.heap = q.heap[:len(q.heap)-1]
	bubbleDown(q, 0)
	return target
}

func bubbleDown(q *minPriorityQueue, index int) {
	for index < len(q.heap) {
		leftChild := leftChild(q, index)
		rightChild := rightChild(q, index)

		if leftChild == nil && rightChild == nil {
			return
		}

		target := smallestIndex(q, leftChild, rightChild)
		swap(q, index, target)
		index = target
	}
}

func (q *minPriorityQueue) Remove(i int) {
	target := search(q, i)
	if target == nil {
		return
	}

	index := target.(int)
	swap(q, index, len(q.heap)-1)
	q.heap = q.heap[:len(q.heap)-1]
	bubbleDown(q, index)
}

func search(q *minPriorityQueue, val int) interface{} {
	for i, v := range q.heap {
		if v == val {
			return i
		}
	}
	return nil
}

func (q *minPriorityQueue) Insert(i int) {
	q.heap = append(q.heap, i)
	bubbleUp(q, len(q.heap)-1)
}

func bubbleUp(q *minPriorityQueue, index int) {
	for index != 0 {
		p := parent(index)
		if q.heap[index] <= q.heap[p] {
			swap(q, index, p)
			index = p
		} else {
			return
		}
	}
}

func newMinPriorityHeap() MinPriorityQueue {
	return &minPriorityQueue{heap: make([]int, 0)}
}

func main() {
	h := newMinPriorityHeap()
	h.Insert(5)
	h.Insert(3)
	h.Insert(-2)
	h.Insert(10)
	polled := h.Poll()
	if polled == nil {
		fmt.Println("polled empty")
	} else {
		fmt.Printf("polled: %d\n", polled.(int))
	}

	h.Remove(3)
	polled = h.Poll()
	if polled == nil {
		fmt.Println("polled empty")
	} else {
		fmt.Printf("polled: %d\n", polled.(int))
	}

	h.Remove(10)
	polled = h.Poll()
	if polled == nil {
		fmt.Println("polled empty")
	} else {
		fmt.Printf("polled: %d\n", polled.(int))
	}
}
