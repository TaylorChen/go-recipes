package main

import (
	. "./collections"
	"fmt"
)

func rightMaxArr(arr []int, len int) []int {
	res := make([]int, len)
	stack := Stack{}
	stack.Push(0)
	for i := 1; i < len; i++ {
		for stack.Len() != 0 && arr[i] > arr[(stack.Peek()).(int)] {
			res[(stack.Peek()).(int)] = arr[i]
			stack.Pop()
		}
		stack.Push(i)
	}

	return res
}

func main() {
	st := []int{1, 2, 3, 8, 7, 6}
	fmt.Println(rightMaxArr(st, len(st)))
}
