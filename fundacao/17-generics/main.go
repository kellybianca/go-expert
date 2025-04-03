package main

import "fmt"

type MyNumber int

type Number interface {
	~int | ~float64 // o ~ considera outros tipo, como exempo o tipo do mynumber
}

func Soma[T Number](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

func Compara[T comparable](a T, b T) bool {
	if a == b {
		return true
	}
	return false
}

func main() {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	m2 := map[string]float64{"a": 1.5, "b": 2.3, "c": 3.7}
	fmt.Println(Soma(m))
	fmt.Println(Soma(m2))
	m3 := map[string]MyNumber{"a": 1, "b": 2, "c": 3}
	fmt.Println(Soma(m3))
	println(Compara(10, 10))
}
