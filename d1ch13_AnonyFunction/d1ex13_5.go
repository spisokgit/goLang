package main

import "fmt"

func main() {
	func() {
		fmt.Println("Anonymous function")
	}() // 익명함수의 실행

	func(a string) {
		fmt.Println(a)
	}("Parameter function") // 익명함수 매개변수로 실행

	res := func(a, b int) int {
		return a + b
	}(5, 10) // 익명함수 매개변수로 실행 ( return 값 변수res로 받음)
	fmt.Println(res) // 변수 res (익명함수 결과 )
}
