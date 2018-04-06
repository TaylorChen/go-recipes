package main

/*
import (
	"fmt"
)
*/

/*
 *
 *
 * Given "abcabcbb", the answer is "abc", which the length is 3.
 * Given "bbbbb", the answer is "b", with the length of 1.
 * Given "pwwkew", the answer is "wke", with the length of 3. Note that the answer must be a substring, "pwke" is a subsequence and not a substring.
 *
 *
 */

func lengthOfLongestSubstring(s string) int {
	slength := len(s)
	ans := 0
	for i := 0; i < slength; i++ {
		for j := i + 1; j <= slength; j++ {
			if judgeUnique(s, i, j) {
				ans = max(ans, j-i)
			}
		}
	}
	return ans
}

func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}

func judgeUnique(s string, start int, end int) bool {
	chars := make(map[string]int)
	for i := start; i < end; i++ {
		_, ok := chars[string(s[i])]
		if ok {
			return false
		} else {
			chars[string(s[i])] = 1
		}
	}
	return true
}

/*
func main() {
	ans := lengthOfLongestSubstring("bbbbb")
	fmt.Println(ans)
	fmt.Println("vim-go")
}
*/
