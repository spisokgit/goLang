package main

import "fmt"

func main() {
	m := map[string]float64{
		"pi":3.1416,
		"e" : 2.71828
	}
	fmt.Println(m)
	for key, value := range m { //order not specified
		fmt.Println(key,value)
	}

// 추가해야 함
}
