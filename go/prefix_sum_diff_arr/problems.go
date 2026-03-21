package prefixsumdiffarr

type Prefix1D struct {
	pre []int
}

func NewPrefix1D(a []int) Prefix1D {
	pre := make([]int, len(a)+1)

	for i, num := range a {
		pre[i+1] = pre[i] + num
	}

	return Prefix1D{
		pre: pre,
	}
}

func (p Prefix1D) Sum(l, r int) int {
	return p.pre[r] - p.pre[l]
}

// nums := []int{1, 7, 3, 6, 5, 6}
// 3
func PivotIndex(nums []int) int {
	leftSum := 0
	total := 0
	for _, num := range nums {
		total += num
	}
	for i, num := range nums {
		rightSum := total - leftSum - num

		if rightSum == leftSum {
			return i
		}

		leftSum += num
	}

	return -1
}

// nums := []int{1, 1, 1}
// k := 2
// answer == 2
func SubarraySum(nums []int, k int) int {
	count := 0
	prefix := 0
	seen := map[int]int{
		0: 1,
	}

	for i := 0; i < len(nums); i++ {
		prefix += nums[i]
		count += seen[prefix-k]
		seen[prefix] += 1
	}

	return count
}

type Diff1D struct {
	diff []int
	n    int
}

func NewDiff1D(n int) *Diff1D {
	return &Diff1D{
		n:    n,
		diff: make([]int, n),
	}
}

// inclusive
func (d *Diff1D) Add(l, r, delta int) {
	if l < d.n {
		d.diff[l] += delta
	}
	if r+1 < d.n {
		d.diff[r+1] -= delta
	}
}

func (d *Diff1D) Build() []int {
	if d.n == 0 {
		return nil
	}
	result := make([]int, d.n)
	running := 0

	for i := 0; i < d.n; i++ {
		running += d.diff[i]
		result[i] = running
	}

	return result
}

func RangeAddition(length int, updates [][]int) []int {
	diff1D := NewDiff1D(length)

	for _, upd := range updates {
		l, r, delta := upd[0], upd[1], upd[2]

		diff1D.Add(l, r, delta)
	}

	return diff1D.Build()
}

func CarPooling(trips [][]int, capacity int) bool {
	maxTo := 0
	for _, trip := range trips {
		if trip[2] > maxTo {
			maxTo = trip[2]
		}
	}

	diff1D := NewDiff1D(maxTo)
	for _, trip := range trips {
		numPassengers, start, stop := trip[0], trip[1], trip[2]
		diff1D.Add(start, stop-1, numPassengers)
	}

	for _, val := range diff1D.Build() {
		if val > capacity {
			return false
		}
	}

	return true
}

//	grid := [][]int{
//		{1, 2},
//		{3, 4},
//	}
type Prefix2D struct {
	pre [][]int
}

func NewPrefix2D(grid [][]int) Prefix2D {
	rows := len(grid) + 1
	cols := len(grid[0]) + 1

	pre := make([][]int, rows)
	for i := range rows {
		pre[i] = make([]int, cols)
	}

	for i := range grid {
		for j := range grid[i] {
			pre[i+1][j+1] =
				grid[i][j] +
					pre[i+1][j] +
					pre[i][j+1] -
					pre[i][j]
		}
	}

	return Prefix2D{
		pre: pre,
	}
}
func (p Prefix2D) Sum(r1, c1, r2, c2 int) int {
	return p.pre[r2][c2] + p.pre[r1][c1] - p.pre[r1][c2] - p.pre[r2][c1]
}

type NumMatrix struct {
	prefix Prefix2D
}

func NewNumMatrix(matrix [][]int) NumMatrix {
	return NumMatrix{
		prefix: NewPrefix2D(matrix),
	}
}

func (n *NumMatrix) SumRegion(row1, col1, row2, col2 int) int {
	return n.prefix.Sum(row1, col1, row2+1, col2+1)
}

func MatrixBlockSum(mat [][]int, k int) [][]int {
	if len(mat) == 0 || len(mat[0]) == 0 {
		return nil
	}

	prefix2D := NewPrefix2D(mat)
	result := make([][]int, len(mat))
	for i := range mat {
		result[i] = make([]int, len(mat[0]))
	}

	for i := range mat {
		for j := range mat[0] {
			r1 := max(i-k, 0)
			c1 := max(j-k, 0)
			r2 := min(i+k+1, len(mat))
			c2 := min(j+k+1, len(mat[0]))

			result[i][j] = prefix2D.Sum(r1, c1, r2, c2)
		}
	}

	return result
}

func SubarraysDivByK(nums []int, k int) int {
	count := 0
	freq := map[int]int{
		0: 1,
	}
	prefix := 0

	for i := 0; i < len(nums); i++ {
		num := nums[i]
		prefix += num

		rem := prefix % k
		if rem < 0 {
			rem += k
		}

		count += freq[rem]
		freq[rem] += 1
	}

	return count
}

func NumSubmatrixSumTarget(matrix [][]int, target int) int {
	prefix2D := NewPrefix2D(matrix)
	seen := map[int]int{
		0: 1,
	}
	count := 0

	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			sum := prefix2D.Sum(0, 0, i, j)
			count += seen[target-sum]
			seen[sum] += 1
		}
	}

	return count
}
