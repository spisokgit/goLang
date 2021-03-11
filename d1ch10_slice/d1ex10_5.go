package main

import "fmt"

func main() {
	data := []int{2, 3, 5, 7, 11, 13}
	for i, v := range data { // python range, enumerate 생각
		fmt.Println(i, v)
	}
}
