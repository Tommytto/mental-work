package twopointers

func TwoSum(nums []int, target int) []int {
	left, right := 0, len(nums)-1

	for left < right {
		sum := nums[left] + nums[right]

		if sum == target {
			return []int{left + 1, right + 1}
		}

		if sum < target {
			left++
		} else {
			right--
		}
	}

	return nil
}

func ContainerWithMostWater(height []int) int {
	if len(height) < 2 {
		return 0
	}

	best := 0
	left, right := 0, len(height)-1

	for left < right {
		h := min(height[left], height[right])
		area := h * (right - left)
		if area > best {
			best = area
		}

		if h == height[left] {
			left++
		} else {
			right--
		}
	}

	return best
}
