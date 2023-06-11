package main

import "fmt"

func main() {
	conversoes()
	condicoes(4)
	tipoVariavel()

	if i := ifComIni(); i > 5 {
		fmt.Println("Ganhou: ", i)
	} else {
		fmt.Println("Perdeu: ", i)
	}

	//Funções Basicas
	f1()
	f2("teste1", "teste2")

	r3, r4 := f3(), f4("teste3", "teste4")

	fmt.Printf(r3 + "\n")
	fmt.Printf(r4)

	r5, r6 := f5()

	fmt.Printf("\nF5: %s %s", r5, r6)

	//--------------------------------------------------------------
	//Pilha de Funções

	pilhaFuncao1()
	//--------------------------------------------------------------
	//Retorno Nomeado

	r7, r8 := trocar(10, 20)
	fmt.Println(r7)
	fmt.Println(r8)
	//--------------------------------------------------------------
	//primeiraClasse

	fmt.Println((soma(2, 3)))
	sub := func(a, b int) int {
		return a - b
	}

	fmt.Println((sub(2, 3)))

	//Funcao como Parametro
	fmt.Println(execuMultiplicacao(multiplicacao, 10, 20))
	//--------------------------------------------------------------
	//Funcoes Variaticas
	fmt.Printf("Média: %2.f", mediaVariatica(7.6, 7.1, 2.4, 9.1, 2.4))

	//Funcoes Variaticas com Slice
	imprimirAprovados([]string{"Teste1", "Teste2", "Teste3"}...)

	//--------------------------------------------------------------
	//Recursividade usando uint
	resultadoFat := fatorial(5)
	fmt.Println(resultadoFat)

	//--------------------------------------------------------------
	//Exemplo de Defer
	defer fmt.Println("Fim!")

}
