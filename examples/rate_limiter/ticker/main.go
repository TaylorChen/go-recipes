package main

import (
	"fmt"
	"time"
)

func main() {
	limiter := time.NewTicker(100 * time.Millisecond)
	defer limiter.Stop()

	for i := 1; i <= 5; i++ {
		<-limiter.C
		fmt.Println("request", i, "at", time.Now().Format(time.RFC3339Nano))
	}
}
