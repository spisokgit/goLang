package main // golang은 package단위로 관리되며, main package에서 실행
// 라이브러리 패키지 읽어 드림

// func main()은 main package안에 선언되는 프로그램 시작점이다.
func main() { // 중괄호는 반듯이 함수의 선언문과 같은 줄에 있어야 한다.
	print("semicolon;\n") // import "fmt"를 하지 않아도 print()가능  // print() 소문자로 사용하면 지역변수로 가능한 built 함수이다.
	print("semi")
	print(("colon;"))
}
