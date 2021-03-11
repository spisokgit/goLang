// 10.3 slicing
package main

import "fmt"

func main() {
	data := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("data=", data)
	fmt.Println("len %d, cap %d", len(data), cap(data))

	data1 = data[1:4]
	fmt.Println("data[1:4]  len %d, cap %d", len(data), cap(data))

	data4 := data[:3]
	fmt.Println("data[:3]  len %d, cap %d", len(data), cap(data))

}
