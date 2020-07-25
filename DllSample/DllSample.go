package main

import "C"

import (
	"fmt"
	"math/rand"
)

//export cgoRandom
func cgoRandom(m C.int) C.int {
	return C.int(rand.Intn(int(m)))
}

//export sayHello
func sayHello() {
	fmt.Printf("Hello there!")
}

// Go screams without a main function defined
func main() {}
