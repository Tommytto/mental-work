#include "code.cpp"
#include <gtest/gtest.h>
#include <algorithm>
#include <vector>

// Helper: check that partition postconditions hold
void assertValidPartition(const std::vector<int> &nums, int left, int right, int pivotIdx) {
  ASSERT_GE(pivotIdx, left);
  ASSERT_LE(pivotIdx, right);
  int pivotVal = nums[pivotIdx];
  for (int i = left; i < pivotIdx; i++) {
    EXPECT_LT(nums[i], pivotVal) << "nums[" << i << "]=" << nums[i] << " should be < pivot " << pivotVal;
  }
  for (int i = pivotIdx + 1; i <= right; i++) {
    EXPECT_GE(nums[i], pivotVal) << "nums[" << i << "]=" << nums[i] << " should be >= pivot " << pivotVal;
  }
}

TEST(PartitionTest, BasicCase) {
  std::vector<int> nums = {3, 1, 4, 1, 5, 9, 2, 6};
  int pivotIdx = partition(nums, 0, nums.size() - 1);
  assertValidPartition(nums, 0, nums.size() - 1, pivotIdx);
}

TEST(PartitionTest, AlreadySorted) {
  std::vector<int> nums = {1, 2, 3, 4, 5};
  int pivotIdx = partition(nums, 0, 4);
  assertValidPartition(nums, 0, 4, pivotIdx);
}

TEST(PartitionTest, ReverseSorted) {
  std::vector<int> nums = {5, 4, 3, 2, 1};
  int pivotIdx = partition(nums, 0, 4);
  assertValidPartition(nums, 0, 4, pivotIdx);
}

TEST(PartitionTest, AllEqual) {
  std::vector<int> nums = {3, 3, 3, 3, 3};
  int pivotIdx = partition(nums, 0, 4);
  assertValidPartition(nums, 0, 4, pivotIdx);
}

TEST(PartitionTest, TwoElements) {
  std::vector<int> nums = {2, 1};
  int pivotIdx = partition(nums, 0, 1);
  assertValidPartition(nums, 0, 1, pivotIdx);
}

TEST(PartitionTest, TwoElementsSorted) {
  std::vector<int> nums = {1, 2};
  int pivotIdx = partition(nums, 0, 1);
  assertValidPartition(nums, 0, 1, pivotIdx);
}

TEST(PartitionTest, SingleElement) {
  std::vector<int> nums = {42};
  int pivotIdx = partition(nums, 0, 0);
  EXPECT_EQ(pivotIdx, 0);
  EXPECT_EQ(nums[0], 42);
}

TEST(PartitionTest, SubrangePartition) {
  std::vector<int> nums = {10, 5, 3, 8, 2, 7, 20};
  int pivotIdx = partition(nums, 1, 5);
  assertValidPartition(nums, 1, 5, pivotIdx);
  // Elements outside the range should be untouched
  EXPECT_EQ(nums[0], 10);
  EXPECT_EQ(nums[6], 20);
}

TEST(PartitionTest, PreservesElements) {
  std::vector<int> nums = {3, 1, 4, 1, 5, 9, 2, 6};
  std::vector<int> sorted_copy = nums;
  std::sort(sorted_copy.begin(), sorted_copy.end());

  partition(nums, 0, nums.size() - 1);

  std::sort(nums.begin(), nums.end());
  EXPECT_EQ(nums, sorted_copy);
}

TEST(PartitionTest, NegativeNumbers) {
  std::vector<int> nums = {-3, -1, -4, -1, -5};
  int pivotIdx = partition(nums, 0, 4);
  assertValidPartition(nums, 0, 4, pivotIdx);
}

TEST(PartitionTest, PivotIsMin) {
  std::vector<int> nums = {5, 3, 4, 2, 1};
  int pivotIdx = partition(nums, 0, 4);
  assertValidPartition(nums, 0, 4, pivotIdx);
  EXPECT_EQ(pivotIdx, 0);
}

