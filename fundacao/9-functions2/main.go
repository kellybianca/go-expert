// Funcoes variadicas
// package main

// import (
// 	"fmt"
// )

// func main() {

// 	fmt.Println(sum(1, 2, 123, 23, 23, 12, 4, 32, 4, 35, 2, 4, 23, 25, 2))
// }

// func sum(numeros ...int) int {
// 	total := 0
// 	for _, numero := range numeros {
// 		total += numero
// 	}
// 	return total
// }

// funcoes anonimas ou closures
package main

import (
	"fmt"
)

func main() {
	total := func() int {
		return sum(1, 2, 123, 23, 23, 12, 4, 32, 4, 35, 2, 4, 23, 25, 2) * 2
	}()
	fmt.Println(total)
}
func sum(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}
