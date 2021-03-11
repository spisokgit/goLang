// 8.1.4 make를 사용한 map 초기화
package main

import . "fmt"

var imap map[int]string

func main() {
	imap = make(map[int]string)
	imap[0] = "North"
	imap[1] = "East"
	imap[2] = "West"
	imap[3] = "South"
	Println(imap)
	Println(imap[1], imap[2])
}
