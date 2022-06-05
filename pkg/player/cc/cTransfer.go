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
	go C.playFMOD()
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

func GetLengthFMOD() uint32 {
	clength := C.getLengthFMOD()
	return uint32(clength)
}

func GetPlayingFMOD() bool {
	cplaying := C.getPlayingFMOD()
	return 0 != cplaying
}

func GetPositionFMOD() uint32 {
	cposition := C.getPositionFMOD()
	return uint32(cposition)
}

func SetPositionFMOD(ms uint32) {
	cms := C.uint(ms)
	C.setPositionFMOD(cms)
	log.Printf("CGO: Set position %d in fmod system.", ms)
}

func SetVolumeFMOD(volume float32) {
	cvolume := C.float(volume)
	C.setVolumeFMOD(cvolume)
	log.Printf("CGO: Set volume %.2f in fmod system.", volume)
}

func SetFrequencyFMOD(frequency float32) {
	cfrequency := C.float(frequency)
	C.setFrequencyFMOD(cfrequency)
	log.Printf("CGO: Set Frequency %.2f in fmod system.", frequency)
}

func GetPitchFMOD() float32 {
	cpitch := C.getPitchFMOD()
	return float32(cpitch)
}

func SetPitchFMOD(pitch float32) {
	cpitch := C.float(pitch)
	C.setPitchDspFMOD(cpitch)
	log.Printf("CGO: Set Pitch DPS %.2f in fmod system.", pitch)

}
