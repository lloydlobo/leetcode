//! # Module solution for 231. Power of Two
//!
//! Given an integer n, return true if it is a power of two. Otherwise, return false.
//! An integer n is a power of two, if there exists an integer x such that n == 2x.
//!
//! # Examples
//!
//! ```
//! use solution::*
//!
//! assert!(is_power_of_two(1), "true");
//! assert!(is_power_of_two(16), "true");
//! assert!(!is_power_of_two(3), "false");
//! ```

/// Implements `IsPowerOfTwo` struct.
#[derive(Debug)]
pub struct IsPowerOfTwo {
    pub is_power_of_two: fn(i32) -> bool,
    pub is_power_of_two_ceil_floor: fn(i32),
}

impl IsPowerOfTwo {
    /// Implements `is_power_of_two_ceil_floor` where given an integer n, return true if
    /// the given number is a power of two.
    ///
    /// * Time Complexity : O(logn)
    /// * Space Complexity : O(1)
    ///
    /// # Examples
    ///
    /// ```
    /// assert!(Solution::is_power_of_two_ceil_floor(2), "2 is a power of two");
    /// assert!(Solution::is_power_of_two_ceil_floor(16), "16 is a power of two");
    /// ```
    pub fn ceil_floor(n: i32) -> bool {
        let n: f32 = (n as f32).log2();
        f32::ceil(n) == f32::floor(n)
    }

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
    pub fn modulus_loop(n: i32) -> bool {
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

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_is_power_of_two() {
        let result = IsPowerOfTwo::modulus_loop(2);
        assert!(result, "2 is a power of two");

        let result = IsPowerOfTwo::modulus_loop(16);
        assert!(result, "16 is a power of two");

        let result = IsPowerOfTwo::modulus_loop(3);
        assert!(!result, "3 is not a power of two");
    }

    #[test]
    fn test_is_power_of_two_ceil_floor() {
        let result = IsPowerOfTwo::ceil_floor(2);
        assert!(result, "2 is a power of two");

        let result = IsPowerOfTwo::ceil_floor(16);
        assert!(result, "16 is a power of two");

        let result = IsPowerOfTwo::ceil_floor(3);
        assert!(!result, "3 is not a power of two");
    }
}