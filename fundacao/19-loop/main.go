package main

func main() {
	// tem apenas o for
	for i := 0; i < 10; i++ {
		println(i)
	}

	numeros := []string{"um", "dois", "tres"}
	for k, v := range numeros {
		println(k, v)
	}

	i := 0
	for i < 10 {
		println(i)
		i++
	}

	for {
		println("loop infinito")
		break
	}
}
