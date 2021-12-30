package main

import "fmt"

func subsets(nums []int) [][]int {
   result := make([][]int,0)
   count := len(nums)-1
   i :=0
   for {
      if i>count {
        break
      }

      for j:=i+1;j<=count;j++ {
        result=append(result,nums[i:j])
      }
      i++
   }
   fmt.Println("--------")
   for i:=0;i<=count;i++ {
     result=append(result,[]int{nums[i]})
   }
  result=append(result,nums,[]int{})
   return result
}
func main() {
  result := subsets([]int{1,2,3,4})
  for _,sub := range result {
    fmt.Println(sub)
  }
}
