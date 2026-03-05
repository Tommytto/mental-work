#include "binsearch.cpp"
#include <gtest/gtest.h>
#include <vector>

TEST(Smoke, Works) { EXPECT_EQ(2 + 2, 4); }

TEST(BaseBinSearchTest, FindElement) {
  std::vector<int> arr = {1, 2, 3, 4, 5};
  int got = BaseBinSearch(arr, 3);
  EXPECT_EQ(got, 2);
}

TEST(BaseBinSearchTest, FindFirstElement) {
  std::vector<int> arr = {1, 2, 3, 4, 5};
  EXPECT_EQ(BaseBinSearch(arr, 1), 0);
}

TEST(BaseBinSearchTest, FindLastElement) {
  std::vector<int> arr = {1, 2, 3, 4, 5};
  EXPECT_EQ(BaseBinSearch(arr, 5), 4);
}

TEST(BaseBinSearchTest, ElementNotFound) {
  std::vector<int> arr = {1, 2, 3, 4, 5};
  EXPECT_EQ(BaseBinSearch(arr, 6), -1);
}

TEST(BaseBinSearchTest, EmptyArray) {
  std::vector<int> arr = {};
  EXPECT_EQ(BaseBinSearch(arr, 1), -1);
}

TEST(BaseBinSearchTest, SingleElementFound) {
  std::vector<int> arr = {42};
  EXPECT_EQ(BaseBinSearch(arr, 42), 0);
}

TEST(BaseBinSearchTest, SingleElementNotFound) {
  std::vector<int> arr = {42};
  EXPECT_EQ(BaseBinSearch(arr, 7), -1);
}

TEST(BaseBinSearchTest, TwoElements) {
  std::vector<int> arr = {10, 20};
  EXPECT_EQ(BaseBinSearch(arr, 10), 0);
  EXPECT_EQ(BaseBinSearch(arr, 20), 1);
  EXPECT_EQ(BaseBinSearch(arr, 15), -1);
}

TEST(BaseBinSearchTest, NegativeNumbers) {
  std::vector<int> arr = {-5, -3, -1, 0, 2, 4};
  EXPECT_EQ(BaseBinSearch(arr, -3), 1);
  EXPECT_EQ(BaseBinSearch(arr, 0), 3);
  EXPECT_EQ(BaseBinSearch(arr, -4), -1);
}

TEST(LowerBoundTest, ExactMatch) {
  std::vector<int> arr = {1, 2, 3, 4, 5};
  EXPECT_EQ(LowerBound(arr, 3), 2);
}

TEST(LowerBoundTest, FirstElement) {
  std::vector<int> arr = {1, 2, 3, 4, 5};
  EXPECT_EQ(LowerBound(arr, 1), 0);
}

TEST(LowerBoundTest, LastElement) {
  std::vector<int> arr = {1, 2, 3, 4, 5};
  EXPECT_EQ(LowerBound(arr, 5), 4);
}

TEST(LowerBoundTest, BetweenElements) {
  std::vector<int> arr = {1, 3, 5, 7, 9};
  EXPECT_EQ(LowerBound(arr, 4), 2);
  EXPECT_EQ(LowerBound(arr, 6), 3);
}

TEST(LowerBoundTest, SmallerThanAll) {
  std::vector<int> arr = {5, 10, 15};
  EXPECT_EQ(LowerBound(arr, 1), 0);
}

TEST(LowerBoundTest, LargerThanAll) {
  std::vector<int> arr = {5, 10, 15};
  EXPECT_EQ(LowerBound(arr, 20), -1);
}

TEST(LowerBoundTest, EmptyArray) {
  std::vector<int> arr = {};
  EXPECT_EQ(LowerBound(arr, 1), -1);
}

TEST(LowerBoundTest, Duplicates) {
  std::vector<int> arr = {1, 3, 3, 3, 5};
  EXPECT_EQ(LowerBound(arr, 3), 1);
}

TEST(LowerBoundTest, SingleElement) {
  std::vector<int> arr = {10};
  EXPECT_EQ(LowerBound(arr, 10), 0);
  EXPECT_EQ(LowerBound(arr, 5), 0);
  EXPECT_EQ(LowerBound(arr, 15), -1);
}

TEST(UpperBoundTest, ExactMatch) {
  std::vector<int> arr = {1, 2, 3, 4, 5};
  EXPECT_EQ(UpperBound(arr, 3), 3);
}

TEST(UpperBoundTest, FirstElement) {
  std::vector<int> arr = {1, 2, 3, 4, 5};
  EXPECT_EQ(UpperBound(arr, 1), 1);
}

TEST(UpperBoundTest, LastElement) {
  std::vector<int> arr = {1, 2, 3, 4, 5};
  EXPECT_EQ(UpperBound(arr, 5), -1);
}

TEST(UpperBoundTest, BetweenElements) {
  std::vector<int> arr = {1, 3, 5, 7, 9};
  EXPECT_EQ(UpperBound(arr, 4), 2);
  EXPECT_EQ(UpperBound(arr, 6), 3);
}

TEST(UpperBoundTest, SmallerThanAll) {
  std::vector<int> arr = {5, 10, 15};
  EXPECT_EQ(UpperBound(arr, 1), 0);
}

TEST(UpperBoundTest, LargerThanAll) {
  std::vector<int> arr = {5, 10, 15};
  EXPECT_EQ(UpperBound(arr, 20), -1);
}

