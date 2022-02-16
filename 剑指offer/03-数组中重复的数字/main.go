package main

import (
	"fmt"
	"sort"
)

func findRepeatNumber_v1(nums []int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	var i, j int
	for {
		j = i + 1
		if j >= len(nums) {
			break
		}
		if nums[i] == nums[j] {
			return nums[i]
		}
		i++
	}
	return -1
}
func findRepeatNumber_v2(nums []int) int {

	i := 0
	for {
		if i >= len(nums) {
			break
		}
		for k := len(nums) - 1; k > i; k-- {
			if nums[k] == nums[i] {
				return nums[k]
			}
		}
		i++
	}
	return -1
}
func findRepeatNumber_v3(nums []int) int {
	count := make(map[int]int)
	repeat :=0
	for _,v := range nums {
		count[v]++
		if count[v] > 1 {
			repeat = v
			break
		}
	}
	return repeat

}
func main() {
	fmt.Println(findRepeatNumber_v1([]int{2, 3, 1, 0, 2, 5, 3}))
	fmt.Println(findRepeatNumber_v2([]int{2, 3, 1, 0, 8, 5, 3}))
}
