// 11.7 구조체의 메소드 오버로딩
package main

import "fmt"

type Point struct{ X, Y int }
type Line struct {
	Point
	name string
	id   int
}

//  매개변수로 p *point
func (p *Point) say() {
	fmt.Println("Say Point")
	return
}

//  매개변수로 p *Line
func (p *Line) say() {
	fmt.Println("Say Line")
	return
}

func main() {
	var p Line
	p.say()       // Say Line
	p.Point.say() // Say Point
}
