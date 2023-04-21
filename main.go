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

}
