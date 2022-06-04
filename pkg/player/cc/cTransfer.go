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
	"log"
	"unsafe"
)

func TestCTransfer() {
	res := C.testGoCString()
	log.Println(C.GoString(res))
	defer C.free(unsafe.Pointer(res))
}

func LaunchFMOD() {
	C.launchFMOD()
	log.Println("CGO: Launch fmod system.")
}

func SetMediaFMOD(path string) {
	cstr := C.CString(path)
	C.setMediaFMOD(cstr)
	defer C.free(unsafe.Pointer(cstr))
	log.Printf("CGO: Set media %s to fmod channel.", path)
}

func PlayFMOD() {
	C.playFMOD()
	log.Println("CGO: Playing with fmod channel.")
}

func PauseFMOD() {
	C.pauseFMOD()
	log.Println("CGO: Pausing with fmod channel.")
}

func ExitFMOD() {
	C.exitFMOD()
	log.Println("CGO: Exit fmod system.")
}
