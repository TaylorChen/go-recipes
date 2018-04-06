package main

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

func TestJudgeUnique(t *testing.T) {
	actual := judgeUnique("abc", 0, 2)
	if !actual {
		t.Errorf("Expected the value of %t but actual got %t!", true, actual)
	}
}
