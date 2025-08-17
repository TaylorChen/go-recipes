package algorithm

// import "fmt"

func bubbleSort(arrs []int, len int) {
	for i := 0; i < len; i++ {
		tag := false
		for j := 0; j < len-1-i; j++ {
			if arrs[j] > arrs[j+1] {
				tmp := arrs[j]
				arrs[j] = arrs[j+1]
				arrs[j+1] = tmp
				tag = true
			}
		}
		if !tag {
			break
		}
	}
}

func insertSort(arrs []int, len int) {
	for i := 1; i < len; i++ {
		val := arrs[i]
		j := i - 1
		for ; j >= 0; j-- {
			if arrs[j] > val {
				arrs[j+1] = arrs[j]
			} else {
				break
			}
		}
		arrs[j+1] = val
	}
}

func selectSort(arrs []int, len int) {
	for i := 0; i < len; i++ {
		min := i
		for j := i; j < len; j++ {
			if arrs[j] < arrs[min] {
				min = j
			}
		}
		if min != i {
			tmp := arrs[i]
			arrs[i] = arrs[min]
			arrs[min] = tmp
		}

	}
}

func quickSort(arrs []int, low int, high int) {
	if low >= high {
		return
	}
	pivot := arrs[low]
	i := low
	j := high
	for i < j {
		for i < j && arrs[j] >= pivot {
			j--
		}
		arrs[i] = arrs[j]

		for i < j && arrs[i] <= pivot {
			i++
		}
		arrs[j] = arrs[i]

	}
	arrs[i] = pivot
	quickSort(arrs, low, i-1)
	quickSort(arrs, i+1, high)
}

// func main() {}
