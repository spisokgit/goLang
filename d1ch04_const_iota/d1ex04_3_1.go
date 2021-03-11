package main

const (
	num0, num1 = iota, 10 * iota
	num2, num3 = iota, 10 * iota
	num4, num5 = iota, 10 * iota
)

func main() {
	println(num0, num1, num2, num3, num4, num5)
}
