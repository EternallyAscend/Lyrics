package main

import (
	"lyrics/pkg/player/gui"
	"os"
)

func main() {
	os.Setenv(gui.FontFyne, gui.FontPath)
	cli := gui.GeneratePlayerClient()
	cli.Start()
	defer cli.Stop()
	defer os.Unsetenv(gui.FontFyne)
}
