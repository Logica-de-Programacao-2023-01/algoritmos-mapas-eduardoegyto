package main

import "fmt"

func calcularDespesas(despesas map[string]float64) map[string]float64 {
	totalDespesas := 0.0
	for _, valor := range despesas {
		totalDespesas += valor
	}

	quantidadePessoas := len(despesas)
	valorPorPessoa := totalDespesas / float64(quantidadePessoas)

	resultado := make(map[string]float64)
	for pessoa, valor := range despesas {
		resultado[pessoa] = valor - valorPorPessoa
	}

	return resultado
}

func main() {
	despesas := map[string]float64{
		"João":  100.0,
		"Maria": 150.0,
		"Pedro": 120.0,
	}

	resultado := calcularDespesas(despesas)
	fmt.Println(resultado)
}
