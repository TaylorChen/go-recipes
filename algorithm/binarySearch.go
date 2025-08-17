package algorithm

func bsearch(arrs []int, len int, val int) int {
	low := 0
	high := len - 1
	for low <= high {
		mid := low + ((high - low) >> 1)
		if arrs[mid] == val {
			return mid
		} else if arrs[mid] < val {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func bsearchRecurise(arrs []int, low int, high int, val int) int {

	if low > high {
		return -1
	}

	mid := low + ((high - low) >> 1)
	if arrs[mid] == val {
		return mid
	} else if arrs[mid] < val {
		return bsearchRecurise(arrs, mid+1, high, val)
	} else {
		return bsearchRecurise(arrs, low, mid-1, val)
	}
}

func bsearchFirstPos(arrs []int, len int, val int) int {
	low := 0
	high := len - 1
	for low <= high {
		mid := low + ((high - low) >> 1)
		if arrs[mid] < val {
			low = mid + 1
		} else if arrs[mid] > val {
			high = mid - 1
		} else {
			if (mid == 0) || (arrs[mid-1] != val) {
				return mid
			} else {
				high = mid - 1
			}
		}
	}
	return -1
}

func bsearchLastPos(arrs []int, len int, val int) int {
	low := 0
	high := len - 1
	for low <= high {
		mid := low + ((high - low) >> 1)
		if arrs[mid] < val {
			low = mid + 1
		} else if arrs[mid] > val {
			high = mid - 1
		} else {
			if (mid == len-1) || (arrs[mid+1] != val) {
				return mid
			} else {
				low = mid + 1
			}
		}
	}
	return -1
}

func bsearchFirstGtePos(arrs []int, len int, val int) int {
	low := 0
	high := len - 1
	for low <= high {
		mid := low + ((high - low) >> 1)
		if arrs[mid] >= val {
			if (mid == 0) || (arrs[mid-1] < val) {
				return mid
			} else {
				high = mid - 1
			}
		} else {
			low = mid + 1
		}
	}
	return -1
}

// func main() {}
