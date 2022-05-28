package body

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type Body struct {
	Object fyne.CanvasObject
}

func GenerateBody() *Body {
	return &Body{}
}

func (that *Body) Combine() {
	that.Object = container.NewBorder(nil, nil, nil, nil, widget.NewLabel("Body!"))
}
