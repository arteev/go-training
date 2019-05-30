package main

/*
#cgo CFLAGS: -I.
#cgo LDFLAGS: -L. -lclibrary
#include <stdio.h>
#include <stdlib.h>
#include "clibrary.h"

int callback_cgo(int i);
*/
import "C"
import (
	"fmt"
	"unsafe"
)

//export callback
func callback(arg int) int {
	fmt.Printf("Callback: arg = %d\n", arg)
	return arg + 1
}

func main() {
	fmt.Printf("calling C function with callback\n")
	C.c_func((C.callback_func)(unsafe.Pointer(C.callback_cgo)))
}
