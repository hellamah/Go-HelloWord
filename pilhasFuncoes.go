package main

import "runtime/debug"

func pilhaFuncao3() {
	debug.PrintStack()
}
func pilhaFuncao2() {
	pilhaFuncao3()
}
func pilhaFuncao1() {
	pilhaFuncao2()
}
