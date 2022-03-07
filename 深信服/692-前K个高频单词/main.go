package main

import "sort"

func topKFrequent(words []string, k int) []string {
	dict:= make(map[string]int)
	for _,word := range words {
		dict[word]++
	}
	type StringCnt struct {
		Val string
		Cnt  int
	}
	vec := make([]StringCnt,0)
	i := 0
	for k,v := range dict {
		vec[i] = StringCnt{k,v}
		i++
	}
	sort.Slice(vec,func(i,j int)bool{
		return vec[i].Cnt> vec[i].Cnt
	})
	result := make([]string,0)
	for i:=0;i<len(vec);i++ {
		result = append(result,vec[i].Val)
	}
	return result

}
