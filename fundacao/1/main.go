package main

const a = "Hello World"

var b bool
var (
	c bool    = true
	d int     = 10
	e string  = "Kelly"
	f float64 = 1.2
)

func main() {
	b = true
	println(b)
	println(c)
	println(d)
	println(e)
	println(f)

	a := "X" // string
	println(a)
}
