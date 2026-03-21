package prefixsumdiffarr

import (
	"slices"
	"testing"
)

func TestPrefix1D(t *testing.T) {
	tests := []struct {
		name string
		a    []int
		l, r int
		want int
	}{
		{"full range", []int{1, 2, 3, 4, 5}, 0, 5, 15},
		{"single element", []int{1, 2, 3, 4, 5}, 2, 3, 3},
		{"first two", []int{1, 2, 3, 4, 5}, 0, 2, 3},
		{"last two", []int{1, 2, 3, 4, 5}, 3, 5, 9},
		{"empty range", []int{1, 2, 3}, 1, 1, 0},
		{"single element array", []int{7}, 0, 1, 7},
		{"with negatives", []int{-1, 2, -3, 4}, 0, 4, 2},
		{"middle subarray", []int{10, 20, 30, 40, 50}, 1, 4, 90},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := NewPrefix1D(tt.a)
			got := p.Sum(tt.l, tt.r)
			if got != tt.want {
				t.Errorf("Prefix1D(%v).Sum(%d, %d) = %d, want %d", tt.a, tt.l, tt.r, got, tt.want)
			}
		})
	}
}

func TestPivotIndex(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"from comments", []int{1, 7, 3, 6, 5, 6}, 3},
		{"pivot at start", []int{0, 1, -1}, 0},
		{"pivot at end", []int{-1, 1, 0}, 2},
		{"no pivot", []int{1, 2, 3}, -1},
		{"single element", []int{5}, 0},
		{"two elements equal", []int{1, 1}, -1},
		{"all zeros", []int{0, 0, 0}, 0},
		{"negative values", []int{-1, -1, -1, -1, 0, 1}, 1},
		{"first valid pivot", []int{2, 1, -1, 1, -1, 1, -1}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := PivotIndex(tt.nums)
			if got != tt.want {
				t.Errorf("PivotIndex(%v) = %d, want %d", tt.nums, got, tt.want)
			}
		})
	}
}

func TestSubarraySum(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want int
	}{
		{"from comments", []int{1, 1, 1}, 2, 2},
		{"single match", []int{1, 2, 3}, 3, 2},
		{"no match", []int{1, 2, 3}, 7, 0},
		{"k is zero", []int{0, 0, 0}, 0, 6},
		{"negative values", []int{-1, -1, 1}, 0, 1},
		{"single element match", []int{5}, 5, 1},
		{"single element no match", []int{5}, 3, 0},
		{"all same", []int{2, 2, 2}, 4, 2},
		{"negative k", []int{-1, -1, -1}, -2, 2},
		{"mixed signs", []int{1, -1, 1, -1}, 0, 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SubarraySum(tt.nums, tt.k)
			if got != tt.want {
				t.Errorf("SubarraySum(%v, %d) = %d, want %d", tt.nums, tt.k, got, tt.want)
			}
		})
	}
}

func TestDiff1D(t *testing.T) {
	tests := []struct {
		name string
		n    int
		ops  [][3]int // {l, r, delta}
		want []int
	}{
		{"single add", 5, [][3]int{{1, 3, 2}}, []int{0, 2, 2, 2, 0}},
		{"full range", 4, [][3]int{{0, 3, 1}}, []int{1, 1, 1, 1}},
		{"overlapping adds", 5, [][3]int{{0, 2, 3}, {1, 4, 1}}, []int{3, 4, 4, 1, 1}},
		{"single element range", 3, [][3]int{{1, 1, 5}}, []int{0, 5, 0}},
		{"no ops", 3, [][3]int{}, []int{0, 0, 0}},
		{"negative delta", 4, [][3]int{{0, 3, 5}, {1, 2, -3}}, []int{5, 2, 2, 5}},
		{"size one", 1, [][3]int{{0, 0, 7}}, []int{7}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := NewDiff1D(tt.n)
			for _, op := range tt.ops {
				d.Add(op[0], op[1], op[2])
			}
			got := d.Build()
			if !slices.Equal(got, tt.want) {
				t.Errorf("Diff1D ops=%v => %v, want %v", tt.ops, got, tt.want)
			}
		})
	}
}

