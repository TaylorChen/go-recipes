package algorithm

//import "fmt"

/*
 * 查找数组中连续2个数字之和等于target值得索引
 * input [1,3,5,2]  7
 * return [2,3]
 *
 */
func twoSum(nums []int, target int) []int {
	var arr []int
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if target-nums[i] == nums[j] {
				arr = []int{i, j}
				goto end
			}
		}
	}
end:
	return arr
}

//func main() {
//	arr := []int{6, 3, 5, 7, 4, 4}
//	arrRst := twoSum(arr, 8)
//	for i := 0; i < len(arrRst); i++ {
//		fmt.Println(arrRst[i])
//	}
//}
