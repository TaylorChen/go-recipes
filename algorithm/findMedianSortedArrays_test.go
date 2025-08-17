package algorithm

import (
	"testing"
)

func TestFindMedianSortedArrays(t *testing.T) {
	nums1 := []int{1, 2}
	nums2 := []int{3, 4}
	median := findMedianSortedArrays(nums1, nums2)
	if median != 2.5 {
		t.Errorf("Expected the %f ,but actual got %f!", 2.5, median)
	}
}
