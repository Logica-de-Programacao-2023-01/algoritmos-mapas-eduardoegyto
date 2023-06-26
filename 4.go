package main

import (
	"fmt"
	"sort"
	"strings"
)

func anagramas(palavras []string) map[string][]string {
	anagramas := make(map[string][]string)

	for _, palavra := range palavras {
		x := sortString(palavra)
		anagramas[key] = append(anagramas[x], palavra)
	}

	return anagramas
}

func sortString(str string) string {
	s := strings.Split(str, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func main() {
	palavras := []string{"amor", "roma", "arroz", "roda", "mora", "carro"}
	gruposAnagramas := anagramas(palavras)
	fmt.Println(gruposAnagramas)
