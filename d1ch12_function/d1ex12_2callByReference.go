// 12.2 function (call by Reference)
package main

import "fmt"

func main() {
	var message string = "hello"
	say(&message)        // 주소 변수 &(임퍼선트)
	fmt.Println(message) // changed
}
func say(message *string) { // main내 local variable과 같은 이름이 있다고 할때 // *포인터 타입으로 선언
	fmt.Println(message) // hello
	*message = "changed" // *message는 포인터변수 message에 대한 대입값 처리를 위한 것
}
