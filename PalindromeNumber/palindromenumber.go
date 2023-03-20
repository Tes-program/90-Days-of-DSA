package main

import "fmt"

// Link: https://leetcode.com/problems/palindrome-number
// Run: go run PalindromeNumber/palindromenumber.go

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	num := x
	rev := 0
	for x != 0 {
		rev = rev*10 + x%10
		x = x / 10
	}
	return rev == num
}

func main() {
	fmt.Println(isPalindrome(121))  // true
	fmt.Println(isPalindrome(-121)) // false
	fmt.Println(isPalindrome(10))   // false
}
