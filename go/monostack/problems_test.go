package monostack

import (
	"slices"
	"testing"
)

func TestNextGreaterElement(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{"example", []int{2, 1, 2, 4, 3}, []int{4, 2, 4, -1, -1}},
		{"empty", []int{}, []int{}},
		{"single", []int{5}, []int{-1}},
		{"ascending", []int{1, 2, 3, 4}, []int{2, 3, 4, -1}},
		{"descending", []int{4, 3, 2, 1}, []int{-1, -1, -1, -1}},
		{"all equal", []int{3, 3, 3}, []int{-1, -1, -1}},
		{"peak", []int{1, 5, 2}, []int{5, -1, -1}},
		{"valley", []int{5, 1, 6}, []int{6, 6, -1}},
		{"two elements", []int{1, 2}, []int{2, -1}},
		{"duplicates with greater", []int{1, 1, 2}, []int{2, 2, -1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NextGreaterElement(tt.nums)
			if !slices.Equal(got, tt.want) {
				t.Errorf("NextGreaterElement(%v) = %v, want %v", tt.nums, got, tt.want)
			}
		})
	}
}
