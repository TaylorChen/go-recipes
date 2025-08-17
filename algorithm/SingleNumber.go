package algorithm

// import "fmt"

func SingleNumber(arrs []int) int {
	ret := 0
	for _, num := range arrs {
		ret ^= num
	}
	return ret
}

// func main() {}
