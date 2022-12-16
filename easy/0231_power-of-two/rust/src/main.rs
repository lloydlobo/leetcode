//! # Given an integer n, return true if it is a power of two. Otherwise, return false.
//! # An integer n is a power of two, if there exists an integer x such that n == 2x.
//! #
//! # Example 1:
//! #
//! # Input: n = 1
//! # Output: true
//! # Explanation: 2^0 = 1
//! #
//! # Example 2:
//! #
//! # Input: n = 16
//! # Output: true
//! # Explanation: 2^4 = 16
//! #
//! # Example 3:
//! #
//! # Input: n = 3
//! # Output: false
//! #
//! # Constraints:
//! #   -2^31 <= n <= 2^31 - 1

fn main() {
    let n = base_pow(2i32, 4u32);
    let result: bool = Solution::is_power_of_two(n);
    dbg!(result);

    let result: bool = Solution::is_power_of_two(3);
    dbg!(result);
}

#[derive(Debug)]
struct Solution {
    // pub cache: HashMap<i32, bool>,
}

impl Solution {
    /// Implements `is_power_of_two` where given an integer n, return true if
    /// it is a power of two. Otherwise, return false.
    ///
    /// An integer n is a power of two, if there exists an integer x
    /// such that n == 2x.
    ///
    /// # Examples
    ///
    /// ```
    /// let result = Solution::is_power_of_two(1);
    /// assert_eq!(result, true); // Explanation: 2^0 = 1.
    /// let result = Solution::is_power_of_two(16);
    /// assert_eq!(result, true); // Explanation: 2^4 = 16.
    /// let result = Solution::is_power_of_two(3);
    /// assert_eq!(result, false); // false
    /// ```
    ///
    /// # Leetcode distribution result:
    ///
    /// * Runtime 0 ms Beats 100%.
    /// * Memory 2 MB Beats 92.31%.
    pub fn is_power_of_two(n: i32) -> bool {
        match n {
            0i32 => false,
            1i32 => true,
            _ => {
                let mut n: i32 = n;
                while n % 2i32 == 0 {
                    n /= 2i32;
                }
                n == 1i32
            }
        }
    }
}

/// Implements `base_pow` function using `core::num::pow(base, exp)`.
///
/// It returns the power of a base i32 integer to an exponent u32 integer.
///
/// # Examples
///
/// ```
///  use std::num::pow
///  assert_eq!(base_pow(2i32, 10u32), pow(2i32, 10u32));
/// ```
fn base_pow(base: i32, exp: u32) -> i32 {
    base.pow(exp)
}

#[cfg(test)]
mod tests {
    use super::*;

    /// Test `is_power_of_two` function.
    ///
    /// # Panics
    ///
    /// Panics if .
    #[test]
    fn test_is_power_of_two() {
        let exp = 4u32;
        let base = 2i32;

        let pow: i32 = base_pow(base, exp);
        assert!(Solution::is_power_of_two(pow), "2^{} = {}", exp, pow);
    }

    /// Test `test_base_pow` function.
    ///
    /// # Panics
    ///
    /// Panics if .
    #[test]
    fn test_base_pow() {
        let exp = 10u32;
        let base = 2i32;

        let pow: i32 = base_pow(base, exp);
        assert!(pow == (2i32).pow(10u32), "2^{} = {}", exp, pow);
    }
}