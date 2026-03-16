package monodeque

import (
	"slices"
	"testing"
)

func TestSlidingWindowMax(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want []int
	}{
		{"example", []int{1, 3, -1, -3, 5, 3, 6, 7}, 3, []int{3, 3, 5, 5, 6, 7}},
		{"k equals len", []int{1, 2, 3}, 3, []int{3}},
		{"k is 1", []int{4, 2, 5, 1}, 1, []int{4, 2, 5, 1}},
		{"single element", []int{7}, 1, []int{7}},
		{"descending", []int{5, 4, 3, 2, 1}, 2, []int{5, 4, 3, 2}},
		{"ascending", []int{1, 2, 3, 4, 5}, 2, []int{2, 3, 4, 5}},
		{"all equal", []int{3, 3, 3, 3}, 2, []int{3, 3, 3}},
		{"negative values", []int{-1, -3, -5, -2}, 2, []int{-1, -3, -2}},
		{"peak in window", []int{1, 5, 1, 1, 1}, 3, []int{5, 5, 1}},
		{"k is 2 alternating", []int{1, 3, 1, 3, 1}, 2, []int{3, 3, 3, 3}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SlidingWindowMax(tt.nums, tt.k)
			if !slices.Equal(got, tt.want) {
				t.Errorf("SlidingWindowMax(%v, %d) = %v, want %v", tt.nums, tt.k, got, tt.want)
			}
		})
	}
}

func TestSlidingWindowMin(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want []int
	}{
		{"example", []int{1, 3, -1, -3, 5, 3, 6, 7}, 3, []int{-1, -3, -3, -3, 3, 3}},
		{"k equals len", []int{3, 1, 2}, 3, []int{1}},
		{"k is 1", []int{4, 2, 5, 1}, 1, []int{4, 2, 5, 1}},
		{"single element", []int{7}, 1, []int{7}},
		{"descending", []int{5, 4, 3, 2, 1}, 2, []int{4, 3, 2, 1}},
		{"ascending", []int{1, 2, 3, 4, 5}, 2, []int{1, 2, 3, 4}},
		{"all equal", []int{3, 3, 3, 3}, 2, []int{3, 3, 3}},
		{"negative values", []int{-1, -3, -5, -2}, 2, []int{-3, -5, -5}},
		{"valley in window", []int{5, 1, 5, 5, 5}, 3, []int{1, 1, 5}},
		{"k is 2 alternating", []int{3, 1, 3, 1, 3}, 2, []int{1, 1, 1, 1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SlidingWindowMin(tt.nums, tt.k)
			if !slices.Equal(got, tt.want) {
				t.Errorf("SlidingWindowMin(%v, %d) = %v, want %v", tt.nums, tt.k, got, tt.want)
			}
		})
	}
}

func TestShortestSubarrayAtLeastK(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want int
	}{
		{"single match", []int{1}, 1, 1},
		{"no match", []int{1, 2}, 4, -1},
		{"with negatives", []int{2, -1, 2}, 3, 3},
		{"exact element", []int{1, 2, 3}, 3, 1},
		{"from comments", []int{84, -37, 32, 40, 95}, 167, 3},
		{"all positive contiguous", []int{1, 1, 1, 1, 1}, 3, 3},
		{"large single element", []int{10}, 5, 1},
		{"sum of all needed", []int{1, 1, 1}, 3, 3},
		{"negative prefix", []int{-1, 4, 2}, 5, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ShortestSubarrayAtLeastK2(tt.nums, tt.k)
			if got != tt.want {
				t.Errorf("ShortestSubarrayAtLeastK(%v, %d) = %d, want %d", tt.nums, tt.k, got, tt.want)
			}
		})
	}
}

func TestLongestSubarrayWithinLimit(t *testing.T) {
	tests := []struct {
		name  string
		nums  []int
		limit int
		want  int
	}{
		{"from comments", []int{10, 1, 2, 4, 7, 2}, 5, 4},
		{"single element", []int{5}, 0, 1},
		{"all equal", []int{3, 3, 3, 3}, 0, 4},
		{"limit zero distinct", []int{1, 2, 3}, 0, 1},
		{"whole array fits", []int{1, 2, 3, 4}, 3, 4},
		{"descending", []int{5, 4, 3, 2, 1}, 2, 3},
		{"ascending", []int{1, 2, 3, 4, 5}, 2, 3},
		{"large gap resets", []int{1, 2, 100, 3, 4}, 2, 2},
		{"alternating", []int{1, 5, 1, 5, 1}, 4, 5},
		{"negative values", []int{-3, -1, -2, -4}, 2, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := LongestSubarrayWithinLimit2(tt.nums, tt.limit)
			if got != tt.want {
				t.Errorf("LongestSubarrayWithinLimit(%v, %d) = %d, want %d", tt.nums, tt.limit, got, tt.want)
			}
		})
	}
}
