package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	var i, j, subStringMaxCount int
	stringLen := len(s)

	for {
		if i >= stringLen {
			break
		}
		subStringDict := make(map[byte]byte)
		subStringDict[s[i]] = s[i]
		for j = i + 1; j < stringLen; j++ {
			if _, ok := subStringDict[s[j]]; !ok {
				subStringDict[s[j]] = s[j]
			} else {
				break
			}
		}
		if len(subStringDict) > subStringMaxCount {
			subStringMaxCount = len(subStringDict)
		}
		i++
	}
	return subStringMaxCount

}
func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("bbbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
	fmt.Println(lengthOfLongestSubstring(""))
}
