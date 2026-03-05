#include <functional>
#include <vector>

int BaseBinSearch(const std::vector<int> &arr, int target) {
  int n = arr.size();
  if (n == 0)
    return -1;

  int left = 0;
  int right = n - 1;

  while (left <= right) {
    int mid = left + (right - left) / 2;

    if (arr[mid] == target)
      return mid;
    if (arr[mid] < target) {
      left = mid + 1;
    } else {
      right = mid - 1;
    }
  }

  return -1;
}

int LowerBound(std::vector<int> &arr, int target) {
  int n = arr.size();
  if (n == 0) {
    return -1;
  }

  int result = -1;
  int left = 0;
  int right = n - 1;

  while (left <= right) {
    int mid = left + (right - left) / 2;

    if (arr[mid] >= target) {
      result = mid;
      right = mid - 1;
    } else {
      left = mid + 1;
    }
  }

  return result;
}

int UpperBound(std::vector<int> &arr, int target) {
  int n = arr.size();
  if (n == 0)
    return -1;

  int result = -1;
  int left = 0;
  int right = n - 1;

  while (left <= right) {
    int mid = left + (right - left) / 2;

    if (arr[mid] > target) {
      result = mid;
      right = mid - 1;
    } else {
      left = mid + 1;
    }
  }

  return result;
}

int first_true(int lo, int hi, std::function<bool(int)> pred) {
  int result = -1;

  int left = lo;
  int right = hi;

  while (left <= right) {
    int mid = left + (right - left) / 2;
    bool ok = pred(mid);

    if (ok) {
      result = mid;
      right = mid - 1;
    } else {
      left = mid + 1;
    }
  }

  return result;
}

int last_true(int lo, int hi, std::function<bool(int)> pred) {
  int result = -1;

  int left = lo;
  int right = hi;

  while (left <= right) {
    int mid = left + (right - left) / 2;
    bool ok = pred(mid);

    if (ok) {
      result = mid;
      left = mid + 1;
    } else {
      right = mid - 1;
    }
  }

  return result;
}

float custom_sqrt(float val, float precision) {
  if (val < 0)
    return -1;

  float left;
  float right;
  if (val < 1) {
    left = 0;
    right = 1;
  } else {
    left = 1;
    right = val;
  }

  while ((right - left) > precision) {
    float mid = left + (right - left) / 2;

    if ((mid * mid) < val) {
      left = mid;
    } else {
      right = mid;
    }
  }

  return left;
}