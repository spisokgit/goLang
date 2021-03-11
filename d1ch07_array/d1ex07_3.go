// 7.3 문자열 배열 [...] 가변길이 + slicing
//문자열 배열 [4] 고정길이 + slicing 가능
package main

import . "fmt"

func main() {
	a := [...]string{"test", "123", "simpletest", "abc"}[1]
	b := [...]string{"test", "123", "simpletest", "abc"}[2]
	c := [...]string{"test", "123", "simpletest", "abc"}[3]
	Println(a)
	Println(b)
	Println(c)
}
