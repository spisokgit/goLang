//8.1.3 할당연산자를 사용한 map 정의
package main

import . "fmt"

func main() {
	var imap = map[int]string{
		0: "North",
		1: "East",
		2: "West",
		3: "South",
	}
	Println(imap)
	Println(imap[1], imap[2])
}
