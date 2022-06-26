package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type AboutWindow struct {
	parent  fyne.Window
	content *fyne.Container
}

func GenerateAboutWindow(window fyne.Window) *AboutWindow {
	aboutWindow := &AboutWindow{
		parent: window,
	}
	aboutWindow.content = container.NewWithoutLayout()
	aboutWindow.content.Resize(fyne.NewSize(AboutWindowX, AboutWindowY))

	info := widget.NewLabel(AboutInfo)
	aboutWindow.content.Add(info)

	fmodImage := canvas.NewImageFromFile(FmodLogoPath)
	aboutWindow.content.Add(fmodImage)
	fmodImage.FillMode = canvas.ImageFillOriginal
	fmodImage.Resize(fyne.NewSize(FmodLogoX, FmodLogoY))
	fmodImage.Move(fyne.NewPos(AboutFmodX, AboutFmodY))

	return aboutWindow
}

func (that *AboutWindow) CanvasContent() *fyne.Container {
	return that.content
}
