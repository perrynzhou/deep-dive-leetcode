package main

import (
	"fmt"
)

func PalindromeHelp(s string) bool {
	if len(s) < 2 {
		return true
	}
	i, j := 0, len(s)-1
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}
func longestPalindrome(s string) string {
	i, j := 0, len(s)
	var res string
	for i < j {
		left, right := i, j
		for left < right {
			ok := PalindromeHelp(s[left:right])
			if ok {
				if len(res) < len(s[left:right]) {
					res = s[left:right]
				}
			}
			right--
		}
		i++
	}
	return res
}

func main() {

	fmt.Println(longestPalindrome("babad"))

	fmt.Println(longestPalindrome("cbbd"))

	fmt.Println(longestPalindrome("ac"))

	fmt.Println(longestPalindrome("a"))
	fmt.Println(longestPalindrome("abb"))

}
