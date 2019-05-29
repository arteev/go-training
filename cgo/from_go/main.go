package main

/*
 #cgo CFLAGS: -I.
 #cgo LDFLAGS: -L. -lsimple
 #include <stdio.h>
 #include <stdlib.h>
 #include "simple.h"
*/
import "C"

import (
	"fmt"
	"unsafe"
)

func main() {
	i := 1000
	for i < 1005 {
		i = int(C.inc(C.int(i)))
		cs := C.CString(fmt.Sprintf("From library: %d", i))
		C.print(cs)
		C.free(unsafe.Pointer(cs))
	}
}