func TestRangeAddition(t *testing.T) {
	tests := []struct {
		name    string
		length  int
		updates [][]int
		want    []int
	}{
		{"from screenshot", 5, [][]int{{1, 3, 2}, {2, 4, 3}, {0, 2, -2}}, []int{-2, 0, 3, 5, 3}},
		{"single update", 3, [][]int{{0, 2, 5}}, []int{5, 5, 5}},
		{"no updates", 4, nil, []int{0, 0, 0, 0}},
		{"single element update", 5, [][]int{{2, 2, 10}}, []int{0, 0, 10, 0, 0}},
		{"full overlap", 3, [][]int{{0, 2, 1}, {0, 2, 2}}, []int{3, 3, 3}},
		{"adjacent ranges", 4, [][]int{{0, 1, 3}, {2, 3, 7}}, []int{3, 3, 7, 7}},
		{"size one", 1, [][]int{{0, 0, 4}}, []int{4}},
		{"negative delta", 3, [][]int{{0, 2, 5}, {1, 1, -3}}, []int{5, 2, 5}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RangeAddition(tt.length, tt.updates)
			if !slices.Equal(got, tt.want) {
				t.Errorf("RangeAddition(%d, %v) = %v, want %v", tt.length, tt.updates, got, tt.want)
			}
		})
	}
}

func TestCarPooling(t *testing.T) {
	tests := []struct {
		name     string
		trips    [][]int
		capacity int
		want     bool
	}{
		{"example 1 true", [][]int{{2, 1, 5}, {3, 5, 7}}, 3, true},
		{"example 2 false", [][]int{{2, 1, 5}, {3, 3, 7}}, 4, false},
		{"single trip fits", [][]int{{3, 0, 5}}, 3, true},
		{"single trip too many", [][]int{{4, 0, 5}}, 3, false},
		{"no overlap", [][]int{{2, 0, 3}, {2, 4, 7}}, 2, true},
		{"exact capacity", [][]int{{2, 0, 5}, {3, 0, 5}}, 5, true},
		{"over capacity", [][]int{{2, 0, 5}, {3, 0, 5}}, 4, false},
		{"sequential trips", [][]int{{1, 0, 2}, {1, 2, 4}, {1, 4, 6}}, 1, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CarPooling(tt.trips, tt.capacity)
			if got != tt.want {
				t.Errorf("CarPooling(%v, %d) = %v, want %v", tt.trips, tt.capacity, got, tt.want)
			}
		})
	}
}

func TestPrefix2D(t *testing.T) {
	t.Run("2x2 grid", func(t *testing.T) {
		grid := [][]int{
			{1, 2},
			{3, 4},
		}
		p := NewPrefix2D(grid)

		tests := []struct {
			name           string
			r1, c1, r2, c2 int
			want           int
		}{
			{"whole grid", 0, 0, 2, 2, 10},
			{"top-left cell", 0, 0, 1, 1, 1},
			{"bottom-right cell", 1, 1, 2, 2, 4},
			{"top row", 0, 0, 1, 2, 3},
			{"bottom row", 1, 0, 2, 2, 7},
			{"left col", 0, 0, 2, 1, 4},
			{"right col", 0, 1, 2, 2, 6},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got := p.Sum(tt.r1, tt.c1, tt.r2, tt.c2)
				if got != tt.want {
					t.Errorf("Prefix2D.Sum(%d,%d,%d,%d) = %d, want %d", tt.r1, tt.c1, tt.r2, tt.c2, got, tt.want)
				}
			})
		}
	})

	t.Run("3x3 grid", func(t *testing.T) {
		grid := [][]int{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
		}
		p := NewPrefix2D(grid)

		tests := []struct {
			name           string
			r1, c1, r2, c2 int
			want           int
		}{
			{"whole grid", 0, 0, 3, 3, 45},
			{"bottom-right 2x2", 1, 1, 3, 3, 28},
			{"top-left 2x2", 0, 0, 2, 2, 12},
			{"middle cell", 1, 1, 2, 2, 5},
			{"first row", 0, 0, 1, 3, 6},
			{"last row", 2, 0, 3, 3, 24},
			{"first col", 0, 0, 3, 1, 12},
			{"last col", 0, 2, 3, 3, 18},
			{"center row", 1, 0, 2, 3, 15},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got := p.Sum(tt.r1, tt.c1, tt.r2, tt.c2)
				if got != tt.want {
					t.Errorf("Prefix2D.Sum(%d,%d,%d,%d) = %d, want %d", tt.r1, tt.c1, tt.r2, tt.c2, got, tt.want)
				}
			})
		}
	})
}

