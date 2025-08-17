package algorithm

import (
	"go-recipes/algorithm/collections"
)

func rightMaxArr(arr []int, len int) []int {
	res := make([]int, len)
	stack := collections.Stack{}
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

// func main() {}
