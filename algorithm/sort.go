package main

import (
	"fmt"
)

func bubbleSort(arrs []int, len int) []int {

	if len <= 0 {
		fmt.Println("bad length")
		return arrs
	}

	for i := 0; i < len; i++ {
		tag := false
		for j := 0; j < len-1-i; j++ {
			if arrs[j] > arrs[j+1] {
				tmp := arrs[j]
				arrs[j] = arrs[j+1]
				arrs[j+1] = tmp
				tag = true
			}
		}
		if !tag {
			break
		}
	}
	return arrs
}

func insertSort(arrs []int, len int) []int {

	if len <= 0 {
		fmt.Println("bad length")
		return arrs
	}

	for i := 1; i < len; i++ {
		val := arrs[i]
		j := i - 1
		for ; j >= 0; j-- {
			if arrs[j] > val {
				arrs[j+1] = arrs[j]
			} else {
				break
			}
		}
		arrs[j+1] = val
	}

	return arrs
}

func selectSort(arrs []int, len int) []int {

	if len <= 0 {
		fmt.Println("bad length")
		return arrs
	}

	for i := 0; i < len; i++ {
		min := i
		for j := i; j < len; j++ {
			if arrs[j] < arrs[min] {
				min = j
			}
		}
		if min != i {
			tmp := arrs[i]
			arrs[i] = arrs[min]
			arrs[min] = tmp
		}

	}
	return arrs
}

func quickSort(arrs []int, low int, high int) {
	if low >= high {
		return
	}
	pivot := arrs[low]
	i := low
	j := high
	for i < j {
		for i < j && arrs[j] >= pivot {
			j--
		}
		arrs[i] = arrs[j]

		for i < j && arrs[i] <= pivot {
			i++
		}
		arrs[j] = arrs[i]

	}
	arrs[i] = pivot
	quickSort(arrs, low, i-1)
	quickSort(arrs, i+1, high)
}

func main() {
	fmt.Println(bubbleSort([]int{5, 10, 3, 2, 1, -1, -2}, 7))
	//fmt.Println(insertSort([]int{5, 4, 3}, 3))
	//fmt.Println(selectSort([]int{5, -1, 10, 1}, 4))
	//arrs := []int{12, -1000, 345, 5, -1, 10, 1}
	//quickSort(arrs, 0, 3)
	//fmt.Println(arrs)
}
