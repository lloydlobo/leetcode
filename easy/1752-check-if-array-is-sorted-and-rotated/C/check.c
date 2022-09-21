/* [[file:../../README.org::1752-check-if-array-is-sorted-and-rotated][1752-check-if-array-is-sorted-and-rotated]]
 */
#include <stdbool.h>
#include <stddef.h>
#include <stdio.h>

#define BOOL bool
// #define SIZE(arr) (sizeof(arr) / sizeof(arr[0]))

// 1752. Check if Array Is Sorted and Rotated
//
// Given an array nums, return true if the array was originally sorted in
// non-decreasing order, then rotated some number of positions (including zero).
// Otherwise, return false. There may be duplicates in the original array. Note:
// An array A rotated by x positions results in an array B of the same length
// such that A[i] == B[(i+x) % A.length], where % is the modulo operation
//
// Constraints:
// 1 <= nums.length <= 100
// 1 <= nums[i] <= 100
//
// Runtime: 2 ms.
// Memory Usage: 5.8 MB.
bool check(int *nums, int numsSize) {
  int count = 0;
  for (int i = 0; i < numsSize; i++) {
    printf("\n%d:%2d |%2d", i, nums[i], count);
    if (nums[i] > nums[(i + 1) % numsSize]) {
      count++;
    }
    if (count > 1) {
      return false;
    }
  }
  return true;
};

void printout(int *nums, bool want) {
  for (int i = 0; i < 5; i++) {
    printf("%d ", nums[i]);
  }
  printf("| want: %s\n", want ? "true" : "false");
}

int main() {
  int nums[5] = {3, 4, 5, 1, 2};
  bool want = true;
  printout(nums, want);
  int n = sizeof(nums) / sizeof(nums[0]); // is size_t but coerced as int.
  bool out = check(nums, n);
  printf("\ngot:%2s|want:%2s\n", out ? "true" : "false",
         want ? "true" : "false");
}
/* 1752-check-if-array-is-sorted-and-rotated ends here */
