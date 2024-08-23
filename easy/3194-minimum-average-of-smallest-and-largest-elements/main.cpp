/*
**
** $ find . -name '*cpp' | entr -crs 'zig c++ -std=c++20 main.cpp; ./a.out'
**
**
**
*/


#include <iostream>
#include <ranges>
#include <vector>


// Finds the minimum sum from pairs formed by the smallest and largest values
// in the sorted vector. This result is then averaged by multiplying by `0.5`,
// which effectively gives a kind of balanced average of these sums.
auto minimum_average(std::vector<int> &nums) -> double
{
    using namespace std::ranges;  // provides `sort`, `min`

    int n = nums.size() / 2;       // n = len(nums) // 2
    sort(nums);                    // s: List[int] = sorted(nums)
    return std::transform_reduce(  // return 0.5 * min(map(operator.add, l, r))
               nums.begin(),       // l, r = s[:n], s[-n:][::-1]
               nums.begin() + n,
               nums.rbegin(), std::numeric_limits<int>::max(),
               min,
               std::plus {})
        * 0.5;
}

// LLM dump:
//
// **How It Works:**
//
// - **Range for Transformation:** `nums.begin()` to `nums.begin() + n` covers
//   the first half of the sorted vector.
// - **Reversed Range:** `nums.rbegin()` starts from the end of the vector and
//   iterates backwards to cover the second half.
// - **Pairing Elements:** The transformation function `std::plus{}` adds
//   corresponding elements from the first and reversed halves.
// - **Reduction Operation:** `std::ranges::min` is used to find the minimum
//   value among these sums.
// - **Initial Value:** `std::numeric_limits<int>::max()` initializes the
//   reduction to a very large number to ensure that the actual minimum value is
//   correctly identified.
//
// Given the vector `nums = {0, 1, 2, 3, 4, 5, 6, 7, 8}`:
//
// 1. **Sorted Vector:** `nums = {0, 1, 2, 3, 4, 5, 6, 7, 8}` (already sorted).
// 2. **Determine `n`:** `n = nums.size() / 2 = 9 / 2 = 4`.
// 3. **Pairing Elements:**
//    - First half: `{0, 1, 2, 3}`.
//    - Reversed second half: `{8, 7, 6, 5}`.
//    - Pair sums: `(0+8)`, `(1+7)`, `(2+6)`, `(3+5)` which are `8, 8, 8, 8`.
// 4. **Minimum Sum:** All sums are `8`, so the minimum is `8`.
// 5. **Final Result:** `0.5 * 8 = 4.0`.
//
int main(int argc, char *argv[])
{
    std::vector<int> nums = { 0, 1, 2, 3, 4, 5, 6, 7, 8 };  //> 4
    std::cout << minimum_average(nums) << "\n";
    return 0;
}
