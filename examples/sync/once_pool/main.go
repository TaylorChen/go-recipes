package main

import (
	"bytes"
	"fmt"
	"sync"
)

var (
	initOnce  sync.Once
	globalCfg string
)

func initConfig() { globalCfg = "initialized-config" }

func main() {
	initOnce.Do(initConfig)
	fmt.Println("cfg:", globalCfg)

	var bufPool = sync.Pool{New: func() any { return new(bytes.Buffer) }}
	b := bufPool.Get().(*bytes.Buffer)
	b.Reset()
	b.WriteString("hello from pool")
	fmt.Println(b.String())
	bufPool.Put(b)
}
