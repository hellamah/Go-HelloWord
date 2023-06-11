package main

type produto struct {
	nome     string
	preco    float64
	desconto float64
}

// Função com receiver (receptor)
func (p produto) precoComDesconto() float64 {
	return p.preco * (1 - p.desconto)
}