TEST(PartitionTest, PivotIsMax) {
  std::vector<int> nums = {1, 3, 2, 4, 5};
  int pivotIdx = partition(nums, 0, 4);
  assertValidPartition(nums, 0, 4, pivotIdx);
  EXPECT_EQ(pivotIdx, 4);
}

// --- quickselect tests ---

TEST(QuickselectTest, FindMin) {
  std::vector<int> nums = {3, 1, 4, 1, 5, 9, 2, 6};
  EXPECT_EQ(quickselect(nums, 0), 1);
}

TEST(QuickselectTest, FindMax) {
  std::vector<int> nums = {3, 1, 4, 1, 5, 9, 2, 6};
  EXPECT_EQ(quickselect(nums, 7), 9);
}

TEST(QuickselectTest, FindMedian) {
  std::vector<int> nums = {7, 2, 5, 3, 1};
  // sorted: {1, 2, 3, 5, 7}, k=2 → 3
  EXPECT_EQ(quickselect(nums, 2), 3);
}

TEST(QuickselectTest, AllEqual) {
  std::vector<int> nums = {4, 4, 4, 4};
  EXPECT_EQ(quickselect(nums, 0), 4);
  // reset since quickselect mutates
  nums = {4, 4, 4, 4};
  EXPECT_EQ(quickselect(nums, 3), 4);
}

TEST(QuickselectTest, SingleElement) {
  std::vector<int> nums = {42};
  EXPECT_EQ(quickselect(nums, 0), 42);
}

TEST(QuickselectTest, TwoElements) {
  std::vector<int> nums = {5, 3};
  EXPECT_EQ(quickselect(nums, 0), 3);
  nums = {5, 3};
  EXPECT_EQ(quickselect(nums, 1), 5);
}

TEST(QuickselectTest, AlreadySorted) {
  std::vector<int> nums = {1, 2, 3, 4, 5};
  EXPECT_EQ(quickselect(nums, 0), 1);
  nums = {1, 2, 3, 4, 5};
  EXPECT_EQ(quickselect(nums, 4), 5);
  nums = {1, 2, 3, 4, 5};
  EXPECT_EQ(quickselect(nums, 2), 3);
}

TEST(QuickselectTest, ReverseSorted) {
  std::vector<int> nums = {5, 4, 3, 2, 1};
  EXPECT_EQ(quickselect(nums, 0), 1);
  nums = {5, 4, 3, 2, 1};
  EXPECT_EQ(quickselect(nums, 4), 5);
}

TEST(QuickselectTest, NegativeNumbers) {
  std::vector<int> nums = {-3, -1, -4, -1, -5};
  // sorted: {-5, -4, -3, -1, -1}
  EXPECT_EQ(quickselect(nums, 0), -5);
  nums = {-3, -1, -4, -1, -5};
  EXPECT_EQ(quickselect(nums, 2), -3);
  nums = {-3, -1, -4, -1, -5};
  EXPECT_EQ(quickselect(nums, 4), -1);
}

TEST(QuickselectTest, Duplicates) {
  std::vector<int> nums = {3, 1, 2, 1, 3};
  // sorted: {1, 1, 2, 3, 3}
  EXPECT_EQ(quickselect(nums, 1), 1);
  nums = {3, 1, 2, 1, 3};
  EXPECT_EQ(quickselect(nums, 3), 3);
}

TEST(QuickselectTest, EveryK) {
  std::vector<int> original = {10, 4, 7, 1, 8, 3};
  std::vector<int> sorted = original;
  std::sort(sorted.begin(), sorted.end());
  for (int k = 0; k < (int)original.size(); k++) {
    std::vector<int> nums = original;
    EXPECT_EQ(quickselect(nums, k), sorted[k]) << "Failed for k=" << k;
  }
}

// --- threeWayPartition tests ---

