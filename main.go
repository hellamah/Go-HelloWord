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

}
