package c

/*
#cgo CFLAGS: -I./inc
#cgo LDFLAGS: -L./lib -lfmod -lfmodL
#include <stdio.h>
#include <stdlib.h>
#include "./entry.c"
*/
import "C"

import (
	"fmt"
	"unsafe"
)

func TestCTransfer() {
	C.testGoC()
	res := C.testGoCString()
	fmt.Println(C.GoString(res))
	C.free(unsafe.Pointer(res))
	//C.freeString(res)
}
