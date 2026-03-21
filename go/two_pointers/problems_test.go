package twopointers

import (
	"slices"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   []int
	}{
		{"basic", []int{2, 7, 11, 15}, 9, []int{1, 2}},
		{"middle pair", []int{2, 3, 4}, 6, []int{1, 3}},
		{"adjacent", []int{-1, 0}, -1, []int{1, 2}},
		{"negatives", []int{-3, -1, 0, 2, 4}, -4, []int{1, 2}},
		{"large target", []int{1, 2, 3, 4, 5}, 9, []int{4, 5}},
		{"duplicates", []int{1, 1, 3, 5}, 2, []int{1, 2}},
		{"two elements", []int{5, 10}, 15, []int{1, 2}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := TwoSum(tt.nums, tt.target)
			if !slices.Equal(got, tt.want) {
				t.Errorf("TwoSum(%v, %d) = %v, want %v", tt.nums, tt.target, got, tt.want)
			}
		})
	}
}

func TestContainerWithMostWater(t *testing.T) {
	tests := []struct {
		name   string
		height []int
		want   int
	}{
		{"example", []int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
		{"equal heights", []int{5, 5, 5, 5}, 15},
		{"two elements", []int{1, 1}, 1},
		{"ascending", []int{1, 2, 3, 4, 5}, 6},
		{"descending", []int{5, 4, 3, 2, 1}, 6},
		{"one tall", []int{1, 1, 1, 100, 1, 1, 1}, 6},
		{"empty", []int{}, 0},
		{"single", []int{5}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ContainerWithMostWater(tt.height)
			if got != tt.want {
				t.Errorf("ContainerWithMostWater(%v) = %d, want %d", tt.height, got, tt.want)
			}
		})
	}
}
