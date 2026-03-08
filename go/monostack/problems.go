package monostack

// for every num in nums return biggest num from the right side, otherwise -1
// nums = [2, 1, 2, 4, 3]
// ans  = [4, 2, 4, -1, -1]
func NextGreaterElement(nums []int) []int {
	n := len(nums)

	stack := make([]int, 0, n)
	result := make([]int, n)

	for idx := range result {
		result[idx] = -1
	}

	for idx, num := range nums {
		for len(stack) > 0 && nums[stack[len(stack)-1]] < num {
			result[stack[len(stack)-1]] = num
			stack = stack[:len(stack)-1]
		}

		stack = append(stack, idx)
	}

	return result
}

// nums = [4, 5, 2, 10, 8]
// ans  = [-1, 4, -1, 2, 2]
func PreviousSmallerElement(nums []int) []int {
	n := len(nums)
	stack := make([]int, 0, n)
	result := make([]int, n)
	for idx := range result {
		result[idx] = -1
	}

	for idx := n - 1; idx >= 0; idx-- {
		num := nums[idx]
		for len(stack) > 0 && nums[stack[len(stack)-1]] > num {
			result[stack[len(stack)-1]] = num
			stack = stack[:len(stack)-1]
		}

		stack = append(stack, idx)
	}

	return result
}

// prices := []int{100, 80, 60, 70, 60, 75, 85}
// answer := []int{1, 1, 1, 2, 1, 4, 6}
func StockSpan(prices []int) []int {
	n := len(prices)
	result := make([]int, n)

	stack := make([]int, 0, n)

	for idx, price := range prices {
		for len(stack) > 0 && prices[stack[len(stack)-1]] <= price {
			stack = stack[:len(stack)-1]
		}

		if len(stack) == 0 {
			result[idx] = idx + 1
		} else {
			result[idx] = idx - stack[len(stack)-1]
		}

		stack = append(stack, idx)
	}

	return result
}
