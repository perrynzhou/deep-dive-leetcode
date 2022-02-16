package main

import "fmt"

func findNumberIn2DArray(matrix [][]int, target int) bool {
	isFound := false
	if len(matrix) <= 0 {
		return isFound
	}
	i, j := 0, len(matrix[0])-1

	for {
		if i >= len(matrix) || j < 0 {
			break
		}
		if matrix[i][j] == target {
			isFound = true
			break
		} else if matrix[i][j] > target {
			j--
		} else {
			i++
		}

	}
	return isFound
}
func main() {
	matrix1 := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}
	matrix := [][]int{{-5}}
	fmt.Println(findNumberIn2DArray(matrix, -2))
	fmt.Println(findNumberIn2DArray(matrix1, 22))

}
