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
      for j := i+1;j<=count;j++ {
          var tmpSub []int
          tmpSub = append(tmpSub,nums[i:j]...)
          fmt.Println("tmpSub=",tmpSub)
          for _,v :=range nums[j:] {
              var tmp []int
              tmp = append(tmp,tmpSub...)
              tmp=append(tmp,v)
              fmt.Println("tmp=",tmp)
              result=append(result,tmp)
          }
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
