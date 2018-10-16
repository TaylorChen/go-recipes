package main

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	Web    = fakeSearch("web")
	Web1   = fakeSearch("web1")
	Image  = fakeSearch("image")
	Image1 = fakeSearch("image1")
	Video  = fakeSearch("video")
	Video1 = fakeSearch("Video1")
)

type Result string
type Search func(query string) string

func fakeSearch(kind string) Search {
	return func(query string) string {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		return fmt.Sprintf("%s result for %q\n", kind, query)
	}

}

func Google(query string) (results []string) {
	c := make(chan string)
	go func() { c <- First(query, Web, Web1) }()
	go func() { c <- First(query, Image, Image1) }()
	go func() { c <- First(query, Video, Video1) }()
	timeout := time.After(80 * time.Millisecond)
	for i := 0; i < 3; i++ {
		select {
		case result := <-c:
			results = append(results, result)
		case <-timeout:
			fmt.Println("timeout")
			return
		}
	}
	return
}

func First(query string, replicas ...Search) string {
	c := make(chan string)
	sarchReplica := func(i int) {
		c <- replicas[i](query)
	}
	for i := range replicas {
		go sarchReplica(i)
	}
	return <-c
}

func main() {
	rand.Seed(time.Now().UnixNano())
	start := time.Now()
	results := Google("hello")
	elapsed := time.Since(start)
	fmt.Println(results)
	fmt.Println(elapsed)
}
