package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	n, index, count := len(s), make([]int, 128), 0

	if n <= 0 {
		return n
	}
	if n <= 1 {
		return n
	}
	for i := 0; i < 128; i++ {
		index[i] = -1
	}
	var srcVal, dstVal int
	for i := 0; i < n; i++ {
		pos := s[i]
		srcVal += i
		if index[pos] < 0 {
			index[pos] = i
			dstVal +=i
		} else {
			diff_count := i - index[pos] 
			if diff_count > count {
				count = diff_count
			}
			index[pos] = i
		}
	}

	if srcVal == dstVal {
		return len(s)
	}
	return count

}
func lengthOfLongestSubstring_v2(s string) int {
	i, n, index := 0, len(s)-1, make([]int, 128)
	var subStringMaxCount int
	for ; i < n; i++ {
		pos := s[i]
		if index[pos] <= 0 {
			index[pos] = i
		} else {
			if i-index[pos] >= subStringMaxCount {
				subStringMaxCount = i - index[pos]
			}
			index[pos] = i
		}

	}

	return subStringMaxCount
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"), lengthOfLongestSubstring_v2("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("bbbb"), lengthOfLongestSubstring_v2("bbbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"), lengthOfLongestSubstring_v2("pwwkeww"))
	fmt.Println(lengthOfLongestSubstring(""), lengthOfLongestSubstring_v2(""))
	fmt.Println(lengthOfLongestSubstring("au"), lengthOfLongestSubstring_v2("au"))

	fmt.Println("------")
	//fmt.Println(lengthOfLongestSubstring_v2("abcabcbb"))
}
