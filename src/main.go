package main

import (
	"lyrics/pkg/lyricsMaker"
	"lyrics/pkg/player/c"
)

func main() {
	c.TestCTransfer()
	lyricsMaker.NewLyricsMakerClient().Start()
}
