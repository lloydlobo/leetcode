pub struct Pow {
    pub base_pow: fn(i32, u32) -> i32,
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