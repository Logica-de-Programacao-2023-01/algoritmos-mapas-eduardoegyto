package main

import "fmt"

func contarPares(slice []int) map[[2]int]int {
	resultado := make(map[[2]int]int)

	for i := 0; i < len(slice)-1; i++ {
		for j := i + 1; j < len(slice); j++ {
			pair := [2]int{slice[i], slice[j]}
			resultado[pair]++
		}
	}

	return resultado
}

func main() {
	slice := []int{1, 2, 3, 2, 4, 1, 5}
	resultado := contarPares(slice)
	fmt.Println(resultado)
}
