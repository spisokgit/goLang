// 8.1.5 make를 사용한 map 초기화
package main

import . "fmt"

func main() {
	imap := make(map[int]string)
	imap = map[int]string{
		0: "North",
		1: "East",
		2: "West",
		3: "South",
	}
	Println(imap)
	Println(imap[1], imap[2])
}