// Helper: verify three-way partition postcondition
void assertValidThreeWay(const std::vector<int> &nums, int k) {
  int n = nums.size();
  int i = 0;
  // Region 1: all < k
  while (i < n && nums[i] < k) i++;
  // Region 2: all == k
  while (i < n && nums[i] == k) i++;
  // Region 3: all > k
  while (i < n && nums[i] > k) i++;
  EXPECT_EQ(i, n) << "Array not properly three-way partitioned around k=" << k;
}

TEST(ThreeWayPartitionTest, BasicCase) {
  std::vector<int> nums = {3, 1, 2, 3, 5, 3, 4};
  threeWayPartition(nums, 3);
  assertValidThreeWay(nums, 3);
}

TEST(ThreeWayPartitionTest, AllEqual) {
  std::vector<int> nums = {3, 3, 3, 3};
  threeWayPartition(nums, 3);
  assertValidThreeWay(nums, 3);
}

TEST(ThreeWayPartitionTest, AllLessThanK) {
  std::vector<int> nums = {1, 2, 1, 2};
  threeWayPartition(nums, 5);
  assertValidThreeWay(nums, 5);
}

TEST(ThreeWayPartitionTest, AllGreaterThanK) {
  std::vector<int> nums = {7, 8, 9, 10};
  threeWayPartition(nums, 5);
  assertValidThreeWay(nums, 5);
}

TEST(ThreeWayPartitionTest, NoEquals) {
  std::vector<int> nums = {5, 1, 4, 2, 3};
  threeWayPartition(nums, 3);
  assertValidThreeWay(nums, 3);
}

TEST(ThreeWayPartitionTest, SingleElement) {
  std::vector<int> nums = {5};
  threeWayPartition(nums, 5);
  assertValidThreeWay(nums, 5);
}

TEST(ThreeWayPartitionTest, TwoElements) {
  std::vector<int> nums = {2, 1};
  threeWayPartition(nums, 1);
  assertValidThreeWay(nums, 1);
}

TEST(ThreeWayPartitionTest, PreservesElements) {
  std::vector<int> nums = {3, 1, 2, 3, 5, 3, 4};
  std::vector<int> sorted_copy = nums;
  std::sort(sorted_copy.begin(), sorted_copy.end());

  threeWayPartition(nums, 3);

  std::sort(nums.begin(), nums.end());
  EXPECT_EQ(nums, sorted_copy);
}

TEST(ThreeWayPartitionTest, NegativeNumbers) {
  std::vector<int> nums = {-1, -5, -3, -3, -2, -3, -4};
  threeWayPartition(nums, -3);
  assertValidThreeWay(nums, -3);
}

TEST(ThreeWayPartitionTest, DutchFlag) {
  // Classic Dutch National Flag: 0s, 1s, 2s
  std::vector<int> nums = {2, 0, 1, 2, 0, 1, 1, 0, 2};
  threeWayPartition(nums, 1);
  assertValidThreeWay(nums, 1);
}

TEST(ThreeWayPartitionTest, KNotPresent) {
  std::vector<int> nums = {1, 5, 2, 4, 3};
  threeWayPartition(nums, 6);
  // All elements < 6, so array should just be all in the "less" region
  assertValidThreeWay(nums, 6);
}

// --- hoarePartition tests ---

// Helper: check Hoare partition postcondition
// After Hoare partition with pivot = nums[left], the returned index p means:
// nums[left..p-1] <= pivot and nums[p..right] >= pivot
void assertValidHoarePartition(const std::vector<int> &nums, int left, int right, int pivot, int p) {
  ASSERT_GE(p, left);
  ASSERT_LE(p, right + 1);
  for (int i = left; i < p; i++) {
    EXPECT_LE(nums[i], pivot) << "nums[" << i << "]=" << nums[i] << " > pivot=" << pivot;
  }
  for (int i = p; i <= right; i++) {
    EXPECT_GE(nums[i], pivot) << "nums[" << i << "]=" << nums[i] << " < pivot=" << pivot;
  }
}

