package main

var soma = func(a, b int) int {
	return a + b
}

func multiplicacao(a, b int) int {
	return a * b
}

func execuMultiplicacao(funcao func(int, int) int, var1, var2 int) int {
	return funcao(var1, var2)
}
