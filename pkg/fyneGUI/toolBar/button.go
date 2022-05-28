package toolBar

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type Button struct {
	label    string
	icon     fyne.Resource
	callback func()
}

func (that *Button) ToolbarObject() fyne.CanvasObject {
	return container.NewBorder(nil, nil, nil, nil, widget.NewButtonWithIcon(that.label, that.icon, that.callback))
}

func NewButton(label string, icon fyne.Resource, callback func()) widget.ToolbarItem {
	return &Button{
		label:    label,
		icon:     icon,
		callback: callback,
	}
}
