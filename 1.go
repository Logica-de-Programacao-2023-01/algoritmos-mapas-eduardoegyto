package main

import (
	"fmt"
	"strings"
)

func palavras(str string) map[string]int {
	ocorrencias := make(map[string]int)
	palavras := strings.Fields(str)

	for _, palavra := range palavras {
		ocorrencias[palavra]++
	}

	return ocorrencias
}

func main() {
	texto := "O gato é preto, o cachorro é marrom, o pássaro é azul."
	ocorrencias :=  palavras(texto)
	fmt.Println(ocorrencias)
}
