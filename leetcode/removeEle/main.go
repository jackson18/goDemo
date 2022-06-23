package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(removeEle(nums, 3))
}

func removeEle(nums []int, val int) int {
	left := 0
	for _, v := range nums {
		if v != val {
			nums[left] = v
			left++
		}
	}
	return left
}
