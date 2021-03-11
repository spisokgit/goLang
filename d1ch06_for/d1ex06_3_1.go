package main

import "fmt"

func main() {
	s := []string{"sun", "mon", "tue", "wed", "thu", "fri", "sat"}
	for i, a := range s {
		fmt.Println(i, a)
	}
}

// range() // python enumrate() 와 같 은 기능
