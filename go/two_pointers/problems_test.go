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

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		wantK    int
		wantNums []int
	}{
		{"example 1", []int{1, 1, 2}, 2, []int{1, 2}},
		{"example 2", []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, 5, []int{0, 1, 2, 3, 4}},
		{"no duplicates", []int{1, 2, 3, 4}, 4, []int{1, 2, 3, 4}},
		{"all same", []int{5, 5, 5, 5}, 1, []int{5}},
		{"single element", []int{1}, 1, []int{1}},
		{"empty", []int{}, 0, []int{}},
		{"two unique", []int{1, 1, 2, 2, 2}, 2, []int{1, 2}},
		{"negatives", []int{-3, -1, -1, 0, 0, 2}, 4, []int{-3, -1, 0, 2}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotK := RemoveDuplicates(tt.nums)
			if gotK != tt.wantK {
				t.Errorf("RemoveDuplicates() k = %d, want %d", gotK, tt.wantK)
			}
			if !slices.Equal(tt.nums[:gotK], tt.wantNums) {
				t.Errorf("RemoveDuplicates() nums[:k] = %v, want %v", tt.nums[:gotK], tt.wantNums)
			}
		})
	}
}

func TestRemoveElementInPlace(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		val      int
		wantK    int
		wantNums []int // expected elements (sorted), order doesn't matter
	}{
		{"example 1", []int{3, 2, 2, 3}, 3, 2, []int{2, 2}},
		{"example 2", []int{0, 1, 2, 2, 3, 0, 4, 2}, 2, 5, []int{0, 0, 1, 3, 4}},
		{"no matches", []int{1, 2, 3}, 4, 3, []int{1, 2, 3}},
		{"all matches", []int{5, 5, 5}, 5, 0, []int{}},
		{"single keep", []int{1}, 2, 1, []int{1}},
		{"single remove", []int{1}, 1, 0, []int{}},
		{"empty", []int{}, 1, 0, []int{}},
		{"val at edges", []int{3, 1, 2, 3}, 3, 2, []int{1, 2}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotK := RemoveElementInPlace(tt.nums, tt.val)
			if gotK != tt.wantK {
				t.Errorf("RemoveElementInPlace() k = %d, want %d", gotK, tt.wantK)
			}
			got := make([]int, gotK)
			copy(got, tt.nums[:gotK])
			slices.Sort(got)
			if !slices.Equal(got, tt.wantNums) {
				t.Errorf("RemoveElementInPlace() sorted nums[:k] = %v, want %v", got, tt.wantNums)
			}
		})
	}
}

func TestSortColors(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{"example 1", []int{2, 0, 2, 1, 1, 0}, []int{0, 0, 1, 1, 2, 2}},
		{"example 2", []int{2, 0, 1}, []int{0, 1, 2}},
		{"already sorted", []int{0, 0, 1, 1, 2, 2}, []int{0, 0, 1, 1, 2, 2}},
		{"reverse sorted", []int{2, 2, 1, 1, 0, 0}, []int{0, 0, 1, 1, 2, 2}},
		{"all zeros", []int{0, 0, 0}, []int{0, 0, 0}},
		{"all ones", []int{1, 1, 1}, []int{1, 1, 1}},
		{"all twos", []int{2, 2, 2}, []int{2, 2, 2}},
		{"single element", []int{1}, []int{1}},
		{"two elements", []int{2, 0}, []int{0, 2}},
		{"empty", []int{}, []int{}},
		{"only 0 and 2", []int{2, 0, 2, 0}, []int{0, 0, 2, 2}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SortColors(tt.nums)
			if !slices.Equal(tt.nums, tt.want) {
				t.Errorf("SortColors() = %v, want %v", tt.nums, tt.want)
			}
		})
	}
}

func TestValidPalindrome(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{"example true", "A man, a plan, a canal: Panama", true},
		{"example false", "race a car", false},
		{"empty string", "", true},
		{"only spaces", "   ", true},
		{"only punctuation", ",.!?", true},
		{"single char", "a", true},
		{"mixed case", "Aa", true},
		{"digits palindrome", "12321", true},
		{"digits not palindrome", "12345", false},
		{"alphanumeric mix", "0P", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ValidPalindrome(tt.s)
			if got != tt.want {
				t.Errorf("ValidPalindrome(%q) = %v, want %v", tt.s, got, tt.want)
			}
		})
	}
}

func TestMoveZeroes(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{"example", []int{0, 1, 0, 3, 12}, []int{1, 3, 12, 0, 0}},
		{"single zero", []int{0}, []int{0}},
		{"no zeroes", []int{1, 2, 3}, []int{1, 2, 3}},
		{"all zeroes", []int{0, 0, 0}, []int{0, 0, 0}},
		{"zeroes at end", []int{1, 2, 0, 0}, []int{1, 2, 0, 0}},
		{"zeroes at start", []int{0, 0, 1, 2}, []int{1, 2, 0, 0}},
		{"empty", []int{}, []int{}},
		{"single nonzero", []int{5}, []int{5}},
		{"alternating", []int{0, 1, 0, 2, 0, 3}, []int{1, 2, 3, 0, 0, 0}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			MoveZeroes(tt.nums)
			if !slices.Equal(tt.nums, tt.want) {
				t.Errorf("MoveZeroes() = %v, want %v", tt.nums, tt.want)
			}
		})
	}
}

func TestSquaresOfSortedArray(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{"example 1", []int{-4, -1, 0, 3, 10}, []int{0, 1, 9, 16, 100}},
		{"example 2", []int{-7, -3, 2, 3, 11}, []int{4, 9, 9, 49, 121}},
		{"all positive", []int{1, 2, 3, 4}, []int{1, 4, 9, 16}},
		{"all negative", []int{-4, -3, -2, -1}, []int{1, 4, 9, 16}},
		{"single element", []int{5}, []int{25}},
		{"single negative", []int{-3}, []int{9}},
		{"zeroes", []int{0, 0, 0}, []int{0, 0, 0}},
		{"symmetric", []int{-3, -2, 0, 2, 3}, []int{0, 4, 4, 9, 9}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SquaresOfSortedArray(tt.nums)
			if !slices.Equal(got, tt.want) {
				t.Errorf("SquaresOfSortedArray(%v) = %v, want %v", tt.nums, got, tt.want)
			}
		})
	}
}
