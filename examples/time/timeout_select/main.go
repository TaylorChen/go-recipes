package main

import (
	"fmt"
	timepkg "time"
)

func main() {
	done := make(chan struct{})

	go func() {
		timepkg.Sleep(200 * timepkg.Millisecond)
		close(done)
	}()

	select {
	case <-done:
		fmt.Println("completed before timeout")
	case <-timepkg.After(100 * timepkg.Millisecond):
		fmt.Println("timed out")
	}
}