TEST(HoarePartitionTest, BasicCase) {
  std::vector<int> nums = {3, 1, 4, 1, 5, 9, 2, 6};
  int pivot = nums[0];
  int p = hoarePartition(nums, 0, 7);
  assertValidHoarePartition(nums, 0, 7, pivot, p);
}

TEST(HoarePartitionTest, AlreadySorted) {
  std::vector<int> nums = {1, 2, 3, 4, 5};
  int pivot = nums[0];
  int p = hoarePartition(nums, 0, 4);
  assertValidHoarePartition(nums, 0, 4, pivot, p);
}

TEST(HoarePartitionTest, ReverseSorted) {
  std::vector<int> nums = {5, 4, 3, 2, 1};
  int pivot = nums[0];
  int p = hoarePartition(nums, 0, 4);
  assertValidHoarePartition(nums, 0, 4, pivot, p);
}

TEST(HoarePartitionTest, AllEqual) {
  std::vector<int> nums = {3, 3, 3, 3};
  int pivot = nums[0];
  int p = hoarePartition(nums, 0, 3);
  assertValidHoarePartition(nums, 0, 3, pivot, p);
}

TEST(HoarePartitionTest, TwoElements) {
  std::vector<int> nums = {5, 3};
  int pivot = nums[0];
  int p = hoarePartition(nums, 0, 1);
  assertValidHoarePartition(nums, 0, 1, pivot, p);
}

TEST(HoarePartitionTest, TwoElementsSorted) {
  std::vector<int> nums = {3, 5};
  int pivot = nums[0];
  int p = hoarePartition(nums, 0, 1);
  assertValidHoarePartition(nums, 0, 1, pivot, p);
}

TEST(HoarePartitionTest, SingleElement) {
  std::vector<int> nums = {42};
  int pivot = nums[0];
  int p = hoarePartition(nums, 0, 0);
  assertValidHoarePartition(nums, 0, 0, pivot, p);
}

TEST(HoarePartitionTest, Subrange) {
  std::vector<int> nums = {99, 3, 1, 4, 1, 5, 99};
  int pivot = nums[1];
  int p = hoarePartition(nums, 1, 5);
  assertValidHoarePartition(nums, 1, 5, pivot, p);
  // Elements outside subrange should be untouched
  EXPECT_EQ(nums[0], 99);
  EXPECT_EQ(nums[6], 99);
}

TEST(HoarePartitionTest, NegativeNumbers) {
  std::vector<int> nums = {-2, -5, -1, -3, -4};
  int pivot = nums[0];
  int p = hoarePartition(nums, 0, 4);
  assertValidHoarePartition(nums, 0, 4, pivot, p);
}

TEST(HoarePartitionTest, PivotIsMin) {
  std::vector<int> nums = {1, 5, 3, 4, 2};
  int pivot = nums[0];
  int p = hoarePartition(nums, 0, 4);
  assertValidHoarePartition(nums, 0, 4, pivot, p);
}

TEST(HoarePartitionTest, PivotIsMax) {
  std::vector<int> nums = {5, 1, 3, 4, 2};
  int pivot = nums[0];
  int p = hoarePartition(nums, 0, 4);
  assertValidHoarePartition(nums, 0, 4, pivot, p);
}

TEST(HoarePartitionTest, PreservesElements) {
  std::vector<int> nums = {3, 1, 4, 1, 5, 9, 2, 6};
  std::vector<int> sorted_copy = nums;
  std::sort(sorted_copy.begin(), sorted_copy.end());

  hoarePartition(nums, 0, 7);

  std::sort(nums.begin(), nums.end());
  EXPECT_EQ(nums, sorted_copy);
}

