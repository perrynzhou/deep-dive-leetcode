package main
import (
	"fmt"
)
func maxArea(height []int) int {
	i,j := 0,len(height)-1
	var maxArea  int 
	for ;i<j;i++ {
		min := height[j]
		if min >= height[i] {
			min = height[i]
		}
		area := min * (j-i)
		if area > maxArea {
			maxArea = area
		}
	}
	return maxArea
}
func main() {
	fmt.Println(maxArea([]int{1,8,6,2,5,4,8,3,7}))
	fmt.Println(maxArea([]int{1,1}))
	fmt.Println(maxArea([]int{14,3,2,1,4}))
	fmt.Println(maxArea([]int{1,2,1}))
}