package main

import (
	"fmt"
	"strings"
)

func Map[T any, R any](in []T, fn func(T) R) []R {
	out := make([]R, len(in))
	for i, v := range in {
		out[i] = fn(v)
	}
	return out
}

func Filter[T any](in []T, fn func(T) bool) []T {
	out := make([]T, 0, len(in))
	for _, v := range in {
		if fn(v) {
			out = append(out, v)
		}
	}
	return out
}

func main() {
	words := []string{"go", "recipes", "are", "great"}
	upper := Map(words, strings.ToUpper)
	long := Filter(upper, func(s string) bool { return len(s) > 3 })
	fmt.Println("upper:", upper)
	fmt.Println("long:", long)
}
