package main

import "fmt"

// 双指针解法
func remove(s []int) int {
	fmt.Println(s)
	i,j :=0,0
	for {
		if j>=len(s) {
			break
		}
		if i<=0 && j<=0 {
			i++
			j++
		}
		if s[j]!=s[i-1] {
			if i!=j {
				s[i]=s[j]
			}
			i++
			j++
		}else {
			j++
		}
	}
	return i
}

func main() {
	s :=[]int{1,1,2,2,3,3,5,6,8}
	n := remove(s)
	fmt.Println(s[:n])
}