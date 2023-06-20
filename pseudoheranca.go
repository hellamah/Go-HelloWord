package main

import (
	"fmt"
)

type carro struct {
	nome            string
	velocidadeAtual int
}
type ferrari struct {
	carro       //campo anonimo
	turboLigado bool
}

func (f ferrari) tipoHeranca() {

	fmt.Printf("A ferrari %s está com turbo ligado? %v\n", f.nome, f.turboLigado)

}