TEST(HoarePartitionTest, Duplicates) {
  std::vector<int> nums = {3, 1, 3, 2, 3, 1, 3};
  int pivot = nums[0];
  int p = hoarePartition(nums, 0, 6);
  assertValidHoarePartition(nums, 0, 6, pivot, p);
}

// --- quicksort tests ---

TEST(QuicksortTest, BasicCase) {
  std::vector<int> nums = {3, 1, 4, 1, 5, 9, 2, 6};
  quicksort(nums, 0, nums.size() - 1);
  std::vector<int> expected = {1, 1, 2, 3, 4, 5, 6, 9};
  EXPECT_EQ(nums, expected);
}

TEST(QuicksortTest, AlreadySorted) {
  std::vector<int> nums = {1, 2, 3, 4, 5};
  quicksort(nums, 0, 4);
  std::vector<int> expected = {1, 2, 3, 4, 5};
  EXPECT_EQ(nums, expected);
}

TEST(QuicksortTest, ReverseSorted) {
  std::vector<int> nums = {5, 4, 3, 2, 1};
  quicksort(nums, 0, 4);
  std::vector<int> expected = {1, 2, 3, 4, 5};
  EXPECT_EQ(nums, expected);
}

TEST(QuicksortTest, AllEqual) {
  std::vector<int> nums = {3, 3, 3, 3};
  quicksort(nums, 0, 3);
  std::vector<int> expected = {3, 3, 3, 3};
  EXPECT_EQ(nums, expected);
}

TEST(QuicksortTest, SingleElement) {
  std::vector<int> nums = {42};
  quicksort(nums, 0, 0);
  EXPECT_EQ(nums[0], 42);
}

TEST(QuicksortTest, TwoElements) {
  std::vector<int> nums = {5, 3};
  quicksort(nums, 0, 1);
  std::vector<int> expected = {3, 5};
  EXPECT_EQ(nums, expected);
}

TEST(QuicksortTest, TwoElementsSorted) {
  std::vector<int> nums = {3, 5};
  quicksort(nums, 0, 1);
  std::vector<int> expected = {3, 5};
  EXPECT_EQ(nums, expected);
}

TEST(QuicksortTest, Duplicates) {
  std::vector<int> nums = {3, 1, 3, 2, 3, 1, 3};
  quicksort(nums, 0, 6);
  std::vector<int> expected = {1, 1, 2, 3, 3, 3, 3};
  EXPECT_EQ(nums, expected);
}

TEST(QuicksortTest, NegativeNumbers) {
  std::vector<int> nums = {-2, -5, -1, -3, -4};
  quicksort(nums, 0, 4);
  std::vector<int> expected = {-5, -4, -3, -2, -1};
  EXPECT_EQ(nums, expected);
}

TEST(QuicksortTest, MixedNegativePositive) {
  std::vector<int> nums = {3, -1, 0, -5, 2, 4};
  quicksort(nums, 0, 5);
  std::vector<int> expected = {-5, -1, 0, 2, 3, 4};
  EXPECT_EQ(nums, expected);
}

TEST(QuicksortTest, Subrange) {
  std::vector<int> nums = {99, 5, 3, 1, 4, 2, 88};
  quicksort(nums, 1, 5);
  // Only indices 1-5 should be sorted
  EXPECT_EQ(nums[0], 99);
  std::vector<int> expected_sub = {1, 2, 3, 4, 5};
  std::vector<int> actual_sub(nums.begin() + 1, nums.begin() + 6);
  EXPECT_EQ(actual_sub, expected_sub);
  EXPECT_EQ(nums[6], 88);
}

TEST(QuicksortTest, LargerArray) {
  std::vector<int> nums = {10, 7, 8, 9, 1, 5, 3, 4, 2, 6};
  quicksort(nums, 0, 9);
  std::vector<int> expected = {1, 2, 3, 4, 5, 6, 7, 8, 9, 10};
  EXPECT_EQ(nums, expected);
}
