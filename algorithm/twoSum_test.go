package main

import (
	"testing"
)

func TestTwoSum(t *testing.T) {
	arr := []int{6, 3, 5, 7, 4, 4}
	target := 8
	actual := twoSum(arr, target)

	for _, val := range actual {
		if !(val == 1 || val == 2) {
			t.Errorf("Expected the index of %d , %d but actual got %d!", 1, 2, val)
		}
	}
}
