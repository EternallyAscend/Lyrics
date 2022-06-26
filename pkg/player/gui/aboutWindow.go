package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type AboutWindow struct {
	window fyne.Window
}

func GenerateAboutWindow(app fyne.App) *AboutWindow {
	aboutWindow := &AboutWindow{
		window: app.NewWindow(AboutWindowTitle),
	}
	aboutWindow.window.Resize(fyne.NewSize(AboutWindowX, AboutWindowY))

	content := container.NewWithoutLayout()
	content.Resize(fyne.NewSize(AboutWindowX, AboutWindowY))
	aboutWindow.window.SetContent(content)

	info := widget.NewLabel(AboutInfo)
	content.Add(info)

	fmodImage := canvas.NewImageFromFile(FmodLogoPath)
	content.Add(fmodImage)
	fmodImage.FillMode = canvas.ImageFillOriginal
	fmodImage.Resize(fyne.NewSize(FmodLogoX, FmodLogoY))
	fmodImage.Move(fyne.NewPos(0, 0))

	return aboutWindow
}

func (that *AboutWindow) Show() {

}
