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
		c := boring("boring !")
		for i := 0; i < 10; i++ {
			fmt.Printf("You say:%q\n", <-c)
		}
	*/

	a := boring("a")
	b := boring("b")
	for i := 0; i < 10; i++ {
		fmt.Println(<-a)
		fmt.Println(<-b)
	}
	fmt.Println("You're boring; I'm leaving.")
}
