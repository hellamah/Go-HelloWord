package main

func fatorial(n uint) uint { //uint é um int sem sinal
	switch {
	case n == 0:
		return 1
	default:
		return n * fatorial(n-1)
	}
}
