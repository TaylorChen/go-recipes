package algorithm

import (
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	actual := lengthOfLongestSubstring("bbbbb")
	if actual != 1 {
		t.Errorf("Expected the value of %d but actual got %d!", 1, actual)
	}
}

func TestMax(t *testing.T) {
	actual := max(1, 3)
	if actual != 3 {
		t.Errorf("Expected the value of %d but actual got %d!", 3, actual)
	}
}
