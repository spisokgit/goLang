//9.3.3 Defer, panic, recover func
package main

import (
	"fmt"
	"os")


func openFile(fn string) {
	defer func() {
		if r := recover(); r != nil{
			fmt.Println("OPEN ERROR", r, "recovered")
		}()
	_, err := os.Open(fn)
	if err != nil {
		fmt.Println("panic routine")
		panic(error)
	}
}
