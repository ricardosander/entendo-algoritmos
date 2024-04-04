package main

import "fmt"

func main() {
	pesquisa_binaria([]int{1, 3, 5, 7, 9}, 1)
	pesquisa_binaria([]int{1, 3, 5, 7, 9}, -1)
}

func pesquisa_binaria(lista []int, item int) {
	fmt.Printf("Pesquisa Binária em %v procurando por %d\n", lista, item)
	retorno, tentativas := algoritmo_pesquisa_binaria(lista, item)
	if retorno != nil {
		fmt.Printf("Encontrado na posição %d com %d tentativas\n", *retorno, tentativas)
	} else {
		fmt.Printf("Não encontrado com %d tentativas\n", tentativas)
	}
}

func algoritmo_pesquisa_binaria(lista []int, item int) (*int, int) {
	baixo := 0
	alto := len(lista) - 1
	tentativas := 0
	for baixo <= alto {
		tentativas++
		meio := (baixo + alto) / 2
		if lista[meio] == item {
			return &meio, tentativas
		}

		if item > lista[meio] {
			baixo = meio + 1
		}

		if item < lista[meio] {
			alto = meio - 1
		}

	}
	return nil, tentativas
}
