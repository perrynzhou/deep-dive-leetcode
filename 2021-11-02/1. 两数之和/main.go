package main

import "fmt"

func twoSum(nums []int, target int) []int {
	dataIndex := make(map[int][]int)
	for index, val := range nums {
		if dataIndex[val]==nil {
			dataIndex[val] = make([]int,0)
		}
		dataIndex[val] = append(dataIndex[val],index)
	}
	var result []int
	for index,value := range nums {
		left := target - value
		done :=false
		if dataIndex[left]!=nil {
			for _,i := range dataIndex[left] {
				if i!=index {
					result = append(result,i)
					result = append(result,index)
					done=true
					break
				}
			}
		}
		if done {
			break
		}
	}
	return result
}
func main() {
	fmt.Println(twoSum([]int{2,7,11,15},9))
	fmt.Println(twoSum([]int{3,2,4},6))
	fmt.Println(twoSum([]int{3,3},6))
}
