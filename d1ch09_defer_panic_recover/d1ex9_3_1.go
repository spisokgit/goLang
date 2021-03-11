//9.3.1 defer
// 종료(정상, error)할때 수행되는 함수 // file open 등에서 file close해야 하는데 사용
package main

import "fmt"

func main() {
	defer fmt.Println("world")
	fmt.Println("hello")
}
