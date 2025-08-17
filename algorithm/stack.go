package algorithm

import (
	"sync"
)

type StackItem struct {
	data string
}

type ItemStack struct {
	items []StackItem
	lock  sync.RWMutex
}

// create stack
func (s *ItemStack) New() *ItemStack {
	s.items = []StackItem{}
	return s
}

// push stack
func (s *ItemStack) Push(t StackItem) {
	s.lock.Lock()
	s.items = append(s.items, t)
	s.lock.Unlock()
}

// pop stack
func (s *ItemStack) Pop() *StackItem {
	s.lock.Lock()
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	s.lock.Unlock()
	return &item
}

// func main() {}
