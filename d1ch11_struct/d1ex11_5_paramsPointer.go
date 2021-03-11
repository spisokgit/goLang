// 11.5 구조체 메소드 매개변수(포인터사용, 미사용) 사용법
package main

import "fmt"

type Vertex struct{ X, Y int }

func (Vert *Vertex) mula(scal int) int {
	Vert.X *= scal
	Vert.Y *= scal
	return Vert.X * Vert.Y
}
func (Vert Vertex) mulb(scal int) int {
	Vert.X *= scal
	Vert.Y *= scal
	return Vert.X * Vert.Y
}

func main() {
	vert := Vertex{10, 20}
	fmt.Println("mula()", vert.mula(10)) // 20000    <== struct 변경됨
	fmt.Println("mulb()", vert.mulb(10)) // 2000000  <== struct 변경되지 않고

}