TEST(UpperBoundTest, EmptyArray) {
  std::vector<int> arr = {};
  EXPECT_EQ(UpperBound(arr, 1), -1);
}

TEST(UpperBoundTest, Duplicates) {
  std::vector<int> arr = {1, 3, 3, 3, 5};
  EXPECT_EQ(UpperBound(arr, 3), 4);
}

TEST(UpperBoundTest, SingleElement) {
  std::vector<int> arr = {10};
  EXPECT_EQ(UpperBound(arr, 10), -1);
  EXPECT_EQ(UpperBound(arr, 5), 0);
  EXPECT_EQ(UpperBound(arr, 15), -1);
}

TEST(FirstTrueTest, BasicThreshold) {
  EXPECT_EQ(first_true(0, 10, [](int x) { return x >= 5; }), 5);
}

TEST(FirstTrueTest, AllTrue) {
  EXPECT_EQ(first_true(0, 10, [](int x) { return true; }), 0);
}

TEST(FirstTrueTest, AllFalse) {
  EXPECT_EQ(first_true(0, 10, [](int x) { return false; }), -1);
}

TEST(FirstTrueTest, FirstElementTrue) {
  EXPECT_EQ(first_true(0, 10, [](int x) { return x >= 0; }), 0);
}

TEST(FirstTrueTest, LastElementTrue) {
  EXPECT_EQ(first_true(0, 10, [](int x) { return x >= 10; }), 10);
}

TEST(FirstTrueTest, SingleElementTrue) {
  EXPECT_EQ(first_true(5, 5, [](int x) { return true; }), 5);
}

TEST(FirstTrueTest, SingleElementFalse) {
  EXPECT_EQ(first_true(5, 5, [](int x) { return false; }), -1);
}

TEST(FirstTrueTest, NegativeRange) {
  EXPECT_EQ(first_true(-10, 10, [](int x) { return x >= -3; }), -3);
}

TEST(FirstTrueTest, SquareRootApprox) {
  // find smallest x where x*x >= 100
  EXPECT_EQ(first_true(0, 100, [](int x) { return x * x >= 100; }), 10);
}

TEST(LastTrueTest, BasicThreshold) {
  EXPECT_EQ(last_true(0, 10, [](int x) { return x <= 5; }), 5);
}

TEST(LastTrueTest, AllTrue) {
  EXPECT_EQ(last_true(0, 10, [](int x) { return true; }), 10);
}

TEST(LastTrueTest, AllFalse) {
  EXPECT_EQ(last_true(0, 10, [](int x) { return false; }), -1);
}

TEST(LastTrueTest, FirstElementTrue) {
  EXPECT_EQ(last_true(0, 10, [](int x) { return x <= 0; }), 0);
}

TEST(LastTrueTest, LastElementTrue) {
  EXPECT_EQ(last_true(0, 10, [](int x) { return x <= 10; }), 10);
}

TEST(LastTrueTest, SingleElementTrue) {
  EXPECT_EQ(last_true(5, 5, [](int x) { return true; }), 5);
}

TEST(LastTrueTest, SingleElementFalse) {
  EXPECT_EQ(last_true(5, 5, [](int x) { return false; }), -1);
}

TEST(LastTrueTest, NegativeRange) {
  EXPECT_EQ(last_true(-10, 10, [](int x) { return x <= -3; }), -3);
}

TEST(LastTrueTest, SquareRootApprox) {
  // find largest x where x*x <= 100
  EXPECT_EQ(last_true(0, 100, [](int x) { return x * x <= 100; }), 10);
}

TEST(CustomSqrtTest, PerfectSquares) {
  EXPECT_NEAR(custom_sqrt(4, 0.001), 2.0, 0.001);
  EXPECT_NEAR(custom_sqrt(9, 0.001), 3.0, 0.001);
  EXPECT_NEAR(custom_sqrt(16, 0.001), 4.0, 0.001);
  EXPECT_NEAR(custom_sqrt(100, 0.001), 10.0, 0.001);
}

TEST(CustomSqrtTest, NonPerfectSquares) {
  EXPECT_NEAR(custom_sqrt(2, 0.001), 1.414, 0.001);
  EXPECT_NEAR(custom_sqrt(3, 0.001), 1.732, 0.001);
  EXPECT_NEAR(custom_sqrt(5, 0.001), 2.236, 0.001);
}

TEST(CustomSqrtTest, Zero) {
  EXPECT_NEAR(custom_sqrt(0, 0.001), 0.0, 0.001);
}

TEST(CustomSqrtTest, One) {
  EXPECT_NEAR(custom_sqrt(1, 0.001), 1.0, 0.001);
}

TEST(CustomSqrtTest, Negative) {
  EXPECT_EQ(custom_sqrt(-1, 0.001), -1);
}

TEST(CustomSqrtTest, FractionLessThanOne) {
  EXPECT_NEAR(custom_sqrt(0.25, 0.001), 0.5, 0.001);
  EXPECT_NEAR(custom_sqrt(0.01, 0.001), 0.1, 0.001);
}

TEST(CustomSqrtTest, HigherPrecision) {
  EXPECT_NEAR(custom_sqrt(2, 0.0001), 1.4142, 0.0001);
}

TEST(CustomSqrtTest, LargeValue) {
  EXPECT_NEAR(custom_sqrt(10000, 0.01), 100.0, 0.01);
}