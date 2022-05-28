package main

import (
	"lyrics/pkg/lyricsMaker"
)

func main() {
	lyricsMaker.NewLyricsMakerClient().Start()
}
