// 12.1 function (call by value)
package main

import "fmt"

func main() {
	var message string = "hello"
	say(message)
	fmt.Println(message) // hello
}
func say(message string) { // main내 local variable과 같은 이름이 있다고 할때
	fmt.Println(message) // hello
	message = "changed"
}
