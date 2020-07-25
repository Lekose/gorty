package main

import (
	"syscall"
)

var (
	h, _        = syscall.LoadLibrary("libtest.dll")
	sayhello, _ = syscall.GetProcAddress(h, "say_hello")
)

func sayHello() (result int) {
	var nargs uintptr = 0
	ret, _, callErr := syscall.Syscall9(uintptr(sayhello),
		nargs,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0)
	if callErr != 0 {

	}

	result = int(ret)
	return
}

func main() {
	defer syscall.FreeLibrary(h)

	sayHello()
}
