package main

import (
	"fmt"
)

func maxProfit(arrs []int) int {
	if len(arrs) < 2 {
		return 0
	}
	minPrice := arrs[0]
	maxDiff := arrs[1] - minPrice

	for i := 2; i < len(arrs); i++ {
		if arrs[i-1] < minPrice {
			minPrice = arrs[i-1]
		}

		currentDiff := arrs[i] - minPrice

		if currentDiff > maxDiff {
			maxDiff = currentDiff
		}
	}

	if maxDiff < 0 {
		return 0
	}

	return maxDiff
}

func main() {

	//arrs := []int{7, 1, 5, 4, 6, 4}
	arrs := []int{7, 6, 4, 3, 1}
	fmt.Println(maxProfit(arrs))

}
