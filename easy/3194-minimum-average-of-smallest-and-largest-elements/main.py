import operator
from typing import List

"""
Alternatives:

    res = 100
    for a, b in zip(l, r): res = min(res, a + b)
    return res / 2
    or
    return 0.5 * min(a + b for a, b in zip(l, r))
"""


def minimum_average(nums: List[int]):
    n = len(nums) // 2
    s: List[int] = sorted(nums)
    l, r = s[:n], s[-n:][::-1]  # (reversed with ::-1)
    return min(map(operator.add, l, r)) * 0.5


def main():
    nums: List[int] = list(range(9))
    res = minimum_average(nums)
    print(res)


if __name__ == "__main__":
    main()
