package main

import "fmt"

func main() {
	f1 := func() {
		fmt.Println("Anonymous function")
	}

	f2 := func(a string) {
		fmt.Println(a)
	}

	f3 := func(a, b int) int {
		return a + b
	}
	f1()
	f2("Parameter function")
	fmt.Println(f3(5, 10))
}
