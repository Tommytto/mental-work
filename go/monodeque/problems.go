package monodeque

// nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
// k := 3
// ответ: [3 3 5 5 6 7]
func SlidingWindowMax(nums []int, k int) []int {
	n := len(nums)
	var out []int
	maxDeque := NewMaxDeque(nums)
	l := 0
	for r := 0; r < n; r++ {
		maxDeque.PopExpired(l)
		maxDeque.Push(r)
		if r-l+1 == k {
			out = append(out, nums[maxDeque.Front()])
			l++
		}
	}

	return out
}

// nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
// k := 3
// ответ: [ -1, -3, -3, -3, 3, 3 ]
func SlidingWindowMin(nums []int, k int) []int {
	n := len(nums)

	minDeque := NewMinDeque(nums)
	out := make([]int, 0, n-k+1)

	for r := 0; r < n; r++ {
		left := r - k + 1
		minDeque.PopExpired(left)
		minDeque.Push(r)

		if left >= 0 {
			out = append(out, nums[minDeque.Front()])
		}
	}

	return out
}

type MaxDeque struct {
	idx  []int
	nums []int
}

func NewMaxDeque(nums []int) *MaxDeque {
	return &MaxDeque{
		idx:  make([]int, 0, len(nums)),
		nums: nums,
	}
}

func (d *MaxDeque) Empty() bool {
	return len(d.idx) == 0
}

func (d *MaxDeque) PopExpired(left int) {
	for len(d.idx) > 0 && d.idx[0] < left {
		d.idx = d.idx[1:]
	}
}

func (d *MaxDeque) Push(i int) {
	for len(d.idx) > 0 && d.nums[d.idx[len(d.idx)-1]] <= d.nums[i] {
		d.idx = d.idx[:len(d.idx)-1]
	}
	d.idx = append(d.idx, i)
}

func (d *MaxDeque) Front() int {
	return d.idx[0]
}

type MinDeque struct {
	idx  []int
	nums []int
}

func NewMinDeque(nums []int) *MinDeque {
	return &MinDeque{
		idx:  make([]int, 0, len(nums)),
		nums: nums,
	}
}

func (d *MinDeque) Empty() bool {
	return len(d.idx) == 0
}

func (d *MinDeque) PopExpired(left int) {
	for len(d.idx) > 0 && d.idx[0] < left {
		d.idx = d.idx[1:]
	}
}

func (d *MinDeque) Push(i int) {
	for len(d.idx) > 0 && d.nums[d.idx[len(d.idx)-1]] >= d.nums[i] {
		d.idx = d.idx[:len(d.idx)-1]
	}
	d.idx = append(d.idx, i)
}

func (d *MinDeque) Front() int {
	return d.idx[0]
}

// ShortestSubarrayAtLeastK([]int{1}, 1)          // 1
// ShortestSubarrayAtLeastK([]int{1, 2}, 4)       // -1
// ShortestSubarrayAtLeastK([]int{2, -1, 2}, 3)   // 3
// ShortestSubarrayAtLeastK([]int{1, 2, 3}, 3)    // 1
// ShortestSubarrayAtLeastK([]int{84, -37, 32, 40, 95}, 167) // 3
func ShortestSubarrayAtLeastK(nums []int, k int) int {
	n := len(nums)

	prefix := make([]int64, n+1)
	for i := 0; i < n; i++ {
		prefix[i+1] = prefix[i] + int64(nums[i])
	}

	best := n + 1
	dq := make([]int, 0, n+1)

	for i := 0; i <= n; i++ {
		for len(dq) > 0 && prefix[i]-prefix[dq[0]] >= int64(k) {
			if i-dq[0] < best {
				best = i - dq[0]
			}

			dq = dq[1:]
		}

		for len(dq) > 0 && prefix[dq[len(dq)-1]] >= prefix[i] {
			dq = dq[:len(dq)-1]
		}

		dq = append(dq, i)
	}

	if best == n+1 {
		return -1
	}

	return best
}

// ShortestSubarrayAtLeastK([]int{1}, 1)          // 1
// ShortestSubarrayAtLeastK([]int{1, 2}, 4)       // -1
// ShortestSubarrayAtLeastK([]int{2, -1, 2}, 3)   // 3
// ShortestSubarrayAtLeastK([]int{1, 2, 3}, 3)    // 1
// ShortestSubarrayAtLeastK([]int{84, -37, 32, 40, 95}, 167) // 3
func ShortestSubarrayAtLeastK2(nums []int, k int) int {
	n := len(nums)
	best := n + 1

	prefixSum := make([]int, n+1)
	for i := 0; i < n; i++ {
		prefixSum[i+1] = prefixSum[i] + nums[i]
	}

	dq := make([]int, 0, n+1)

	for i := 0; i <= n; i++ {
		for len(dq) > 0 && prefixSum[i]-prefixSum[dq[0]] >= k {
			if i-dq[0] < best {
				best = i - dq[0]
			}

			dq = dq[1:]
		}

		for len(dq) > 0 && prefixSum[dq[len(dq)-1]] >= prefixSum[i] {
			dq = dq[:len(dq)-1]
		}

		dq = append(dq, i)
	}

	if best == n+1 {
		return -1
	}

	return best
}

// nums = [10, 1, 2, 4, 7, 2]
// limit = 5
// look for max length array where diff between max and min <= then limit
func LongestSubarrayWithinLimit(nums []int, limit int) int {
	n := len(nums)
	best := -1
	l := 0

	maxDQ := make([]int, 0, n)
	minDQ := make([]int, 0, n)
	for r := 0; r < n; r++ {
		for len(maxDQ) > 0 && nums[maxDQ[len(maxDQ)-1]] <= nums[r] {
			maxDQ = maxDQ[:len(maxDQ)-1]
		}
		maxDQ = append(maxDQ, r)

		for len(minDQ) > 0 && nums[minDQ[len(minDQ)-1]] >= nums[r] {
			minDQ = minDQ[:len(minDQ)-1]
		}
		minDQ = append(minDQ, r)

		for len(maxDQ) > 0 && len(minDQ) > 0 && nums[maxDQ[0]]-nums[minDQ[0]] > limit {
			if l == maxDQ[0] {
				maxDQ = maxDQ[1:]
			}
			if l == minDQ[0] {
				minDQ = minDQ[1:]
			}
			l++
		}

		if r-l+1 > best {
			best = r - l + 1
		}
	}

	return best
}

// nums = [10, 1, 2, 4, 7, 2]
// limit = 5
// look for max length array where diff between max and min <= then limit
func LongestSubarrayWithinLimit2(nums []int, limit int) int {
	n := len(nums)
	best := 0

	l := 0
	minDQ := make([]int, 0, n)
	maxDQ := make([]int, 0, n)

	for r := 0; r < n; r++ {
		for len(maxDQ) > 0 && nums[maxDQ[len(maxDQ)-1]] <= nums[r] {
			maxDQ = maxDQ[:len(maxDQ)-1]
		}
		maxDQ = append(maxDQ, r)

		for len(minDQ) > 0 && nums[minDQ[len(minDQ)-1]] >= nums[r] {
			minDQ = minDQ[:len(minDQ)-1]
		}
		minDQ = append(minDQ, r)

		for len(minDQ) > 0 && len(maxDQ) > 0 && nums[maxDQ[0]]-nums[minDQ[0]] > limit {
			if l == maxDQ[0] {
				maxDQ = maxDQ[1:]
			}
			if l == minDQ[0] {
				minDQ = minDQ[1:]
			}
			l++
		}

		if r-l+1 > best {
			best = r - l + 1
		}
	}

	return best
}
