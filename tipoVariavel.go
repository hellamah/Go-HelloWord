package main

import (
	"fmt"
	"math"
)

func tipoVariavel() {
	fmt.Println("Esta é a função A.")

	const PI float64 = 3.1415
	var raio = 3.2

	//forma reduzida é a mais indicada
	area := PI * math.Pow(raio, 2)

	fmt.Println("A área da circunferência é ", area)

	const (
		a = 1
		b = 2
	)

	var (
		c = 3
		d = 4
	)

	fmt.Println(a, b, c, d)

	var e, f bool = true, false
	fmt.Println(e, f)

	g, h, i := 2, false, "epa"

	fmt.Println(g, h, i)

	//Maps e slices
	aprovados := make(map[int]string)

	aprovados[12345678978] = "Maria"
	aprovados[98765432100] = "Pedro"
	aprovados[95135745682] = "Ana"

	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	fmt.Println(aprovados[95135745682])
	delete(aprovados, 95135745682)

	fmt.Println(aprovados[95135745682])

}
