package main

func main() {
	a := 10
	b := 2
	c := 3

	switch a {
	case 1:
		println("a is 1")
	case 2:
		println("a is 2")
	case 3:
		println("a is 3")
	default:
		println("a is not 1 or 2")
	}

	if a == b {
		println("a is equal to b")
	} else if a == c {
		println("a is equal to c")
	} else {
		println("a is not equal to b or c")
	}
}
