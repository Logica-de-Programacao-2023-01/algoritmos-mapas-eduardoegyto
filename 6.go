package main

import "fmt"

func somarContagens(contagens []map[string]int) map[string]int {
	resultado := make(map[string]int)

	for _, contagem := range contagens {
		for palavra, quantidade := range contagem {
			resultado[palavra] += quantidade
		}
	}

	return resultado
}

func main() {
	contagens := []map[string]int{
		{"palavra1": 2, "palavra2": 3},
		{"palavra2": 5, "palavra3": 1},
		{"palavra1": 1, "palavra3": 4},
	}

	resultado := somarContagens(contagens)
	fmt.Println(resultado)
}
