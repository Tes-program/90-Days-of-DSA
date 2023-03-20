package main

import "fmt"

// Link: https://leetcode.com/problems/roman-to-integer
// Run: go run RomanToInteger/main.go

func romanToInt(s string) int {
	hash := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	sum := 0
	for i := 0; i < len(s); i++ {
		curr, _ := hash[string(s[i])]
		next := 0
		prev := 0
		if i < len(s)-1 {
			next, _ = hash[string(s[i+1])]
		}
		if i > 0 {
			prev, _ = hash[string(s[i-1])]
		}
		if curr < next {
			sum = sum + (next - curr)
		} else if curr > prev && prev != 0 {
			continue
		} else {
			sum += curr
		}
	}
	return sum
}

func main() {
	fmt.Println(romanToInt("III"))     // 3
	fmt.Println(romanToInt("LVIII"))   // 58
	fmt.Println(romanToInt("MCMXCIV")) // 1994
}
