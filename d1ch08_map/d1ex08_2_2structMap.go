//8.2.2 struct를 사용한 map
package main

import . "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": Vertex{40.68433, -74.39967},
	"Google":    Vertex{37.42202, -122.08408},
}

func main() {
	Println(m)              // map[Bell Labs:{40.68433 -74.39967}]
	Println(m["Bell Labs"]) //{40.68433 -74.39967}
}
