package main

import "fmt"

func main() {
	salarios := map[string]int{"Kelly": 1000, "Jose": 5000, "Maria": 2000}
	fmt.Println(salarios["Kelly"])
	delete(salarios, "Kelly")
	salarios["Kel"] = 5000 // adiciona um novo salario
	fmt.Println(salarios["Kel"])

	// 	sal := make(map[string]int) // cria um mapa inicial
	// sal1 := map[string]int{}
	// sal1["Kelly"] = 1000

	for nome, salario := range salarios {
		fmt.Printf("O salario de %s eh %d\n", nome, salario)
	}
	for _, salario := range salarios {
		fmt.Printf("O salario eh %d\n", salario)
	}
}
