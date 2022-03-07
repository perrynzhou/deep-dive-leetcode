package main

import "fmt"

func reverseOnlyLetters(s string) string {
	res := make([]byte,len(s))
	for i:=0;i<len(s);i++ {
		res[i]  = s[i]
	}
	i,j := 0,len(res)-1
	for i <j {
		match1 := res[i]>='A' && res[i]<='Z'|| res[i] >='a'&&res[i]<='z'
		match2 :=res[j]>='A' && res[j]<='Z'|| res[j] >='a'&&res[j]<='z'
		if match1 && match2{
			 c := res[i]
			 res[i]=res[j]
			 res[j]=c
			 i++
			 j--
		}
		if !match1 {
			i++
		}
		if !match2 {
			j--
		}

	}
	return string(res)
}
func main() {
	fmt.Println(reverseOnlyLetters("ab-cd"))
	fmt.Println(reverseOnlyLetters("a-bC-dEf-ghIj"))
	fmt.Println(reverseOnlyLetters("7_28]"))

}
