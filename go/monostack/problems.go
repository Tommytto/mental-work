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
