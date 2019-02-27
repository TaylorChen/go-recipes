package main

import (
	"fmt"
)

func SingleNumber(arrs []int) int {
	ret := 0
	for _, num := range arrs {
		ret ^= num
	}
	return ret
}

func main() {
	arrs := []int{2, 2, 1, 4, 5, 4, 6, 5, 6}
	ret := SingleNumber(arrs)
	fmt.Println(ret)
}
