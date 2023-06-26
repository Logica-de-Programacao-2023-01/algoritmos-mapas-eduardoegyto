package main

import "fmt"

func soma(valores map[string]int) int {
	soma := 0

	for _, valor := range valores {
		soma += valor
	}

	return soma
}

func main() {
	valores := map[string]int{"a": 1, "b": 2, "c": 3}
	soma := soma(valores)
	fmt.Println(soma) 
}
