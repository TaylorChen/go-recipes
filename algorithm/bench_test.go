package algorithm

import "testing"

func BenchmarkTwoSum(b *testing.B) {
	nums := make([]int, 1000)
	for i := range nums {
		nums[i] = i
	}
	for n := 0; n < b.N; n++ {
		_ = twoSum(nums, 1500)
	}
}
