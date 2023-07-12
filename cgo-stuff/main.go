package main

// #cgo CFLAGS: -g -Wall
// #include <stdlib.h>
// #include <stdio.h>
import "C"

import (
	"fmt"
	"unsafe"
)

func main() {
	name := C.CString("name")
	defer C.free(unsafe.Pointer(name))

	year := C.int(2018)

	C.printf("%c", name)
	fmt.Println(year)
}
