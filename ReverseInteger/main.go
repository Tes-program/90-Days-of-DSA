package main

import (
	"fmt"
	"math"
)

// Link: https://leetcode.com/problems/reverse-integer/
// Run: go run ReverseInteger/main.go

func reverse(x int) int {
	rev := 0
	for x != 0 {
		rev = rev*10 + x%10
		x = x / 10
	}
	if rev > math.MaxInt32 || rev < math.MinInt32 {
		return 0
	}
	return rev
}
func main() {
	fmt.Println(reverse(123))  // 321
	fmt.Println(reverse(-123)) // -321
	fmt.Println(reverse(120))  // -21
}
