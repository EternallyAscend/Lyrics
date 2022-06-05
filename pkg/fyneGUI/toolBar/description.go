package toolBar

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type TextDescription struct {
	description *widget.Label
}

func GenerateTextDescriptionPointer(description string) *TextDescription {
	return &TextDescription{
		description: widget.NewLabel(description),
	}
}

func NewTextDescriptionByPointer(pointer *TextDescription) widget.ToolbarItem {
	return pointer
}

func NewTextDescription(description string) widget.ToolbarItem {
	return GenerateTextDescriptionPointer(description)
}

func (that *TextDescription) Update(description string) {
	that.description.SetText(description)
	that.description.Refresh()
}

func (that *TextDescription) ToolbarObject() fyne.CanvasObject {
	return container.NewBorder(nil, nil, nil, nil, that.description)
}
