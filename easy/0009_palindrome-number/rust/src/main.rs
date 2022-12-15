/// Given an integer x, return true if x is a palindrome, and false otherwise.
fn main() {
    let result: bool = is_palindrome(4324234);
    dbg!(result);
}

/// Implements `is_palindrome` function to validate if a number is a palindrome.
///
/// Returns true if the given number is a palindrome.
///
/// Args:
///     x: An integer.
///
/// # Examples
///
/// ```
/// use leetcode::easy::0009_palindrome_number::is_palindrome;
/// assert!(is_palindrome(4324234), "4324234 is palindrome.");
/// assert!(is_palindrome(121i32), "121 should be a palindrome");
/// ```
///
/// # Panics
///
/// ```
/// assert!(is_palindrome(-121));
/// ```
///
/// Returns:
///     True if x is a palindrome, False otherwise.
pub fn is_palindrome(x: i32) -> bool {
    x.to_string() == x.to_string().chars().rev().collect::<String>()
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_is_palindrome() {
        assert!(is_palindrome(121i32), "121 should be a palindrome");
        assert!(!is_palindrome(10i32), "10 should not be a palindrome");
    }

    #[test]
    #[should_panic]
    fn test_is_palindrome_panics() {
        assert!(is_palindrome(-121), "-121 should not be a palindrome");
    }
}

/*
    Given an integer x, return true if x is a
    palindrome
    , and false otherwise.

    Example 1:

    Input: x = 121
    Output: true
    Explanation: 121 reads as 121 from left to right and from right to left.

    Example 2:

    Input: x = -121
    Output: false
    Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.

    Example 3:

    Input: x = 10
    Output: false
    Explanation: Reads 01 from right to left. Therefore it is not a palindrome.

    Constraints:

        -2^31 <= x <= 2^31 - 1


    Follow up: Could you solve it without converting the integer to a string?
*/
