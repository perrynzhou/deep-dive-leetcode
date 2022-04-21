package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	res := make([]int, 0)
	dict := make(map[int][]int)
	for i, v := range nums {
		if dict[v] == nil {
			dict[v] = make([]int, 0)
		}
		dict[v] = append(dict[v], i)
		if dict[target-v] != nil {
			if len(dict[target-v]) > 1 {
				res = dict[target-v][:2]
			} else {
				if dict[v][0] != dict[target-v][0] {
					res = append(res, dict[v][0], dict[target-v][0])
					break
				}
			}
		}
	}
	/*
		for value, indexes := range dict {
			left := target - value
			if dict[left] != nil {
				if len(dict[left]) > 1 {
					res = indexes[:2]
				} else {
					if indexes[0] != dict[left][0] {
						res = append(res, indexes[0], dict[left][0])
						break
					}
				}

			}
		}

	*/
	return res
}
func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
	fmt.Println(twoSum([]int{3, 3}, 6))
}
