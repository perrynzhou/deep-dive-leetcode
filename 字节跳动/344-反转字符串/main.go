package main

import "fmt"

func reverseString(s []byte)  {
	fmt.Println(string(s))
	i,j := 0,len(s)-1
	for i<j {
		 c := s[i]
		 s[i]=s[j]
		 s[j] =c
		 i++
		 j--
	}
	fmt.Println(string(s))
}
func main() {
	reverseString([]byte("abcde"))
}

