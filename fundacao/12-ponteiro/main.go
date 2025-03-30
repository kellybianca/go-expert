package main

func main() {
	a := 10
	var ponteiro *int = &a
	println(ponteiro)
	*ponteiro = 20
	b := &a
	*b = 30
	println(ponteiro)
	println(a)

}
