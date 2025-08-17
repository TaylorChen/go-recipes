package algorithm

import "fmt"

func findThePivotElement(arr []int, len int) {
	rightMin := make([]int, len)
	r_min := arr[len-1]
	for i := len - 2; i >= 0; i-- {
		if arr[i] < r_min {
			r_min = arr[i]
		}
		rightMin[i] = r_min
	}

	l_max := arr[0]
	for i := 1; i < len-1; i++ {
		if arr[i] > l_max {
			l_max = arr[i]
			if arr[i] < rightMin[i+1] {
				fmt.Println(arr[i])
			}
		}
	}

}

// func main() {}
