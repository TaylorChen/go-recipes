package main

import (
	"fmt"
	"strings"
)

var space = [][2]uint16{
	{0x0009, 0x000d},
	{0x0020, 0x0020},
	{0x0085, 0x0085},
	{0x00a0, 0x00a0},
	{0x1680, 0x1680},
	{0x2000, 0x200a},
	{0x2028, 0x2029},
	{0x202f, 0x202f},
	{0x205f, 0x205f},
	{0x3000, 0x3000},
}

func isSpace(r rune) bool {
	if r >= 1<<16 {
		return false
	}
	rx := uint16(r)
	for _, rng := range space {
		if rx < rng[0] {
			return false
		}
		if rx <= rng[1] {
			return true
		}
	}
	return false
}

func notSpace(r rune) bool {
	return !isSpace(r)
}

func test() {
	a := strings.Replace("中国 上海", " ", "@", -1)
	fmt.Println(a)
	a = strings.Replace("中国\u00a0上海", " ", "@", -1)
	fmt.Println(a)
	a = strings.Replace("中国\u0020上海", " ", "@", -1)
	fmt.Println(a)
	a = strings.Replace("中国\u3000上海", " ", "@", -1)
	fmt.Println(a)
}

func main() {
	test()
	// a := "中国\u00a0上海"
	// a := "中国\u0020上海"
	a := "中国\u3000上海"
	var r = []rune(a)
	for _, r := range r {
		fmt.Println(isSpace(r))
	}
}
