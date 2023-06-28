package main

import "fmt"

func gerarSequenciaFibonacci(n int) map[int]int {
	resultado := make(map[int]int)
	a, b := 0, 1

	for i := 0; i <= n; i++ {
		resultado[i] = a
		a, b = b, a+b
	}

	return resultado
}

func main() {
	n := 10
	resultado := gerarSequenciaFibonacci(n)
	fmt.Println(resultado)
}
