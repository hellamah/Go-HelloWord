package main

import "fmt"

func mediaVariatica(numeros ...float64) float64 {
	total := 0.0

	for _, num := range numeros {
		total += num
	}

	return total / float64(len(numeros))
}

func imprimirAprovados(aprovados ...string) {
	fmt.Println(("Lista de Aprovados"))
	for i, aprovado := range aprovados {
		fmt.Printf("%d) %s \n", i+1, aprovado)
	}

}
