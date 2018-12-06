package main

import (
	"fmt"
	"math/rand"
	"time"
)

func boring(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}

func main() {

	/*
		c := boring("Joe")
		for {
			select {
			case s := <-c:
				fmt.Println(s)
			case <-time.After(100 * time.Millisecond):
				fmt.Println("You are to slow")
			}
		}
	*/

	c := boring("Joe")
	timeout := time.After(100 * time.Millisecond)
	for {
		select {
		case s := <-c:
			fmt.Println(s)
		case <-timeout:
			fmt.Println("You are to slow")
		}
	}

	fmt.Println("You're boring; I'm leaving.")
}
