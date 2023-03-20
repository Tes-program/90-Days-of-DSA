
"""https://leetcode.com/problems/reverse-integer/description/"""
class Solution:
    def reverse(self, x: int) -> int:
        sign = -1 if x < 0 else 1
        x = abs(x)
        result = 0
        while x > 0:
            result = result * 10 + x % 10
            x //= 10
        if result > 2**31 - 1:
            return 0
        return sign * result