package main

import "fmt"

func add(a, b int) int {
	return a + b
}

var a, b int = 1, 2 // global 변수의 선언 및 정의
func main() { // 실행프로그램의 진입점
	result := add(a, b)                                // 함수의 사용 // Short Assignment Statement ( := )  var 사용하지 않고 local variable
	fmt.Printf("Result=%d=add(%d + %d)", result, a, b) //fmt 패키지 사용하여 fmt.Printf 대문자로 시작 // package의 scope가 있다고
}

// thirt part git 바로 사용할 수 있다.
// get get git주소으로 가져 올 수 있다. 외부 프로그램 저장되는 장소가 있다.  ==> git 이용하여 협업 시스템 과 배포에 강점이 있다.
// get get git주소  import
