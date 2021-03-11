// 11.2 구조체 포인터
package main

import "fmt"

type Vertex struct {
	x int
	y int
}

func main() {
	v := Vertex{1, 2}
	p := &v
	p.x = 1e9 // (*p).x = 1e9 동일함
	fmt.Println(v)

}
