/* [[file:../../README.org::1752-check-if-array-is-sorted-and-rotated][1752-check-if-array-is-sorted-and-rotated]]
 */
// 1752. Check if Array Is Sorted and Rotated
#include <math.h>
#include <stdbool.h>
#include <stddef.h>
#include <stdio.h>

#define BOOL bool

void print_example(int *nums, bool want);
bool is_f_if_unsorted_more_than_once(int count);
bool check(int arr[], int len);

// Array is internally a pointer so dont spend too much time thinking about the
// impossible. For your reference, it can look like:—
//
// C does not advocate to return the address of a local variable to outside of
// the function, so you would have to define the local variable as static
// variable.
int *nums(int *num, int size) {
  static int x[5];
  int i;
  int j;
  int nums_1[5] = {3, 4, 5, 1, 2};
  // int nums_2[4] = {2, 1, 3, 4};
  // int nums_3[3] = {1, 2, 3};

  for (int i = 0; i < size; i++) {
    printf("%i", num[i]);
  }
  printf("\n\n"); // print new line.
  // Set the seed.
  // srand((unsigned)time(NULL));

  int next_num_counter = 0;

  for (i = 0; i < 5; i++) {
    for (j = 0; j < 5; j++) {

      x[i] = nums_1[i];
    }
    next_num_counter++;
    printf("%i", x[i]);
  }
  printf("\n\n"); // print new line.

  return x;
}

int main() {
  int nums_1[5] = {3, 4, 5, 1, 2};
  int len = 5;
  static int x[5];
  for (int i = 0; i < 5; i++) {
    x[i] = *nums(nums_1, len);
    printf("%i\n", x[i]);
  }
  bool want = true;            // expected results.
  print_example(nums_1, want); // print results for convenience.

  int n = sizeof(nums_1) / sizeof(nums_1[0]); // is size_t but coerced as int.
  bool out = check(nums_1, n);
  printf("\ngot:%2s|want:%2s\n", out ? "true" : "false",
         want ? "true" : "false");

  printf("\n"); // add new line.
}

void print_example(int *nums, bool want) {
  for (int i = 0; i < 5; i++) {
    printf("%d ", nums[i]);
  }
  printf("| want: %s\n", want ? "true" : "false");
}

bool is_f_if_unsorted_more_than_once(int count) {
  bool output = true;
  if (count > 1) {
    output = false;
  }

  return output;
}

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
bool check(int arr[], int len) {
  int count = 0;

  for (int i = 0; i < len; i++) {
    bool curr_more_than_next = (arr[i] > arr[(i + 1) % len]);
    if (curr_more_than_next) {
      count++;
    }

    return is_f_if_unsorted_more_than_once(count);
  }

  return true;
}

/* 1752-check-if-array-is-sorted-and-rotated ends here */
