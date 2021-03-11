// 11. 6 구조체 메서드 임베딩
package main

import "fmt"

type Point struct{ X, Y int }
type Line struct {
	Point
	name string
	id   int
}

func (p *Point) mula(scal int) int {
	p.X *= scal
	p.Y *= scal
	return p.X * p.Y
}

func main() {
	var p Line
	p.Point.X = 2
	p.Point.Y = 3
	fmt.Println(p.Point.mula(10)) //
	fmt.Println(p.mula(10))       //

}
