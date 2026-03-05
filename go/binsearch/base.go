package binsearch

func Base(arr []int, target int) int {
	if len(arr) == 0 {
		return -1
	}

	left := 0
	right := len(arr) - 1

	for left <= right {
		mid := int(left + ((right - left) / 2))

		if arr[mid] == target {
			return mid
		}
		if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}

// find the first index i where a[i] >= target
func LowerBound(arr []int, target int) int {
	if len(arr) == 0 {
		return -1
	}

	left := 0
	right := len(arr) - 1

	result := -1
	for left <= right {
		mid := int(left + ((right - left) / 2))

		if arr[mid] >= target {
			result = mid
			right = mid - 1
		} else if arr[mid] < target {
			left = mid + 1
		}
	}

	return result
}

// find the first index i where a[i] > target
func UpperBound(arr []int, target int) int {
	if len(arr) == 0 {
		return -1
	}

	left := 0
	right := len(arr) - 1

	result := -1
	for left <= right {
		mid := int(left + ((right - left) / 2))

		if arr[mid] > target {
			result = mid
			right = mid - 1
		} else if arr[mid] <= target {
			left = mid + 1
		}
	}

	return result
}

func MonotonicPredicateFirstTrue(arr []bool) int {
	if len(arr) == 0 {
		return -1
	}

	left := 0
	right := len(arr) - 1

	result := -1

	for left <= right {
		mid := int(left + (right-left)/2)

		if arr[mid] == true {
			result = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return result
}

func MonotonicPredicateLastTrue(arr []bool) int {
	if len(arr) == 0 {
		return -1
	}

	left := 0
	right := len(arr) - 1

	result := -1

	for left <= right {
		mid := int(left + (right-left)/2)

		if arr[mid] == true {
			result = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return result
}

func FirstTrue(lo, hi int, ok func(int) bool) int {
	if !ok(hi) {
		return -1
	}

	left := lo
	right := hi

	result := -1
	for left <= right {
		mid := int(left + ((right - left) / 2))

		val := ok(mid)
		if val == true {
			result = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return result
}

func DoSqrt(num float64, precision float64) float64 {
	if num < 0 {
		return -1
	}

	if num == 1 {
		return 1
	}
	if num == 0 {
		return 0
	}

	var left float64
	var right float64

	if num > 1 {
		left = 1
		right = num
	} else {
		left = 0
		right = 1
	}

	for left <= right {
		mid := left + ((right - left) / 2)
		var doubled float64 = mid * mid
		diff := right - left
		// fmt.Println(left, right, doubled, diff, mid)

		if diff <= precision {
			return mid
		}

		if doubled > num {
			right = mid
		} else {
			left = mid
		}
	}

	return -1
}
