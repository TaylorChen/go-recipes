package main

import (
	"fmt"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	size, i, j := len(nums1)+len(nums2), 0, 0
	slice := make([]int, size, size)
	for k := 0; k < size; k++ {
		if i > len(nums1)-1 && j <= len(nums2)-1 {
			slice[k] = nums2[j]
			j++
		} else if i <= len(nums1)-1 && j > len(nums2)-1 {
			slice[k] = nums1[i]
			i++
		} else if nums1[i] < nums2[j] {
			slice[k] = nums1[i]
			i++
		} else {
			slice[k] = nums2[j]
			j++
		}
	}
	cnt := len(slice)
	ret := 0.0
	if cnt%2 == 0 {
		ret = float64(slice[(cnt/2)]+slice[(cnt/2)-1]) / 2
	} else {
		ret = float64(slice[cnt/2])
	}
	return ret
}

func main() {
	nums1 := []int{1, 2}
	nums2 := []int{3, 4}
	median := findMedianSortedArrays(nums1, nums2)
	fmt.Println(median)
}
