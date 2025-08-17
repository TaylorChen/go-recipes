package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context, id int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("worker", id, "stopped:", ctx.Err())
			return
		case <-time.After(200 * time.Millisecond):
			fmt.Println("worker", id, "tick")
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 800*time.Millisecond)
	defer cancel()
	go worker(ctx, 1)
	go worker(ctx, 2)
	time.Sleep(1 * time.Second)
}
