package main

import "fmt"

func main() {
	res := func(a, b int) int {
		return a + b
	}(1, 3)
	fmt.Println(res)
}
