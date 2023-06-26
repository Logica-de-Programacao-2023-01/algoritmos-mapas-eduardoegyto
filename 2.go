package main

import "fmt"

func combinacao(mapa1, mapa2 map[string]int) map[string]int {
	resultado := make(map[string]int)

	for chave, valor := range mapa1 {
		resultado[chave] = valor
	}

	for chave, valor := range mapa2 {
		resultado[chave] = valor
	}

	return resultado
}

func main() {
	mapa1 := map[string]int{"a": 1, "b": 2}
	mapa2 := map[string]int{"b": 3, "c": 4}
	resultado := combinacao(mapa1, mapa2)
	fmt.Println(resultado) 
}
