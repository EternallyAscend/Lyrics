package mainWindow

import (
	"fyne.io/fyne/v2/widget"
	"time"
)

type StatusBar struct {
	object *widget.Label
}

func GenerateStatusBar() *StatusBar {
	return &StatusBar{}
}

func (that *StatusBar) Combine() {
	that.object = widget.NewLabel(time.Now().String()[0:19] + " Ready.")
}

func (that *StatusBar) UpdateStatus(status string) {
	that.object.SetText(time.Now().String()[0:19] + " " + status)
}
