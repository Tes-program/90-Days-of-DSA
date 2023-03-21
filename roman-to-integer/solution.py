class Solution:
    def romanToInt(self, s: str) -> int:
        table = {
            "M" : 1000,
            "D" : 500,
            "C" : 100,
            "L" : 50,
            "X" : 10,
            "V" : 5,
            "I" : 1
        }

        num = 0

        s = s.replace("IV", "IIII") \
             .replace("IX", "VIIII") \
             .replace("XL", "XXXX") \
             .replace("XC", "LXXXX") \
             .replace("CD", "CCCC") \
             .replace("CM", "DCCCC")

        for lis in s:
            num += table[lis]
        
        return num