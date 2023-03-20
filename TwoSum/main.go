package main

import "fmt"

// Link: https://leetcode.com/problems/two-sum
// Run: go run TwoSum/main.go

func twoSum(nums []int, target int) []int {
	hash := make(map[int]int, 0)
	result := make([]int, 0)
	for i, v := range nums {
		second := target - v
		if _, ok := hash[second]; !ok {
			hash[v] = i
		} else {
			result = append(result, hash[second], i)
		}
	}
	return result
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9)) // [0,1]
	fmt.Println(twoSum([]int{3, 2, 4}, 6))      // [1,2]
	fmt.Println(twoSum([]int{3, 3}, 6))         // [0,1]
}
