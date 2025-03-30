package main

import "fmt"

type ID int

var (
	e float64 = 1.2
	f ID      = 1
)

func main() {
	var meuArray [3]int
	meuArray[0] = 1
	meuArray[1] = 2
	meuArray[2] = 15

	fmt.Println(meuArray[len(meuArray)-1]) // valor da ultima posicao do array
	fmt.Println(len(meuArray) - 1)         // tamanho do array

	for i, v := range meuArray {
		fmt.Printf("O valor do indice eh %d e o valor eh %d\n", i, v)
	}
}
