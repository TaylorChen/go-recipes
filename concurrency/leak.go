package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

type result struct {
	record string
	err    error
}

func search(term string) (string, error) {
	time.Sleep(200 * time.Millisecond)
	return term, nil
}

func process(term string) error {

	// 创建一个在 100 ms 内取消上下文的 context
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	// 为 Goroutine 创建一个传递结果的 channel
	ch := make(chan result)

	// 启动一个 Goroutine 来寻找记录，然后得到结果
	// 将返回值从 channel 中返回
	go func() {
		record, err := search(term)
		ch <- result{record, err}
	}()

	// 阻塞等待从 Goroutine 接收值
	// 通过 channel 和 context 来取消上下文操作
	select {
	case <-ctx.Done():
		return errors.New("search canceled")
	case result := <-ch:
		fmt.Println("11111")
		if result.err != nil {
			return result.err
		}
		fmt.Println("Received:", result.record)
		return nil
	}
}

func main() {
	fmt.Println(process("hello world"))
}
