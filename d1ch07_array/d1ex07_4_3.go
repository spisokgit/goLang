package main

import . "fmt"

func main() {
	var c []int
	var d [][]int
	c = []int{1, 11, 111, 1111}
	d = [][]int{{1, 11, 111},
		{2, 22, 222},
		{3, 33, 333},
	}
	Println(c, c[0], c[1])
	Println(d)
	Println("d[0], [d1], d[1][1]", d[0], d[1], d[1][1])
}
