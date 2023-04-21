package main

import (
	"math/rand"
	"time"
)

func ifComIni() int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return r.Intn(10)
}
