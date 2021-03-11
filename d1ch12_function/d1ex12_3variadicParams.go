// 12.3 가변인자 함수
package main

func main() {
	msg("how", "much", "is", "this?")
	msg("10dallars!")
}
func msg(mesg ...string) { // 가변인자 string // ... map이면 가변인자 map
	for _, ms := range mesg { // _ index 무시
		println(ms)
	}
}
