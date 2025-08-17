package algorithm

import (
	"sync"
)

type QueueItem struct {
	data string
}

type ItemQueue struct {
	items []QueueItem
	lock  sync.RWMutex
}

// create stack
func (q *ItemQueue) New() *ItemQueue {
	q.items = []QueueItem{}
	return q
}

// En queue
func (q *ItemQueue) Enqueue(t QueueItem) {
	q.lock.Lock()
	q.items = append(q.items, t)
	q.lock.Unlock()
}

// De queue
func (q *ItemQueue) Dequeue() *QueueItem {
	q.lock.Lock()
	item := q.items[0]
	q.items = q.items[1:len(q.items)]
	q.lock.Unlock()
	return &item
}

// front item
func (q *ItemQueue) Front() *QueueItem {
	q.lock.Lock()
	item := q.items[0]
	q.lock.Unlock()
	return &item
}

func (q *ItemQueue) IsEmpty() bool {
	return len(q.items) == 0
}

// size
func (q *ItemQueue) Size() int {
	return len(q.items)
}

// func main() {}
