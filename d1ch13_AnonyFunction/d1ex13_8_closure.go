package main

import "fmt"

func main() {
	sum := 0
	addr := func(x int) int {
		sum += x
		return sum
	}
	aum := 0
	neg := func(x int) int {
		aum += x
		return aum
	}
	for i := 0; i < 10; i++ {
		fmt.Println(addr(i), neg(-2*i))
	}
}
