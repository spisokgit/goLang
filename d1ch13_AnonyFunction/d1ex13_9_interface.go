package main

import "fmt"

type Itype int

func (i Itype) Print() {
	fmt.Println(i)
}

type Rectangle struct {
	width, height int
}

func (r Rectangle) Print() {
	fmt.Println(r.width, r.height)
}

type Printer interface {
	Print()
}

func main() {
	var i Itype = 5
	r := Rectangle{10, 20}
	var p Printer // 인터페이스 사용
	p = i
	p.Print() // Itype변수의 메서드 호출
	p = r
	p.Print() // Rectangle 메서드 호출
}
