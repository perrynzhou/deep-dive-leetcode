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
func lengthOfLongestSubstring_v2(s string) int {
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
			dstVal += i
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

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"), lengthOfLongestSubstring_v2("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("bbbb"), lengthOfLongestSubstring_v2("bbbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"), lengthOfLongestSubstring_v2("pwwkeww"))
	fmt.Println(lengthOfLongestSubstring(""), lengthOfLongestSubstring_v2(""))
	fmt.Println("*****")
	fmt.Println(lengthOfLongestSubstring("aa"), lengthOfLongestSubstring_v2("aa"))
	fmt.Println(lengthOfLongestSubstring("au"), lengthOfLongestSubstring_v2("au"))
	fmt.Println(lengthOfLongestSubstring("aab"), lengthOfLongestSubstring_v2("aab"))

	fmt.Println("------")
	//fmt.Println(lengthOfLongestSubstring_v2("abcabcbb"))
}
