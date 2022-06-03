package cc

/*
#cgo CFLAGS: -I./inc
#cgo LDFLAGS: -L./lib -lfmod -lfmodL
#include <stdio.h>
#include <stdlib.h>
#include "./music.h"
#include "./entry.h"
*/
import "C"

import (
	"fmt"
	"unsafe"
)

func TestCTransfer() {
	C.testFmod()
	C.testGoC()
	res := C.testGoCString()
	fmt.Println(C.GoString(res))
	defer C.free(unsafe.Pointer(res))
}

func LaunchFMOD() {
	C.launchFMOD()
}

func SetMediaFMOD(path string) {
	cstr := C.CString(path)
	C.setMediaFMOD(cstr)
	defer C.free(unsafe.Pointer(cstr))
}

func PlayFMOD() {
	C.playFMOD()
}

func ExitFMOD() {
	C.exitFMOD()
}
