//7.1 문자열 배열 [] 고정길이
package main

import . "fmt"

func main() {
	var c [2]string
	var d []string
	c[0] = "first"
	c[1] = "second"
	d = []string{
		"aa",
		"aaa",
		"aaaaa"}
	Println(c, c[0], c[1], c[0][0]) // [first second] first second 102
	Println(d)
	Println((d[0]))                              // [aa aaa aaaaa]  // aa
	Println("d[0][0],d[0][1]", d[0][0], d[0][1]) // d[0][0],d[0][1] 97 97
}
