// 11.1 구조체
package main

import "fmt"

type Vertex struct {
	x int
	y int
}

func main() {
	v := Vertex{1, 2}
	v.x := 4
	fmt.Println(v)
}
