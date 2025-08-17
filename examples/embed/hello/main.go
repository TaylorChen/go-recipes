package main

import (
	_ "embed"
	"fmt"
)

//go:embed static/hello.txt
var hello []byte

func main() {
	fmt.Printf("embedded: %s\n", string(hello))
}
