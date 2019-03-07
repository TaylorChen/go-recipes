package main

import (
	"fmt"
	"github.com/robfig/cron"
	"time"
)

type Job struct {
}

func (this TestJob) Run() {
	fmt.Println("testJob1...")
}

type Test2Job struct {
}

func (this Test2Job) Run() {
	fmt.Println("testJob2...")
}

func main() {
	i := 0
	j := 0
	cron := cron.New()
	cron.Start()
	defer cron.Stop()
	cron.AddFunc("* * * * * *", func() {
		i++
		fmt.Println("cron runing a !!", i)
	})

	cron.AddFunc("* * * * * *", func() {
		j++
		fmt.Println("cron runing b !!", j)
	})
	select {
	case <-time.After(60 * time.Second):
		fmt.Println("run done")
		return
	}
}
