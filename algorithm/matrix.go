package main

import (
	"fmt"
)

func Find(arrs [][]int, rows, columns, num int) bool {
	isFind := false
	if len(arrs) != 0 && rows > 0 && columns > 0 {
		row := 0
		column := columns - 1

		for row < rows && column >= 0 {
			if arrs[row][column] == num {
				isFind = true
				break
			} else if arrs[row][column] > num {
				column--
			} else {
				row++
			}
		}
	}
	return isFind
}

func main() {
	martix := [][]int{
		{1, 2, 3},
		{2, 4, 6},
		{3, 8, 9},
	}
	fmt.Println(Find(martix, 3, 3, 8))
}
