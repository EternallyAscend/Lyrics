package toolBar

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type Float64Controller struct {
	object      fyne.CanvasObject
	value       float64
	progressBar *widget.ProgressBar
	down        fyne.Resource
	up          fyne.Resource
	step        float64
	limitCheck  func(result float64) bool
	callback    func()
}

func (that *Float64Controller) SetValue() {
	that.progressBar.SetValue(that.value)
	that.progressBar.Refresh()
	that.callback()
}

func (that *Float64Controller) Up() {
	if that.limitCheck(that.value + that.step) {
		that.value += that.step
		that.SetValue()
	}
}

func (that *Float64Controller) Down() {
	if that.limitCheck(that.value - that.step) {
		that.value -= that.step
		that.SetValue()
	}
}

func (that *Float64Controller) ToolbarObject() fyne.CanvasObject {
	that.progressBar = widget.NewProgressBar()
	that.progressBar.SetValue(that.value)

	that.object = container.NewBorder(nil, nil,
		widget.NewButtonWithIcon("", that.down, func() {
			that.Down()
		}),
		widget.NewButtonWithIcon("", that.up, func() {
			that.Up()
		}), that.progressBar)
	return that.object
}

func NewIntController(value float64, down fyne.Resource, up fyne.Resource, step float64, limitChecker func(result float64) bool, callback func()) widget.ToolbarItem {
	return &Float64Controller{
		object:      nil,
		value:       value,
		progressBar: nil,
		down:        down,
		up:          up,
		step:        step,
		limitCheck:  limitChecker,
		callback:    callback,
	}
}
