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

	fmt.Printf("%p", h)
}
