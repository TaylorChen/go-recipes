package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var counter int64
	var m sync.Map

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				atomic.AddInt64(&counter, 1)
				m.Store(fmt.Sprintf("k-%d-%d", id, j), j)
			}
		}(i)
	}
	wg.Wait()
	fmt.Println("counter:", counter)

	size := 0
	m.Range(func(_, _ any) bool { size++; return true })
	fmt.Println("map size:", size)
}
