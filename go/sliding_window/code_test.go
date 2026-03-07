package slidingwindow

import (
	"fmt"
	"math"
	"testing"
)

func TestMaxSumOfArrayWindow(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		k    int
		want int
	}{
		{"simple one", []int{1, 2, 3, 4, 5, 6}, 3, 15},
		{"empty", []int{}, 3, math.MinInt64},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MaxSumOfArrayWindow(tt.arr, tt.k)

			if got != tt.want {
				t.Errorf("TestMaxSumOfArrayWindow(%v, %d) = %d, want %d", tt.arr, tt.k, got, tt.want)
			}
		})
	}
}

func TestMaxVowels(t *testing.T) {
	tests := []struct {
		name string
		s    string
		k    int
		want int
	}{
		{"simple one", "abciiidef", 3, 3},
		{"no vowels", "bcdfg", 2, 0},
		{"all vowels", "aeiou", 3, 3},
		{"vowels at end", "bbbbaei", 3, 3},
		{"single char window", "aeiou", 1, 1},
		{"window equals length", "abc", 3, 1},
		{"single consonant", "b", 1, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MaxVowels(tt.s, tt.k)

			if got != tt.want {
				t.Errorf("TestMaxVowels(%s, %d) = %d, want %d", tt.s, tt.k, got, tt.want)
			}
		})
	}
}

func TestLongestSubNonRep(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{"simple", "abcabcbb", 3},
		{"all same", "bbbbb", 1},
		{"mixed", "pwwkew", 3},
		{"empty string", "", 0},
		{"single char", "a", 1},
		{"all unique", "abcdef", 6},
		{"repeat at end", "abcda", 4},
		{"repeat at start", "aab", 2},
		{"check previous", "abba", 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := LongestSubNonRep(tt.s)

			if got != tt.want {
				t.Errorf("LongestSubNonRep(%s) = %d, want %d", tt.s, got, tt.want)
			}
		})
	}
}

func TestSlidingWindowMax(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want []int
	}{
		{"simple", []int{1, 2, 3, 4, 5}, 3, []int{3, 4, 5}},
		{"with negative", []int{-1, 2, -3, 4, -5}, 3, []int{2, 4, 4}},
		{"descending", []int{5, 4, 3, 2, 1}, 3, []int{5, 4, 3}},
		{"all same", []int{3, 3, 3, 3}, 2, []int{3, 3, 3}},
		{"window equals length", []int{1, 2, 3}, 3, []int{3}},
		{"window of one", []int{4, 2, 7, 1}, 1, []int{4, 2, 7, 1}},
		{"max at front of window", []int{9, 1, 1, 1, 1}, 3, []int{9, 1, 1}},
		{"k greater than length", []int{1, 2}, 3, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SlidingWindowMax(tt.nums, tt.k)

			sGot := fmt.Sprintf("%v", got)
			sWant := fmt.Sprintf("%v", tt.want)

			if sGot != sWant {
				t.Errorf("TestSlidingWindowMax(%v, %d) = %v, want %v", tt.nums, tt.k, got, tt.want)
			}
		})
	}
}

func TestAtMostKDistinct(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want int
	}{
		// example from comment: [1,2,1,2,3], k=2 → 12
		{"example", []int{1, 2, 1, 2, 3}, 2, 12},
		// all elements identical → every subarray is valid: n*(n+1)/2
		{"all same", []int{5, 5, 5, 5}, 1, 10},
		{"all same k2", []int{5, 5, 5}, 2, 6},
		// k=0 → no distinct elements allowed → 0
		{"k zero", []int{1, 2, 3}, 0, 0},
		// empty array
		{"empty", []int{}, 3, 0},
		// single element
		{"single element", []int{7}, 1, 1},
		{"single element k0", []int{7}, 0, 0},
		// k >= number of distinct elements → all subarrays valid
		{"k covers all distinct", []int{1, 2, 3}, 3, 6},
		{"k exceeds distinct", []int{1, 2, 3}, 5, 6},
		// each element is unique, k=1 → only single-element subarrays
		{"all unique k1", []int{1, 2, 3, 4}, 1, 4},
		// two distinct elements interleaved
		{"alternating", []int{1, 2, 1, 2, 1}, 2, 15},
		// distinct count drops when window shrinks
		{"shrink window", []int{1, 2, 3, 1, 2}, 2, 9},
		// large k, small array
		{"large k", []int{1}, 100, 1},
		// negative numbers
		{"negatives", []int{-1, -2, -1}, 2, 6},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AtMostKDistinct(tt.nums, tt.k)

			if got != tt.want {
				t.Errorf("AtMostKDistinct(%v, %d) = %d, want %d", tt.nums, tt.k, got, tt.want)
			}
		})
	}
}

