// 12.4 Local variable(total) return
package main

func main() {
	msg("how", "much", "is", "this?")
	msg("10dallars!")
}
func msg(mesg ...string) {
	var mes string
	for _, ms := range mesg {
		println(ms)
		mes = mes + ms
	}
	return mes //howmuchisthis?
}
