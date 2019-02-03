package main

import (
	"fmt"
	"sync"
)

type Item struct {
	data string
}

type ItemStack struct {
	items []Item
	lock  sync.RWMutex
}

// create stack
func (s *ItemStack) New() *ItemStack {
	s.items = []Item{}
	return s
}

//push stack
func (s *ItemStack) Push(t Item) {
	s.lock.Lock()
	s.items = append(s.items, t)
	s.lock.Unlock()
}

//pop stack
func (s *ItemStack) Pop() *Item {
	s.lock.Lock()
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	s.lock.Unlock()
	return &item
}

func main() {
	istack := ItemStack{}
	istack.New()
	istack.Push(Item{data: "a"})
	istack.Push(Item{data: "b"})
	istack.Push(Item{data: "c"})
	tmpItem := istack.Pop()
	fmt.Println(tmpItem.data)
}
