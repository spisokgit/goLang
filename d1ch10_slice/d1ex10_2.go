package main

import "fmt"

func main() {
	data := [6]int{2, 3, 5, 7, 11, 13}
	var s []int = data[1:4]
	fmt.Println(s, len(s), cap(s))
}
