//! Module helpers for creating helper functions.

pub struct Pow {
    pub base: fn(i32, u32) -> i32,
}

impl Pow {
    /// Implements `base` function using `core::num::pow(base, exp)`.
    ///
    /// It returns the power of a base i32 integer to an exponent u32 integer.
    ///
    /// # Examples
    ///
    /// ```
    ///  use std::num::pow
    ///  assert_eq!(base_pow(2i32, 10u32), pow(2i32, 10u32));
    /// ```
    pub fn base(base: i32, exp: u32) -> i32 {
        base.pow(exp)
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    /// Computes the power of a base number.
    ///
    /// # Examples
    ///
    /// ```
    /// let base = 2;
    /// let exp = 10;
    /// let pow = Pow::base(base, exp);
    /// assert_eq!(1024, pow);
    /// ```
    fn test_base_pow() {
        let exp = 10u32;
        let base = 2i32;

        let pow = Pow::base(base, exp);
        assert_eq!(1024, pow);
    }

    use crate::helpers;

    #[test]
    /// This function tests the base_pow_exp function
    ///
    /// # Examples
    ///
    /// ```
    /// use helpers::Pow;
    /// let exp = 10u32;
    /// let base = 2i32;
    /// assert_eq!(1024, Pow::base(base, exp));
    /// ```
    fn test_base_pow_exp() {
        let exp = 10u32;
        let base = 2i32;
        assert_eq!(1024, helpers::Pow::base(base, exp));
    }
}