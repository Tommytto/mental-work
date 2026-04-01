#include <vector>

void swap(std::vector<int> &nums, int i, int j) {
  int tmp = nums[i];
  nums[i] = nums[j];
  nums[j] = tmp;
}

// Lomuto partition
int partition(std::vector<int> &nums, int left, int right) {
  int pivot = nums[right];
  int write = left;

  for (int i = left; i < right; i++) {
    if (nums[i] < pivot) {
      swap(nums, i, write);
      write++;
    }
  }
  swap(nums, write, right);
  return write;
}

// Нужно реализовать функцию,
// которая возвращает k-й по величине smallest element по индексу
// в отсортированном порядке, то есть: k = n-1 → максимум
int quickselect(std::vector<int> &nums, int k) {
  int left = 0;
  int right = nums.size() - 1;

  while (left <= right) {
    int p = partition(nums, left, right);

    if (p == k) {
      return nums[p];
    } else if (p < k) {
      left = p + 1;
    } else {
      right = p - 1;
    }
  }
  return -1;
}

int threeWayPartition(std::vector<int> &nums, int k) {
  int n = nums.size();
  int low = 0;
  int mid = 0;
  int high = n - 1;

  while (mid <= high) {
    if (nums[mid] < k) {
      swap(nums, low, mid);
      low++;
      mid++;
    } else if (nums[mid] == k) {
      mid++;
    } else {
      swap(nums, mid, high);
      high--;
    }
  }

  return mid;
}

void sortColors(std::vector<int> &nums) {
  int low = 0;
  int mid = 0;
  int high = nums.size() - 1;

  while (mid <= high) {
    if (nums[mid] == 0) {
      swap(nums, mid, low);
      low++;
      mid++;
    } else if (nums[mid] == 1) {
      mid++;
    } else {
      swap(nums, mid, high);
      high--;
    }
  }
}

void moveZeroes(std::vector<int> &nums) {
  int left = 0;
  int n = nums.size();

  for (int i = 0; i < n; i++) {
    if (nums[i] != 0) {
      nums[left] = nums[i];
      left++;
    }
  }

  for (int i = left; i < n; i++) {
    nums[i] = 0;
  }
}

int removeElement(std::vector<int> &nums, int val) {
  int n = nums.size();

  int left = 0;

  for (int i = 0; i < n; i++) {
    if (nums[i] != val) {
      nums[left] = nums[i];
      left++;
    }
  }

  return left;
}

int removeDuplicates(std::vector<int> &nums) {
  int n = nums.size();
  int left = 1;

  for (int i = 1; i < n; i++) {
    if (nums[i] != nums[i - 1]) {
      nums[left] = nums[i];
      left++;
    }
  }

  return left;
}

int removeDuplicates2(std::vector<int> &nums) {
  int n = nums.size();
  int left = 2;

  for (int i = 2; i < n; i++) {
    if (nums[i] != nums[left - 2]) {
      nums[left] = nums[i];
      left++;
    }
  }

  return left;
}

int keepAtMostK(std::vector<int> &nums, int k) {
  int n = nums.size();
  int left = k;

  for (int i = k; i < n; i++) {
    if (nums[i] != nums[left - k]) {
      nums[left] = nums[i];
      left++;
    }
  }

  return left;
}

int hoarePartition(std::vector<int> &nums, int left, int right) {
  int i = left;
  int j = right;
  int pivot = nums[left];

  while (true) {
    while (nums[i] < pivot) {
      i++;
    }

    while (nums[j] > pivot) {
      j--;
    }

    if (i >= j) {
      return j;
    }

    swap(nums, i, j);
    i++;
    j--;
  }

  return j;
}

void quicksort(std::vector<int> &nums, int left, int right) {
  if (left >= right) {
    return;
  }

  int p = hoarePartition(nums, left, right);
  quicksort(nums, left, p);
  quicksort(nums, p + 1, right);
}

int findKthLargest(std::vector<int> &nums, int k) {
  return quickselect(nums, nums.size() - k);
}