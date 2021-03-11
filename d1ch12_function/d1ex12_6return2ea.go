// 12.6 Local variable(total) return 2ea
package main

func main() {
	msg("how", "much", "is", "this?")
	msg("10dallars!")
}
func msg(mesg ...string) (mes string, index int) { // 여기에 return variable 2ea 적시
	for _, ms := range mesg {
		println(ms)
		mes = mes + ms
	}
	index = len(mesg)
	return // return에 사용하지 않고
}
