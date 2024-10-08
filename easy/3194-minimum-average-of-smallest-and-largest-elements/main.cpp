/*
**
** $ find . -name '*cpp' | entr -crs 'zig c++ -std=c++20 main.cpp; ./a.out'
**
**
*/

#include <array>
#include <cassert>
#include <iostream>
#include <ranges>
#include <vector>

#include "flux.hpp"  // [Open in browser](https://github.com/tcbrindle/flux?tab=readme-ov-file)

// copied from https://flux.godbolt.org/z/KKcEbYnTx
int run_flux_example()
{
    constexpr auto result
        = flux::ints()                           // 0,1,2,3,...
              .filter(flux::pred::even)          // 0,2,4,6,...
              .map([](int i) { return i * 2; })  // 0,4,8,12,...
              .take(3)                           // 0,4,8
              .sum();                            // 12

    static_assert(result == 12);
    return result;
}

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

void debug_printvec(std::vector<int> vec)
{
    std::cout << "[ ";
    for (auto x : vec)
        std::cout << x << " ";
    std::cout << "]\n";
}


void run_examples()
{
    {
        auto is_small = [](int i) { return i < 10; };

        std::vector<int> vec { 1, 2, 3, 4, 5 };
        debug_printvec(vec);
        bool all_small = flux::all(vec, is_small);  // check whether every element is small
        assert(all_small);
    }

    {
        auto is_even = [](int i) { return i % 2 == 0; };
        std::vector<int> vec { 1, 3, 4, 5, 7, 9 };
        debug_printvec(vec);
        assert(flux::any(vec, is_even));  // found an even element
    }
}


int main(int argc, char *argv[])
{
    run_examples();

    std::vector<int> nums = { 0, 1, 2, 3, 4, 5, 6, 7, 8 };  //> 4
    std::cout << minimum_average(nums) << "\n";

    std::cout << run_flux_example() << "\n";

    std::cout << std::endl;
    return 0;
}
