package twopointers

import (
	"unicode"
)

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

// go left to right
// left pointer writes
// right pointer reads
func RemoveDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	l := 1
	for r := 1; r < len(nums); r++ {
		if nums[r] != nums[l-1] {
			nums[l] = nums[r]
			l++
		}
	}

	return l
}

// 5
// 0 1 2 3
func RemoveElementInPlace(nums []int, target int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	l := 0

	for r := 0; r < len(nums); r++ {
		if nums[r] != target {
			nums[l] = nums[r]
			l++
		}
	}

	return l
}

// given colors 0,1,2 = sort in place
func SortColors(nums []int) {
	swap := func(i, j int) {
		nums[i], nums[j] = nums[j], nums[i]
	}
	low, mid, high := 0, 0, len(nums)-1

	for mid <= high {
		midVal := nums[mid]
		if midVal == 0 {
			swap(low, mid)
			low++
			mid++
		} else if midVal == 1 {
			mid++
		} else if midVal == 2 {
			swap(mid, high)
			high--
		}
	}
}

func ValidPalindrome(s string) bool {
	rs := []rune(s)
	isAlphaNumeric := func(r rune) bool {
		return unicode.IsDigit(r) || unicode.IsLetter(r)
	}
	l, r := 0, len(rs)-1

	for l < r {
		if !isAlphaNumeric(rs[l]) {
			l++
			continue
		}
		if !isAlphaNumeric(rs[r]) {
			r--
			continue
		}
		if unicode.ToLower(rs[l]) == unicode.ToLower(rs[r]) {
			l++
			r--
			continue
		}

		return false
	}

	return true
}

func MoveZeroes(nums []int) {
	write := 0
	for read := 0; read < len(nums); read++ {
		if nums[read] != 0 {
			nums[write] = nums[read]
			write++
		}
	}

	for write < len(nums) {
		nums[write] = 0
		write++
	}
}

// Input:
// [-4,-1,0,3,10]

// Answer:
// [0,1,9,16,100]
func SquaresOfSortedArray(nums []int) []int {
	l, r := 0, len(nums)-1
	result := make([]int, len(nums))

	for i := len(nums) - 1; i >= 0; i-- {
		lSquare := nums[l] * nums[l]
		rSquare := nums[r] * nums[r]

		if lSquare > rSquare {
			result[i] = lSquare
			l++
		} else {
			result[i] = rSquare
			r--
		}
	}

	return result
}
