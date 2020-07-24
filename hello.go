package main

import "C"

import (
	"math/rand"
)

//export cgoRandom
func cgoRandom(m C.int) C.int {
	return C.int(rand.Intn(int(m)))
}
