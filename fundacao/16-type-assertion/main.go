package main

import "fmt"

func main() {
	var minhaVar interface{} = "Kelly Silva"
	println(minhaVar.(string))
	res, ok := minhaVar.(int)
	fmt.Printf("O valor de res é %v e o resultado de ok é %v", res, ok)
	// res := minhaVar // dar um erro pois n\o tem ok para retornar corretamente
	// fmt.Printf("O valor de res é %v", res)
}
