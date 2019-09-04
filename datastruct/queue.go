package datastruct

import "fmt"

type Queue struct {
	queue []int
}

func NewQueue() *Queue {
	return &Queue{
		queue: make([]int, 0),
	}
}

func (q *Queue) Push(k int) {
	q.queue = append(q.queue, k)
}

func (q *Queue) Pop() (int, error) {
	if q.Empty() {
		return 0, fmt.Errorf("Pop from empty queue")
	}
	v := q.queue[0]
	q.queue = q.queue[1:]
	return v, nil
}

func (q *Queue) Front() (int, error) {
	if q.Empty() {
		return 0, fmt.Errorf("Front from empty queue")
	}
	return q.queue[0], nil
}

func (q *Queue) Back() (int, error) {
	if q.Empty() {
		return 0, fmt.Errorf("Back from empty queue")
	}
	return q.queue[len(q.queue)-1], nil
}

func (q *Queue) Empty() bool {
	return len(q.queue) == 0
}
