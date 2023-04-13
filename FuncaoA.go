package main

import (
	"fmt"
	"math"
)

func FuncaoA() {
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
}