func TestSubarraysWithExactlyKDistinct(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want int
	}{
		// [1,2,1,2,3] k=2: atMost(2)=12, atMost(1)=5 → 7
		{"example", []int{1, 2, 1, 2, 3}, 2, 7},
		// all same elements, k=1: every subarray has exactly 1 distinct → n*(n+1)/2
		{"all same k1", []int{5, 5, 5, 5}, 1, 10},
		// all same elements, k=2: no subarray has 2 distinct → 0
		{"all same k2", []int{5, 5, 5}, 2, 0},
		// empty array
		{"empty", []int{}, 2, 0},
		// single element, k=1
		{"single k1", []int{7}, 1, 1},
		// single element, k=2: can't have 2 distinct with 1 element
		{"single k2", []int{7}, 2, 0},
		// k=0: no subarray can have 0 distinct (every non-empty subarray has ≥1)
		{"k zero", []int{1, 2, 3}, 0, 0},
		// all unique, k=1: only single-element subarrays
		{"all unique k1", []int{1, 2, 3, 4}, 1, 4},
		// all unique, k equals length: only 1 subarray (the whole array)
		{"all unique k=n", []int{1, 2, 3}, 3, 1},
		// [1,2,1,3,4], k=3: manually counted
		// exactly 3 distinct subarrays: [1,2,1,3], [2,1,3], [1,3,4] → 3
		{"mixed k3", []int{1, 2, 1, 3, 4}, 3, 3},
		// alternating two values, k=2: all subarrays have ≤2 distinct,
		// subtract those with exactly 1 → 15 - 5 = 10
		{"alternating k2", []int{1, 2, 1, 2, 1}, 2, 10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SubarraysWithExactlyKDistinct(tt.nums, tt.k)

			if got != tt.want {
				t.Errorf("SubarraysWithExactlyKDistinct(%v, %d) = %d, want %d", tt.nums, tt.k, got, tt.want)
			}
		})
	}
}

func TestMinSubarrayLenAtLeast(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		S     int
		want  int
	}{
		{"simple", []int{2, 3, 1, 2, 4, 3}, 7, 2},
		{"exact match", []int{1, 2, 3, 4, 5}, 15, 5},
		{"single element enough", []int{1, 2, 10, 3}, 10, 1},
		{"no valid subarray", []int{1, 1, 1}, 100, 0},
		{"empty", []int{}, 5, 0},
		{"s is zero", []int{1, 2, 3}, 0, 1},
		{"all ones", []int{1, 1, 1, 1, 1}, 3, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MinSubarrayLenAtLeast(tt.input, tt.S)

			if got != tt.want {
				t.Errorf("MinSubarryaLenAtLeast(%v, %d) = %d, want %d", tt.input, tt.S, got, tt.want)
			}
		})
	}
}

func TestFindAnagrams(t *testing.T) {
	tests := []struct {
		name string
		s    string
		p    string
		want []int
	}{
		// examples from comment
		{"example1", "cbaebabacd", "abc", []int{0, 6}},
		{"example2", "aaaa", "a", []int{0, 1, 2, 3}},
		// no anagrams found
		{"no match", "abcdef", "zz", nil},
		// empty pattern
		{"empty pattern", "abc", "", nil},
		// empty string
		{"empty string", "", "abc", nil},
		// both empty
		{"both empty", "", "", nil},
		// pattern longer than string
		{"pattern longer", "ab", "abcd", nil},
		// pattern equals string
		{"exact match", "abc", "abc", []int{0}},
		{"exact match scrambled", "bca", "abc", []int{0}},
		// single char repeated, pattern is that char repeated
		{"single char repeated", "aaaa", "aa", []int{0, 1, 2}},
		// overlapping anagrams
		{"overlapping", "abab", "ab", []int{0, 1, 2}},
		// pattern with repeated chars
		{"repeated in pattern", "abacbabc", "abc", []int{1, 2, 3, 5}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FindAnagrams(tt.s, tt.p)

			sGot := fmt.Sprintf("%v", got)
			sWant := fmt.Sprintf("%v", tt.want)

			if sGot != sWant {
				t.Errorf("FindAnagrams(%q, %q) = %v, want %v", tt.s, tt.p, got, tt.want)
			}
		})
	}
}
