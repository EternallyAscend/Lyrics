package lyricsMaker

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"lyrics/pkg/lyricsMaker/config"
	"lyrics/pkg/lyricsMaker/mainWindow"
)

type LyricsMaker struct {
	daemon     fyne.App
	mainWindow *mainWindow.MainWindow
}

func NewLyricsMakerClient() *LyricsMaker {
	config.LoadConfig()
	application := &LyricsMaker{
		daemon: app.New(),
	}
	application.mainWindow = mainWindow.GenerateMainWindow(application.daemon, config.ApplicationTitle)
	return application
}

func (that *LyricsMaker) Start() {
	that.mainWindow.Load()
	defer that.Close()
}

func (that *LyricsMaker) UpdateStatus(status string) {
	that.mainWindow.UpdateStatus(status)
}

func (that *LyricsMaker) Close() {
	config.Close()
}
