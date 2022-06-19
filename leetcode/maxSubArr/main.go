package main

import "fmt"

func main() {
	nums := []int{-1, 2, -3, 4, -1, 2}
	fmt.Println(maxSubArr(nums))
}

func maxSubArr(nums []int) int {
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] + nums[i - 1] > nums[i] {
			nums[i] += nums[i - 1]
		}

		if nums[i] > max {
			max = nums[i]
		}
	}

	return max
}
