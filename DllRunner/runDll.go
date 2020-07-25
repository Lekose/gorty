package main

import (
	"fmt"
	"log"
	"syscall"
)

func main() {
	h, e := syscall.LoadLibrary("libtest.dll")
	if e != nil {
		log.Fatal(e)
	}

	rand, err := syscall.GetProcAddress(h, "cgoRandom")
	if err != nil {
		log.Fatal(err)
	}

	n, _, _ := syscall.Syscall9(uintptr(rand), 0, 2, 2, 2, 2, 2, 0, 0, 0, 0)

	fmt.Printf("%x\n", h)
	fmt.Printf("%x\n", &rand)
	fmt.Printf("Function returns %d\n", n)
}
