//8.2.3 struct를 사용한 map의 삭제
package main

import "fmt"

type Vertex struct{ Lat, Long float64 }

var m = map[string]Vertex{
	"Bell Labs": Vertex{40.68433, -74.39967},
	"Google":    Vertex{37.42202, -122.08408},
}

func main() {
	m["test"] = Vertex{10.68433, -54.39967}
	//map append
	m["test1"] = Vertex{20.68433, -154.39967}
	for key, val := range m {
		fmt.Println(key, val)
	}
	// delete
	delete(m, "test")
	delete(m, "test1")
	for key, val := range m {
		fmt.Println(key, val)
	}
}
