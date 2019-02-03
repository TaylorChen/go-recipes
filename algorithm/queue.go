package main

import (
	"fmt"
	"sync"
)

type Item struct {
	data string
}

type ItemQueue struct {
	items []Item
	lock  sync.RWMutex
}

// create stack
func (q *ItemQueue) New() *ItemQueue {
	q.items = []Item{}
	return q
}

//En queue
func (q *ItemQueue) Enqueue(t Item) {
	q.lock.Lock()
	q.items = append(q.items, t)
	q.lock.Unlock()
}

//De queue
func (q *ItemQueue) Dequeue() *Item {
	q.lock.Lock()
	item := q.items[0]
	q.items = q.items[1:len(q.items)]
	q.lock.Unlock()
	return &item
}

//front item
func (q *ItemQueue) Front() *Item {
	q.lock.Lock()
	item := q.items[0]
	q.lock.Unlock()
	return &item
}

func (q *ItemQueue) IsEmpty() bool {
	return len(q.items) == 0
}

//size
func (q *ItemQueue) Size() int {
	return len(q.items)
}

func main() {
	queue := ItemQueue{}
	queue.New()
	queue.Enqueue(Item{data: "a"})
	queue.Enqueue(Item{data: "b"})
	tmpItem := queue.Dequeue()
	fmt.Println(tmpItem.data)
	fmt.Println(queue.items)
}
