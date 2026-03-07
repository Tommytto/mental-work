package slidingwindow

import (
	"math"
)

// maximum sum of any subarray of length K
func MaxSumOfArrayWindow(arr []int, K int) int {
	sum := 0
	best := math.MinInt64

	n := len(arr)
	l := 0

	for r := range n {
		size := r - l + 1
		sum += arr[r]

		if size == K {
			if sum > best {
				best = sum
			}
			sum -= arr[l]
			l += 1
		}
	}

	return best
}

// Maximum number of vowels in any substring of length K
func MaxVowels(s string, K int) int {
	n := len(s)

	l := 0

	isVowel := func(c byte) bool {
		switch c {
		case 'a', 'e', 'i', 'o', 'u':
			return true
		}
		return false
	}

	sum := 0
	best := 0
	for r := range n {
		if isVowel(s[r]) {
			sum += 1
		}

		if r-l+1 == K {
			if sum > best {
				best = sum
			}

			if isVowel(s[l]) {
				sum -= 1
			}
			l += 1
		}
	}

	return best
}

// Longest substring without repeating characters
// sliding window
// remember what letters on what position were
// when duplicate found move left cursor to next position
func LongestSubNonRep(s string) int {
	runes := []rune(s)
	seen := make(map[rune]int)

	l := 0
	best := 0

	for r, ch := range runes {
		if pos, ok := seen[ch]; ok && pos >= l {
			l = pos + 1
		}

		currLen := r - l + 1
		if currLen > best {
			best = currLen
		}

		seen[ch] = r
	}

	return best
}

// find max value in window K and return all such values for all windows
func SlidingWindowMax(nums []int, k int) []int {
	n := len(nums)
	if k <= 0 || k > n {
		return nil
	}

	dq := make([]int, 0, k)
	out := make([]int, 0, n-k+1)

	l := 0

	for r := range n {
		for len(dq) > 0 && nums[dq[len(dq)-1]] <= nums[r] {
			dq = dq[:len(dq)-1]
		}
		dq = append(dq, r)

		if dq[0] < l {
			dq = dq[1:]
		}

		if r-l+1 == k {
			out = append(out, nums[dq[0]])
			l++
		}
	}

	return out
}

// minimal subarray length, where items sum >= S
func MinSubarrayLenAtLeast(nums []int, S int) int {
	n := len(nums)

	if n == 0 {
		return 0
	}

	if S == 0 {
		return 1
	}

	l := 0
	sum := 0
	bestLength := n + 1

	for r := range n {
		sum += nums[r]

		for sum >= S {
			length := r - l + 1
			if length < bestLength {
				bestLength = length
			}
			sum -= nums[l]
			l++
		}
	}

	if bestLength == n+1 {
		return 0
	}

	return bestLength
}

// Subarrays with at most K distinct integers
// edge cases
// eg. nums = [1,2,1,2,3], k = 2 → answer 12
func AtMostKDistinct(nums []int, k int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if k <= 0 {
		return 0
	}
	counter := make(map[int]int)

	l := 0
	distinct := 0
	subArraysCount := 0

	for r := range n {
		if count := counter[nums[r]]; count == 0 {
			distinct += 1
			counter[nums[r]] = 1
		} else {
			counter[nums[r]] += 1
		}

		for distinct > k {
			counter[nums[l]] -= 1
			if count := counter[nums[l]]; count == 0 {
				distinct -= 1
			}

			l++
		}

		subArraysCount += r - l + 1
	}

	return subArraysCount
}

func SubarraysWithExactlyKDistinct(nums []int, k int) int {
	return AtMostKDistinct(nums, k) - AtMostKDistinct(nums, k-1)
}

// return list of starting indices of start of anagrams
// s = "cbaebabacd", p = "abc" → [0, 6]
// s = "cbaebabacd", p = "abb" → [4, 5]
// s = "aaaa", p = "a" -> [0,1,2,3]
func FindAnagrams(s, p string) []int {
	sr := []rune(s)
	pr := []rune(p)

	n, m := len(s), len(p)
	if m == 0 || m > n {
		return nil
	}

	need := make(map[rune]int)
	for _, ch := range pr {
		need[ch]++
	}
	required := len(need)

	have := make(map[rune]int)
	matched := 0

	var res []int

	l := 0
	for r, ch := range sr {
		if _, ok := need[ch]; ok {
			have[ch]++
			if have[ch] == need[ch] {
				matched++
			} else if have[ch] == need[ch]+1 {
				matched--
			}
		}

		if r-l+1 > m {
			left := sr[l]
			if _, ok := need[left]; ok {
				if have[left] == need[left] {
					matched--
				} else if have[left] == need[left]+1 {
					matched++
				}
				have[left]--
				if have[left] == 0 {
					delete(have, left)
				}
			}
			l++
		}

		if r-l+1 == m && matched == required {
			res = append(res, l)
		}
	}

	return res
}
