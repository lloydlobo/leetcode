//! 231: Power Of Two.
//!
//! Given an integer n, return true if it is a power of two. Otherwise, return false.
//! An integer n is a power of two, if there exists an integer x such that n == 2x.

mod solution;
mod helpers;

use solution::IsPowerOfTwo;

/// Implements `main` function entry point.
fn main() {
    dbg!(IsPowerOfTwo::modulus_loop(helpers::Pow::base(2i32, 4u32)));
    dbg!(IsPowerOfTwo::modulus_loop(3));
    dbg!(IsPowerOfTwo::ceil_floor(3));
}

#[cfg(test)]
mod tests {
    use super::*;

    /// Test `is_power_of_two` function.
    ///
    /// # Panics
    ///
    /// Panics if .
    /// # Examples
    ///
    /// ```
    /// use leetcode::base_pow;
    /// let exp = 4u32;
    /// let base = 2i32;
    /// let pow: i32 = base_pow(base, exp);
    /// assert!(solution::IsPowerOfTwo::modulus_loop(pow), "2^{} = {}", exp, pow);
    /// ```
    #[test]
    fn test_is_power_of_two() {
        let exp = 4u32;
        let base = 2i32;

        let pow: i32 = helpers::Pow::base(base, exp);
        assert!(solution::IsPowerOfTwo::modulus_loop(pow), "2^{} = {}", exp, pow);
    }

    /// Returns true if the given number is a power of two.
    ///
    /// # Examples
    ///
    /// ```
    /// use leetcode_rs::Solution;
    /// ```
    #[test]
    fn test_is_power_of_two_ceil_floor() {
        assert!(solution::IsPowerOfTwo::ceil_floor(2), "2 is a power of two");
        assert!(solution::IsPowerOfTwo::ceil_floor(16), "16 is a power of two");
        assert!(!solution::IsPowerOfTwo::ceil_floor(3), "3 is not a power of two");
        assert!(!solution::IsPowerOfTwo::ceil_floor(1200), "1200 is not a power of two");
    }

    #[test]
    /// Computes the power of a base number.
    ///
    /// # Examples
    ///
    /// ```
    /// let base = 2;
    /// let exp = 10;
    ///
    /// assert_eq!(1024, base_pow(base, exp));
    /// ```
    fn test_base_pow() {
        let exp = 10u32;
        let base = 2i32;

        let pow: i32 = helpers::Pow::base(base, exp);
        assert!(pow == (2i32).pow(10u32), "2^{} = {}", exp, pow);
    }
}