package main
import (
	"fmt"
)
func replaceSpace(s string) string {
	srcBytes := []byte(s)
	dstBytes := make([]byte,0)
	for _,c:=range srcBytes {
		if c == ' ' {
			dstBytes =  append(dstBytes,'%','2','0')
		}else {
			dstBytes =  append(dstBytes,c)
		}
	}
	return string(dstBytes)
}
func main() {
	s :="We are happy."
	fmt.Println(replaceSpace(s))
}