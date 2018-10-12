package main

import (
	"fmt"
)

func bsearch(arrs []int, len int, val int) int {
	low := 0
	high := len - 1
	for low <= high {
		mid := low + (high-low)>>1
		if arrs[mid] == val {
			return mid
		} else if arrs[mid] < val {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func bsearchRecurise(arrs []int, low int, high int, val int) int {

	if low > high {
		return -1
	}

	mid := low + (high-low)>>1
	if arrs[mid] == val {
		return mid
	} else if arrs[mid] < val {
		return bsearchRecurise(arrs, mid+1, high, val)
	} else {
		return bsearchRecurise(arrs, low, mid-1, val)
	}
}

func main() {
	arrs := []int{1, 100, 12, 8, 34, 56, 90}
	fmt.Println(bsearch(arrs, 7, 90))
	fmt.Println(bsearchRecurise(arrs, 0, 6, 90))
}
