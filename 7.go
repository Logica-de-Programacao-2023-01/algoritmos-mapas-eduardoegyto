package main

import "fmt"

func contarLetras(frase string) map[string]map[rune]int {
	resultado := make(map[string]map[rune]int)
	palavras := splitPalavras(frase)

	for _, palavra := range palavras {
		contagem := make(map[rune]int)

		for _, letra := range palavra {
			contagem[letra]++
		}

		resultado[palavra] = contagem
	}

	return resultado
}

func splitPalavras(frase string) []string {
	return []string{"palavra1", "palavra2", "palavra3"}
}

func main() {
	frase := "Olá mundo, como vai?"
	resultado := contarLetras(frase)
	fmt.Println(resultado)
}
