package main

import (
	"fmt"
)

func multiply(num1 string, num2 string) string {
	if num1=="0" || num2=="0" {
		return "0"
	}
	tmp := make([]int,len(num1)+len(num2))
	for i:=len(num1)-1;i>=0;i-- {
		n := int(num1[i]-'0')
		for j:=len(num2)-1;j>=0;j-- {
			m := int(num2[j]-'0')
			tmp[i+j+1]  =tmp[i+j+1] + m *n
		}
	}
	for i:=len(tmp)-1;i>0;i-- {
		if tmp[i]>=10 {
			tmp[i-1]= tmp[i-1] +tmp[i]/10
			tmp[i]= tmp[i]%10
		}
	}
	res := ""
	var idx int
	if tmp[0]==0 {
		idx = 1
	}
	for ;idx<len(tmp);idx++ {
		res +=string(rune(tmp[idx]+'0'))
	}
	return res

}
/*
由于 \textit{num}_1num
1
​
  和 \textit{num}_2num
2
​
  的乘积的最大长度为 m+nm+n，因此创建长度为 m+nm+n 的数组 \textit{ansArr}ansArr 用于存储乘积。对于任意 0 \le i < m0≤i<m 和 0 \le j < n0≤j<n，\textit{num}_1[i] \times \textit{num}_2[j]num
1
​
 [i]×num
2
​
 [j] 的结果位于 \textit{ansArr}[i+j+1]ansArr[i+j+1]，如果 \textit{ansArr}[i+j+1] \ge 10ansArr[i+j+1]≥10，则将进位部分加到 \textit{ansArr}[i+j]ansArr[i+j]。

最后，将数组 \textit{ansArr}ansArr 转成字符串，如果最高位是 00 则舍弃最高位。

 */
func main() {
	// "498828660196"
	//"840477629533"
	fmt.Println(multiply("498828660196","840477629533"))
	// "419254329864656431168468"
	// "15177301459451591636"
}
