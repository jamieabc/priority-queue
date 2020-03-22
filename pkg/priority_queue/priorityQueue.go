package priority_queue

// PriorityQueue - provide methods to operate priority queue
type PriorityQueue interface {
	Poll() interface{}
	Offer(interface{})
	Peek() interface{}
	Comparator(func(interface{}, interface{}) bool)
	Size() int
}

type priorityQueue struct {
	data        []interface{}
	size        int
	compareFunc func(interface{}, interface{}) bool
}

func (q *priorityQueue) Poll() interface{} {
	if len(q.data) == 0 {
		return nil
	}

	polled := q.data[0]
	q.swap(0, len(q.data)-1)
	q.data = q.data[:len(q.data)-1]

	q.bubbleDown(0)
	return polled
}

func (q *priorityQueue) Offer(i interface{}) {
	if len(q.data) < q.size {
		q.data = append(q.data, i)
		idx := q.bubbleUp(len(q.data) - 1)
		q.bubbleDown(idx)
	} else {
		if q.compareFunc(q.data[q.size-1], i) {
			q.data[q.size-1] = i
			idx := q.bubbleUp(q.size - 1)
			q.bubbleDown(idx)
		}
	}
}

func (q *priorityQueue) Size() int {
	return len(q.data)
}

func (q *priorityQueue) Peek() interface{} {
	if len(q.data) == 0 {
		return nil
	}
	return q.data[0]
}

func (q *priorityQueue) Comparator(f func(interface{}, interface{}) bool) {
	q.compareFunc = f
}

func parent(idx int) int {
	return (idx - 1) / 2
}

func (q *priorityQueue) swap(i, j int) {
	q.data[i], q.data[j] = q.data[j], q.data[i]
}

func (q *priorityQueue) bubbleUp(idx int) int {
	for current := idx; current != 0; {
		p := parent(current)
		if q.compareFunc(q.data[p], q.data[current]) {
			return current
		} else {
			q.swap(p, current)
			current = p
		}
	}
	return 0
}

func leftChild(idx int) int {
	return idx*2 + 1
}

func rightChild(idx int) int {
	return idx*2 + 2
}

func (q *priorityQueue) bubbleDown(idx int) {
	for current := idx; current < len(q.data)-1; {
		l := leftChild(idx)
		r := rightChild(idx)

		ls, rs := true, true
		if l < len(q.data) {
			ls = q.compareFunc(q.data[current], q.data[l])
		}

		if r < len(q.data) {
			rs = q.compareFunc(q.data[current], q.data[r])
		}

		if ls && rs {
			return
		}

		// in case right child not exist or l is end of data
		if l == len(q.data)-1 || r >= len(q.data) {
			q.swap(l, current)
			current = l
			continue
		}

		if q.compareFunc(q.data[l], q.data[r]) {
			q.swap(l, current)
			current = l
		} else {
			q.swap(r, current)
			current = r
		}
	}
}

func New(size int) PriorityQueue {
	return &priorityQueue{
		data: make([]interface{}, 0),
		size: size,
		compareFunc: func(i, j interface{}) bool {
			return i.(int) <= j.(int)
		},
	}
}
