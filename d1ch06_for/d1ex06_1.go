package main

import "fmt"

func main() {
Loop:
	for i := 1; i < 10; i++ {
		if i == 5 {
			break Loop
		}
		fmt.Println(i)
	}
	for j := 1; j < 10; j++ {
		if j == 5 {
			break Loop
		}
		fmt.Println(j)
	}
}

// error
