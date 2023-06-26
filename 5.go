package main

import (
	"fmt"
)

func conta(str string) map[string]int {
	frequencia := make(map[string]int)

	for _, char := range str {
		frequencia[string(char)]++
	}

	return frequencia
}

func main() {
	texto := "abracadabra"
	frequencia := conta(texto)
	fmt.Println(frequencia) 
  }