func TestNumMatrix(t *testing.T) {
	matrix := [][]int{
		{3, 0, 1, 4, 2},
		{5, 6, 3, 2, 1},
		{1, 2, 0, 1, 5},
		{4, 1, 0, 1, 7},
		{1, 0, 3, 0, 5},
	}
	n := NewNumMatrix(matrix)

	tests := []struct {
		name                   string
		row1, col1, row2, col2 int
		want                   int
	}{
		{"from screenshot", 2, 1, 4, 3, 8},
		{"whole grid", 0, 0, 4, 4, 58},
		{"single cell", 0, 0, 0, 0, 3},
		{"first row", 0, 0, 0, 4, 10},
		{"first col", 0, 0, 4, 0, 14},
		{"top-left 2x2", 0, 0, 1, 1, 14},
		{"bottom-right 2x2", 3, 3, 4, 4, 13},
		{"middle 3x3", 1, 1, 3, 3, 16},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := n.SumRegion(tt.row1, tt.col1, tt.row2, tt.col2)
			if got != tt.want {
				t.Errorf("SumRegion(%d,%d,%d,%d) = %d, want %d", tt.row1, tt.col1, tt.row2, tt.col2, got, tt.want)
			}
		})
	}
}

func TestMatrixBlockSum(t *testing.T) {
	tests := []struct {
		name string
		mat  [][]int
		k    int
		want [][]int
	}{
		{
			name: "3x3 k=1",
			mat:  [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			k:    1,
			want: [][]int{{12, 21, 16}, {27, 45, 33}, {24, 39, 28}},
		},
		{
			name: "2x2 k=1",
			mat:  [][]int{{1, 2}, {3, 4}},
			k:    1,
			want: [][]int{{10, 10}, {10, 10}},
		},
		{
			name: "1x1",
			mat:  [][]int{{5}},
			k:    0,
			want: [][]int{{5}},
		},
		{
			name: "k larger than matrix",
			mat:  [][]int{{1, 2}, {3, 4}},
			k:    5,
			want: [][]int{{10, 10}, {10, 10}},
		},
		{
			name: "k=0 returns original",
			mat:  [][]int{{1, 2, 3}, {4, 5, 6}},
			k:    0,
			want: [][]int{{1, 2, 3}, {4, 5, 6}},
		},
		{
			name: "empty matrix",
			mat:  [][]int{},
			k:    1,
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MatrixBlockSum(tt.mat, tt.k)
			if len(got) != len(tt.want) {
				t.Fatalf("MatrixBlockSum() returned %d rows, want %d", len(got), len(tt.want))
			}
			for i := range got {
				if !slices.Equal(got[i], tt.want[i]) {
					t.Errorf("row %d = %v, want %v", i, got[i], tt.want[i])
				}
			}
		})
	}
}

func TestSubarraysDivByK(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want int
	}{
		{"example", []int{4, 5, 0, -2, -3, 1}, 5, 7},
		{"all zeros", []int{0, 0, 0}, 3, 6},
		{"single divisible", []int{6}, 3, 1},
		{"single not divisible", []int{1}, 3, 0},
		{"negatives", []int{-1, -9, -4, 0}, 5, 2},
		{"negative sum divisible", []int{-1, 2, 9}, 2, 2},
		{"k=1 all divisible", []int{3, 7, -2}, 1, 6},
		{"two elements", []int{5, 10}, 5, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SubarraysDivByK(tt.nums, tt.k)
			if got != tt.want {
				t.Errorf("SubarraysDivByK(%v, %d) = %d, want %d", tt.nums, tt.k, got, tt.want)
			}
		})
	}
}

func TestNumSubmatrixSumTarget(t *testing.T) {
	tests := []struct {
		name   string
		matrix [][]int
		target int
		want   int
	}{
		{
			name:   "example 1",
			matrix: [][]int{{0, 1, 0}, {1, 1, 1}, {0, 1, 0}},
			target: 0,
			want:   4,
		},
		{
			name:   "single element match",
			matrix: [][]int{{5}},
			target: 5,
			want:   1,
		},
		{
			name:   "single element no match",
			matrix: [][]int{{5}},
			target: 3,
			want:   0,
		},
		{
			name:   "all ones target 2",
			matrix: [][]int{{1, 1}, {1, 1}},
			target: 2,
			want:   4,
		},
		{
			name:   "negatives",
			matrix: [][]int{{1, -1}, {-1, 1}},
			target: 0,
			want:   5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NumSubmatrixSumTarget(tt.matrix, tt.target)
			if got != tt.want {
				t.Errorf("NumSubmatrixSumTarget(%v, %d) = %d, want %d", tt.matrix, tt.target, got, tt.want)
			}
		})
	}
}
