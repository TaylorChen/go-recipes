package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Job struct{ ID int }
type Result struct {
	ID      int
	Latency time.Duration
}

func worker(id int, jobs <-chan Job, results chan<- Result) {
	for j := range jobs {
		d := time.Duration(50+rand.Intn(150)) * time.Millisecond
		time.Sleep(d)
		results <- Result{ID: j.ID, Latency: d}
		fmt.Println("worker", id, "done job", j.ID)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	jobs := make(chan Job)
	results := make(chan Result)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	go func() {
		for i := 1; i <= 5; i++ {
			jobs <- Job{ID: i}
		}
		close(jobs)
	}()

	for i := 0; i < 5; i++ {
		r := <-results
		fmt.Println("result: job", r.ID, "latency", r.Latency)
	}
}
