package main

import (
	"log"
	"lyrics/pkg/lyricsMaker"
	"lyrics/pkg/player/cc"
	"lyrics/pkg/player/fmod"
)

func main() {
	player := fmod.GeneratePlayerFMOD()
	err := player.LoadMedia("./assert/test.wav")
	if nil != err {
		log.Println("Load err: ", err)
	}
	go player.Play()
	defer player.Close()
	cc.TestCTransfer()
	lyricsMaker.NewLyricsMakerClient().Start()
}
