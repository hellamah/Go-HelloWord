package main

import "fmt"

func condicoes(nota float64) {
	if nota == 7 {
		fmt.Println("(cuidado com sua nota) Aprovado com nota", nota)
	} else if nota > 7 {
		fmt.Println("Aprovado com nota", nota)
	} else {
		fmt.Println("Aprovado com nota", nota)
	}
}
