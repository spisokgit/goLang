//9.3.2 Panic("Error message")
// 현재함수를 즉시 멈추고 현재함수에 defer함수들을 모두 실행한 후 즉시 리턴한다.
// 그리고 마지막에는 프로그램이 에러를 내고 종료하게 된다.
package main

import "os"

func openFile(fn string) {
	f, err := os.Open(fn)
	if err != nil {
		panic(err)
	}
	defer f.Close()
}

func main() {
	openFile("Invalid.txt")
	println("Done")
}
