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

func TestPreviousSmallerElement(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{"example", []int{4, 5, 2, 10, 8}, []int{-1, 4, -1, 2, 2}},
		{"empty", []int{}, []int{}},
		{"single", []int{5}, []int{-1}},
		{"ascending", []int{1, 2, 3, 4}, []int{-1, 1, 2, 3}},
		{"descending", []int{4, 3, 2, 1}, []int{-1, -1, -1, -1}},
		{"all equal", []int{3, 3, 3}, []int{-1, -1, -1}},
		{"valley", []int{5, 1, 6}, []int{-1, -1, 1}},
		{"two elements", []int{3, 1}, []int{-1, -1}},
		{"duplicates with smaller", []int{2, 2, 1}, []int{-1, -1, -1}},
		{"smaller at start", []int{1, 5, 3}, []int{-1, 1, 1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := PreviousSmallerElement(tt.nums)
			if !slices.Equal(got, tt.want) {
				t.Errorf("PreviousSmallerElement(%v) = %v, want %v", tt.nums, got, tt.want)
			}
		})
	}
}

func TestStockSpan(t *testing.T) {
	tests := []struct {
		name   string
		prices []int
		want   []int
	}{
		{"example", []int{100, 80, 60, 70, 60, 75, 85}, []int{1, 1, 1, 2, 1, 4, 6}},
		{"single", []int{50}, []int{1}},
		{"ascending", []int{10, 20, 30, 40}, []int{1, 2, 3, 4}},
		{"descending", []int{40, 30, 20, 10}, []int{1, 1, 1, 1}},
		{"all equal", []int{5, 5, 5, 5}, []int{1, 2, 3, 4}},
		{"two elements up", []int{10, 20}, []int{1, 2}},
		{"two elements down", []int{20, 10}, []int{1, 1}},
		{"valley then peak", []int{50, 30, 40, 60}, []int{1, 1, 2, 4}},
		{"peak then valley", []int{30, 60, 40, 20}, []int{1, 2, 1, 1}},
		{"dip and recover", []int{100, 60, 70, 65, 80, 105}, []int{1, 1, 2, 1, 4, 6}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := StockSpan(tt.prices)
			if !slices.Equal(got, tt.want) {
				t.Errorf("StockSpan(%v) = %v, want %v", tt.prices, got, tt.want)
			}
		})
	}
}
