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

// Daily Temperatures
// temps = [73,74,75,71,69,72,76,73]
// ans   = [1,1,4,2,1,1,0,0]
func DailyTemperatures(temp []int) []int {
	n := len(temp)
	out := make([]int, n)
	stack := make([]int, 0, n)

	for idx := n - 1; idx >= 0; idx-- {
		for len(stack) > 0 && temp[stack[len(stack)-1]] <= temp[idx] {
			stack = stack[:len(stack)-1]
		}

		if len(stack) == 0 {
			out[idx] = 0
		} else {
			out[idx] = stack[len(stack)-1] - idx
		}

		stack = append(stack, idx)
	}

	return out
}

// Largest Rectangle in Histogram
// heights = [2,1,5,6,2,3]
// answer = 10
func LargestRectangleHistogram(heights []int) int {
	n := len(heights)
	if n == 1 {
		return heights[0]
	}
	answer := 0

	prevSmaller := make([]int, n)
	nextSmaller := make([]int, n)

	stack := make([]int, 0, n)
	for idx, h := range heights {
		for len(stack) > 0 && heights[stack[len(stack)-1]] >= h {
			stack = stack[:len(stack)-1]
		}

		if len(stack) == 0 {
			prevSmaller[idx] = -1
		} else {
			prevSmaller[idx] = stack[len(stack)-1]
		}

		stack = append(stack, idx)
	}

	stack = make([]int, 0, n)
	for idx := n - 1; idx >= 0; idx-- {
		for len(stack) > 0 && heights[stack[len(stack)-1]] >= heights[idx] {
			stack = stack[:len(stack)-1]
		}

		if len(stack) == 0 {
			nextSmaller[idx] = n
		} else {
			nextSmaller[idx] = stack[len(stack)-1]
		}

		stack = append(stack, idx)
	}

	for idx := 0; idx < len(heights); idx++ {
		r := nextSmaller[idx]
		l := prevSmaller[idx]

		width := r - l - 1
		area := heights[idx] * width

		if area > answer {
			answer = area
		}
	}

	return answer
}

// Largest Rectangle in Histogram
// heights = [2,1,5,6,2,3]
// answer = 10
func LargestRectangleHistogramOnePass(heights []int) int {
	n := len(heights)

	best := 0
	stack := make([]int, 0, n)

	for i := 0; i <= n; i++ {
		currHeight := 0
		if i < n {
			currHeight = heights[i]
		}

		for len(stack) > 0 && heights[stack[len(stack)-1]] >= currHeight {
			mid := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			left := -1
			if len(stack) > 0 {
				left = stack[len(stack)-1]
			}

			width := i - left - 1
			height := heights[mid]

			area := width * height
			if area > best {
				best = area
			}
		}

		stack = append(stack, i)
	}

	return best
}

// Largest Rectangle in Histogram
// heights = [2,1,5,6,2,3]
// answer = 10
func LargestRectangle2(nums []int) int {
	n := len(nums)

	// find next smaller
	// find prev smaller
	// after we know borders, calc rectangle
	nextSmaller := make([]int, n)
	prevSmaller := make([]int, n)

	stack := make([]int, 0, n)
	best := 0

	for i := 0; i < n; i++ {
		for len(stack) > 0 && nums[stack[len(stack)-1]] >= nums[i] {
			stack = stack[:len(stack)-1]
		}

		if len(stack) == 0 {
			prevSmaller[i] = -1
		} else {
			prevSmaller[i] = stack[len(stack)-1]

		}
		stack = append(stack, i)
	}

	stack = make([]int, 0, n)
	for i := n - 1; i >= 0; i-- {
		for len(stack) > 0 && nums[stack[len(stack)-1]] >= nums[i] {
			stack = stack[:len(stack)-1]
		}

		if len(stack) == 0 {
			nextSmaller[i] = n
		} else {
			nextSmaller[i] = stack[len(stack)-1]

		}
		stack = append(stack, i)
	}

	for i := 0; i < n; i++ {
		height := nums[i]
		width := nextSmaller[i] - prevSmaller[i] - 1
		area := width * height
		if area > best {
			best = area
		}
	}

	return best
}

// Largest Rectangle in Histogram
// heights = [2,1,5,6,2,3]
// answer = 10
func LargestRectangleOnePass2(nums []int) int {
	n := len(nums)
	best := 0

	stack := make([]int, 0, n)
	for i := 0; i <= n; i++ {
		currentHeight := -1
		if i < n {
			currentHeight = nums[i]
		}

		for len(stack) > 0 && nums[stack[len(stack)-1]] > currentHeight {
			mid := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			r := i

			l := -1
			if len(stack) > 0 {
				l = stack[len(stack)-1]
			}

			width := r - l - 1
			height := nums[mid]
			area := width * height

			if area > best {
				best = area
			}
		}

		stack = append(stack, i)
	}

	return best
}
