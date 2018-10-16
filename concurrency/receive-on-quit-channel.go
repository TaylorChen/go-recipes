package main

import (
	"fmt"
	"math/rand"
	"time"
)

func boring(msg string, quit chan string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			select {
			case c <- fmt.Sprintf("%s %d", msg, i):
			case <-quit:
				quit <- "see you!"
				return
			}
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}

func main() {
	quit := make(chan string)
	c := boring("a", quit)
	rd := rand.Intn(30)
	fmt.Println(rd)
	for i := rd; i >= 0; i-- {
		fmt.Println(<-c)
	}
	quit <- "bye~"
	fmt.Println("a says:%q\n", <-quit)
}
