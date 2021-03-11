package main

const (
	num0, num1 = iota, 10 * iota
	num2, num3
	num4, num5
)

func main() {
	println(num0, num1, num2, num3, num4, num5)
}
