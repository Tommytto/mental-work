package binsearch

import (
	"math"
	"testing"
)

func TestBase(t *testing.T) {
	tests := []struct {
		name   string
		arr    []int
		target int
		want   int
	}{
		{"found middle", []int{1, 2, 3, 4, 5}, 3, 2},
		{"found first", []int{1, 2, 3, 4, 5}, 1, 0},
		{"found last", []int{1, 2, 3, 4, 5}, 5, 4},
		{"not found", []int{1, 2, 3, 4, 5}, 6, -1},
		{"empty array", []int{}, 1, -1},
		{"single element found", []int{5}, 5, 0},
		{"single element not found", []int{5}, 3, -1},
		{"two elements found first", []int{1, 2}, 1, 0},
		{"two elements found last", []int{1, 2}, 2, 1},
		{"two elements not found", []int{1, 3}, 2, -1},
		{"negative numbers", []int{-5, -3, -1, 0, 2}, -3, 1},
		{"target smaller than all", []int{2, 4, 6}, 1, -1},
		{"target between elements", []int{1, 3, 5, 7}, 4, -1},
		{"duplicates finds one", []int{1, 2, 2, 2, 3}, 2, 2},
		{"large array", []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23, 25, 27, 29}, 21, 10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Base(tt.arr, tt.target)
			if got != tt.want {
				t.Errorf("Base(%v, %d) = %d, want %d", tt.arr, tt.target, got, tt.want)
			}
		})
	}
}

func TestUpperBound(t *testing.T) {
	tests := []struct {
		name   string
		arr    []int
		target int
		want   int
	}{
		{"found first", []int{1, 2, 2, 4, 5}, 2, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := UpperBound(tt.arr, tt.target)
			if got != tt.want {
				t.Errorf("UpperBound(%v, %d) = %d, want %d", tt.arr, tt.target, got, tt.want)
			}
		})
	}
}

func TestMonotonicPredicateFirstTrue(t *testing.T) {
	tests := []struct {
		name string
		arr  []bool
		want int
	}{
		{"first true at middle", []bool{false, false, true, true, true}, 2},
		{"first true early", []bool{false, true, true, true}, 1},
		{"all true", []bool{true, true, true, true, true}, 0},
		{"empty", []bool{}, -1},
		{"only one false", []bool{false}, -1},
		{"only one true", []bool{true}, 0},
		{"only falses", []bool{false, false, false}, -1},
		{"only trues", []bool{true, true, true}, 0},
		{"first true at last", []bool{false, false, false, true}, 3},
		{"two elements false true", []bool{false, true}, 1},
		{"two elements false false", []bool{false, false}, -1},
		{"two elements true true", []bool{true, true}, 0},
		{"first true near end", []bool{false, false, false, true, true}, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MonotonicPredicateFirstTrue(tt.arr)
			if got != tt.want {
				t.Errorf("MonotonicPredicateFirstTrue(%v) = %d, want %d", tt.arr, got, tt.want)
			}
		})
	}
}

func TestMonotonicPredicateLastTrue(t *testing.T) {
	tests := []struct {
		name string
		arr  []bool
		want int
	}{
		{"last true at middle", []bool{true, true, true, false, false}, 2},
		{"last true at first", []bool{true, false, false, false}, 0},
		{"all true", []bool{true, true, true, true, true}, 4},
		{"empty", []bool{}, -1},
		{"only one false", []bool{false}, -1},
		{"only one true", []bool{true}, 0},
		{"only falses", []bool{false, false, false}, -1},
		{"only trues", []bool{true, true, true}, 2},
		{"two elements true false", []bool{true, false}, 0},
		{"two elements true true", []bool{true, true}, 1},
		{"two elements false false", []bool{false, false}, -1},
		{"last true near end", []bool{true, true, true, true, false}, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MonotonicPredicateLastTrue(tt.arr)
			if got != tt.want {
				t.Errorf("MonotonicPredicateLastTrue(%v) = %d, want %d", tt.arr, got, tt.want)
			}
		})
	}
}

func predicate(val int) bool {
	return (val * val) >= 30
}
func TestFirstTrue(t *testing.T) {

	// find smallest x, where x*x>=30
	result := FirstTrue(0, 30, func(i int) bool {
		return i*i >= 30
	})

	if result != 6 {
		t.Errorf("TestFirstTrue failed")
	}
}

func TestDoSqrt(t *testing.T) {
	precision := 1e-6

	t.Run("negative returns -1", func(t *testing.T) {
		got := DoSqrt(-1, precision)
		if got != -1 {
			t.Errorf("DoSqrt(-1) = %f, want -1", got)
		}
	})

	positiveTests := []struct {
		name string
		val  float64
	}{
		{"zero", 0},
		{"one", 1},
		{"perfect square 4", 4},
		{"perfect square 9", 9},
		{"perfect square 25", 25},
		{"perfect square 100", 100},
		{"non-perfect 2", 2},
		{"non-perfect 3", 3},
		{"non-perfect 7", 7},
		{"non-perfect 10", 10},
		{"fraction 0.25", 0.25},
		{"fraction 0.5", 0.5},
		{"fraction 0.01", 0.01},
		{"large 10000", 10000},
		{"large 123456", 123456},
	}

	for _, tt := range positiveTests {
		t.Run(tt.name, func(t *testing.T) {
			got := DoSqrt(tt.val, precision)
			want := math.Sqrt(tt.val)
			if math.Abs(got-want) > precision {
				t.Errorf("DoSqrt(%f) = %f, want %f", tt.val, got, want)
			}
		})
	}
}
