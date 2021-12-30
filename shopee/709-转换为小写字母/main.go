package main

import "fmt"

func toLowerCase(s string) string {
	sBytes := []byte(s)
	count := len(sBytes)

	for i:=0;i<count;i++ {
		if sBytes[i] < 'A' || sBytes[i] > 'z' {
			continue
		}
		if sBytes[i] >='A' && sBytes[i] <='Z' {
			sBytes[i] +=32
		}
	}
	return string(sBytes)
}
func main() {
	fmt.Println(toLowerCase("Hello@[]"))
}
