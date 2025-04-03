package matematica

// todas as funcoes que sao exportadas para fora tem que iniciar com a inicial maiuscula,
// isso serve para struct, valores dentro da struct e metodos
func Soma[T int | float64](a T, b T) T {
	return a + b
}
