package toolBar

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type Figure struct {
	figure *canvas.Image
}

func NewFigureAndResize(path string, x float32, y float32) widget.ToolbarItem {
	image := canvas.NewImageFromFile(path)
	image.FillMode = canvas.ImageFillOriginal
	image.ScaleMode = canvas.ImageScaleSmooth
	image.SetMinSize(fyne.NewSize(x, y))
	image.Resize(fyne.NewSize(x, y))
	image.Refresh()
	image.Show()
	return &Figure{figure: image}
}

func NewFigure(path string) widget.ToolbarItem {
	image := canvas.NewImageFromFile(path)
	image.FillMode = canvas.ImageFillOriginal
	image.ScaleMode = canvas.ImageScaleSmooth
	image.Refresh()
	image.Show()
	return &Figure{figure: image}
}

func (that *Figure) ToolbarObject() fyne.CanvasObject {
	return container.NewBorder(nil, nil, nil, nil, that.figure)
}
